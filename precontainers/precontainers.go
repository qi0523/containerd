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

	"github.com/containerd/containerd/containers"
)

// Store interacts with the underlying container storage
type Store interface {

	// Preload a container in the store from the provided container.
	Preload(ctx context.Context, container containers.Container) (containers.Container, error)

	// Get a container using the id.
	//
	// Container object is returned on success. If the id is not known to the
	// store, an error will be returned.
	Get(ctx context.Context, function string) (containers.Container, error)

	// Delete a container using the function.
	//
	// nil will be returned on success. If the container is not known to the
	// store, ErrNotFound will be returned.
	Delete(ctx context.Context, function string) error
}
