/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package precontainers

import (
	"context"

	api "github.com/containerd/containerd/api/services/precontainers/v1"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/events"
	"github.com/containerd/containerd/metadata"
	"github.com/containerd/containerd/plugin"
	"github.com/containerd/containerd/precontainers"
	ptypes "github.com/containerd/containerd/protobuf/types"
	"github.com/containerd/containerd/services"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/grpc"
	grpcm "google.golang.org/grpc/metadata"
)

func init() {
	plugin.Register(&plugin.Registration{
		Type: plugin.ServicePlugin,
		ID:   services.PrecontainersService,
		Requires: []plugin.Type{
			plugin.EventPlugin,
			plugin.MetadataPlugin,
		},
		InitFn: func(ic *plugin.InitContext) (interface{}, error) {
			m, err := ic.Get(plugin.MetadataPlugin)
			if err != nil {
				return nil, err
			}
			ep, err := ic.Get(plugin.EventPlugin)
			if err != nil {
				return nil, err
			}

			db := m.(*metadata.DB)
			return &local{
				Store:     metadata.NewPreContainerStore(db),
				db:        db,
				publisher: ep.(events.Publisher),
			}, nil
		},
	})
}

type local struct {
	precontainers.Store
	db        *metadata.DB
	publisher events.Publisher
}

var _ api.PrecontainersClient = &local{}

func (l *local) Get(ctx context.Context, req *api.GetPreContainerRequest, _ ...grpc.CallOption) (*api.GetPreContainerResponse, error) {
	var resp api.GetPreContainerResponse

	return &resp, errdefs.ToGRPC(l.withStoreView(ctx, func(ctx context.Context) error {
		container, err := l.Store.Get(ctx, req.Function)
		if err != nil {
			return err
		}
		containerpb := precontainerToProto(&container)
		resp.Container = containerpb

		return nil
	}))
}

func (l *local) Preload(ctx context.Context, req *api.PreloadContainerRequest, _ ...grpc.CallOption) (*api.PreloadContainerResponse, error) {
	var resp api.PreloadContainerResponse

	if err := l.withStoreUpdate(ctx, func(ctx context.Context) error {
		container := precontainerFromProto(req.Container)

		created, err := l.Store.Preload(ctx, container)
		if err != nil {
			return err
		}

		resp.Container = precontainerToProto(&created)

		return nil
	}); err != nil {
		return &resp, errdefs.ToGRPC(err)
	}

	return &resp, nil
}

func (l *local) Delete(ctx context.Context, req *api.DeletePreContainerRequest, _ ...grpc.CallOption) (*ptypes.Empty, error) {
	if err := l.withStoreUpdate(ctx, func(ctx context.Context) error {
		return l.Store.Delete(ctx, req.Function)
	}); err != nil {
		return &ptypes.Empty{}, errdefs.ToGRPC(err)
	}

	return &ptypes.Empty{}, nil
}

func (l *local) withStore(ctx context.Context, fn func(ctx context.Context) error) func(tx *bolt.Tx) error {
	return func(tx *bolt.Tx) error {
		return fn(metadata.WithTransactionContext(ctx, tx))
	}
}

func (l *local) withStoreView(ctx context.Context, fn func(ctx context.Context) error) error {
	return l.db.View(l.withStore(ctx, fn))
}

func (l *local) withStoreUpdate(ctx context.Context, fn func(ctx context.Context) error) error {
	return l.db.Update(l.withStore(ctx, fn))
}

type localStream struct {
	ctx        context.Context
	containers []*api.Container
	i          int
}

func (s *localStream) Context() context.Context {
	return s.ctx
}

func (s *localStream) CloseSend() error {
	return nil
}

func (s *localStream) Header() (grpcm.MD, error) {
	return nil, nil
}

func (s *localStream) Trailer() grpcm.MD {
	return nil
}

func (s *localStream) SendMsg(m interface{}) error {
	return nil
}

func (s *localStream) RecvMsg(m interface{}) error {
	return nil
}
