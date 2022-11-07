package tmpimages

import (
	"context"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type TmpImage struct {
	// Name of the image.
	//
	// To be pulled, it must be a reference compatible with resolvers.
	//
	// This field is required.
	Name string

	// Target describes the root content for this image. Typically, this is
	// a manifest, index or manifest list.
	Target ocispec.Descriptor
}

type Store interface {
	GetTmpImage(ctx context.Context, name string) (TmpImage, error)
	InsertTmpImage(ctx context.Context, image TmpImage) (TmpImage, error)
	Delete(ctx context.Context, name string) error
}
