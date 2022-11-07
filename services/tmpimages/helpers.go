package tmpimages

import (
	tmpimagesapi "github.com/containerd/containerd/api/services/tmpimages/v1"
	"github.com/containerd/containerd/api/types"
	"github.com/containerd/containerd/tmpimages"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

func tmpImageToProto(tmpimage *tmpimages.TmpImage) *tmpimagesapi.TmpImage {
	return &tmpimagesapi.TmpImage{
		Name:   tmpimage.Name,
		Target: descToProto(&tmpimage.Target),
	}
}

func tmpImageFromProto(tmpimagepb *tmpimagesapi.TmpImage) tmpimages.TmpImage {
	return tmpimages.TmpImage{
		Name:   tmpimagepb.Name,
		Target: descFromProto(tmpimagepb.Target),
	}
}

func descFromProto(desc *types.Descriptor) ocispec.Descriptor {
	return ocispec.Descriptor{
		MediaType:   desc.MediaType,
		Size:        desc.Size,
		Digest:      digest.Digest(desc.Digest),
		Annotations: desc.Annotations,
	}
}

func descToProto(desc *ocispec.Descriptor) *types.Descriptor {
	return &types.Descriptor{
		MediaType:   desc.MediaType,
		Size:        desc.Size,
		Digest:      desc.Digest.String(),
		Annotations: desc.Annotations,
	}
}
