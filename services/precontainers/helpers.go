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
	api "github.com/containerd/containerd/api/services/precontainers/v1"
	"github.com/containerd/containerd/containers"
	"github.com/containerd/containerd/protobuf"
	"github.com/containerd/containerd/protobuf/types"
	"github.com/containerd/typeurl"
)

func precontainerToProto(container *containers.Container) *api.Container {
	extensions := make(map[string]*types.Any)
	for k, v := range container.Extensions {
		extensions[k] = protobuf.FromAny(v)
	}
	return &api.Container{
		Id:     container.ID,
		Labels: container.Labels,
		Image:  container.Image,
		Runtime: &api.Container_Runtime{
			Name:    container.Runtime.Name,
			Options: protobuf.FromAny(container.Runtime.Options),
		},
		Spec:        protobuf.FromAny(container.Spec),
		Snapshotter: container.Snapshotter,
		SnapshotKey: container.SnapshotKey,
		CreatedAt:   protobuf.ToTimestamp(container.CreatedAt),
		UpdatedAt:   protobuf.ToTimestamp(container.UpdatedAt),
		Extensions:  extensions,
		Sandbox:     container.SandboxID,
	}
}

func precontainerFromProto(containerpb *api.Container) containers.Container {
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
		Extensions:  extensions,
		SandboxID:   containerpb.Sandbox,
	}
}
