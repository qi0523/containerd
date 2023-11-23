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

package containerd

import (
	"context"

	precontainersapi "github.com/containerd/containerd/api/services/precontainers/v1"
	"github.com/containerd/containerd/containers"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/precontainers"
	"github.com/containerd/containerd/protobuf"
	ptypes "github.com/containerd/containerd/protobuf/types"
	"github.com/containerd/typeurl"
)

type remotePrecontainers struct {
	client precontainersapi.PrecontainersClient
}

var _ precontainers.Store = &remotePrecontainers{}

// NewRemoteContainerStore returns the container Store connected with the provided client
func NewRemotePrecontainerStore(client precontainersapi.PrecontainersClient) precontainers.Store {
	return &remotePrecontainers{
		client: client,
	}
}

func (r *remotePrecontainers) Preload(ctx context.Context, container containers.Container) (containers.Container, error) {
	created, err := r.client.Preload(ctx, &precontainersapi.PreloadContainerRequest{
		Container: precontainerToProto(&container),
	})
	if err != nil {
		return containers.Container{}, errdefs.FromGRPC(err)
	}

	return precontainerFromProto(created.Container), nil

}

func (r *remotePrecontainers) Get(ctx context.Context, function string) (containers.Container, error) {
	resp, err := r.client.Get(ctx, &precontainersapi.GetPreContainerRequest{
		Function: function,
	})
	if err != nil {
		return containers.Container{}, errdefs.FromGRPC(err)
	}

	return precontainerFromProto(resp.Container), nil
}

func (r *remotePrecontainers) Delete(ctx context.Context, function string) error {
	_, err := r.client.Delete(ctx, &precontainersapi.DeletePreContainerRequest{
		Function: function,
	})

	return errdefs.FromGRPC(err)

}

func precontainerToProto(container *containers.Container) *precontainersapi.Container {
	extensions := make(map[string]*ptypes.Any)
	for k, v := range container.Extensions {
		extensions[k] = protobuf.FromAny(v)
	}
	return &precontainersapi.Container{
		Id:     container.ID,
		Labels: container.Labels,
		Image:  container.Image,
		Runtime: &precontainersapi.Container_Runtime{
			Name:    container.Runtime.Name,
			Options: protobuf.FromAny(container.Runtime.Options),
		},
		Spec:        protobuf.FromAny(container.Spec),
		Snapshotter: container.Snapshotter,
		SnapshotKey: container.SnapshotKey,
		Extensions:  extensions,
		Sandbox:     container.SandboxID,
	}
}

func precontainerFromProto(containerpb *precontainersapi.Container) containers.Container {
	var runtime containers.RuntimeInfo
	if containerpb.Runtime != nil {
		runtime = containers.RuntimeInfo{
			Name:    containerpb.Runtime.Name,
			Options: containerpb.Runtime.Options,
		}
	}
	extensions := make(map[string]typeurl.Any)
	for k, v := range containerpb.Extensions {
		v := v
		extensions[k] = v
	}
	return containers.Container{
		ID:          containerpb.Id,
		Labels:      containerpb.Labels,
		Image:       containerpb.Image,
		Runtime:     runtime,
		Spec:        containerpb.Spec,
		Snapshotter: containerpb.Snapshotter,
		SnapshotKey: containerpb.SnapshotKey,
		CreatedAt:   protobuf.FromTimestamp(containerpb.CreatedAt),
		UpdatedAt:   protobuf.FromTimestamp(containerpb.UpdatedAt),
		Extensions:  extensions,
		SandboxID:   containerpb.Sandbox,
	}
}
