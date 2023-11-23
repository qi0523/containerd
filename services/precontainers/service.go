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
	"errors"

	api "github.com/containerd/containerd/api/services/precontainers/v1"
	"github.com/containerd/containerd/plugin"
	ptypes "github.com/containerd/containerd/protobuf/types"
	"github.com/containerd/containerd/services"
	"google.golang.org/grpc"
)

func init() {
	plugin.Register(&plugin.Registration{
		Type: plugin.GRPCPlugin,
		ID:   "precontainers",
		Requires: []plugin.Type{
			plugin.ServicePlugin,
		},
		InitFn: func(ic *plugin.InitContext) (interface{}, error) {
			plugins, err := ic.GetByType(plugin.ServicePlugin)
			if err != nil {
				return nil, err
			}
			p, ok := plugins[services.PrecontainersService]
			if !ok {
				return nil, errors.New("precontainers service not found")
			}
			i, err := p.Instance()
			if err != nil {
				return nil, err
			}
			return &service{local: i.(api.PrecontainersClient)}, nil
		},
	})
}

type service struct {
	local api.PrecontainersClient
	api.UnimplementedPrecontainersServer
}

var _ api.PrecontainersServer = &service{}

func (s *service) Register(server *grpc.Server) error {
	api.RegisterPrecontainersServer(server, s)
	return nil
}

func (s *service) Get(ctx context.Context, req *api.GetPreContainerRequest) (*api.GetPreContainerResponse, error) {
	return s.local.Get(ctx, req)
}

func (s *service) Preload(ctx context.Context, req *api.PreloadContainerRequest) (*api.PreloadContainerResponse, error) {
	return s.local.Preload(ctx, req)
}

func (s *service) Delete(ctx context.Context, req *api.DeletePreContainerRequest) (*ptypes.Empty, error) {
	return s.local.Delete(ctx, req)
}
