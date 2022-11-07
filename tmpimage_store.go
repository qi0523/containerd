package containerd

import (
	"context"
	"strings"

	tmpimagesapi "github.com/containerd/containerd/api/services/tmpimages/v1"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/tmpimages"
)

type remoteTmpImages struct {
	client tmpimagesapi.TmpImagesClient
}

func NewTmpImageStoreFromClient(client tmpimagesapi.TmpImagesClient) tmpimages.Store {
	return &remoteTmpImages{
		client: client,
	}
}

func (s *remoteTmpImages) GetTmpImage(ctx context.Context, name string) (tmpimages.TmpImage, error) {
	resp, err := s.client.GetTmpImage(ctx, &tmpimagesapi.GetTmpImageRequest{
		Name: name,
	})
	if err != nil {
		return tmpimages.TmpImage{}, errdefs.FromGRPC(err)
	}
	return tmpimageFromProto(resp.Image), nil
}

func (s *remoteTmpImages) InsertTmpImage(ctx context.Context, image tmpimages.TmpImage) (tmpimages.TmpImage, error) {
	created, err := s.client.InsertTmpImage(ctx, &tmpimagesapi.CreateTmpImageRequest{
		Image: tmpimageToProto(&image),
	})
	if err != nil {
		return tmpimages.TmpImage{}, errdefs.FromGRPC(err)
	}
	return tmpimageFromProto(created.Image), nil
}

func (s *remoteTmpImages) Delete(ctx context.Context, name string) error {
	_, err := s.client.Delete(ctx, &tmpimagesapi.DeleteTmpImageRequest{
		Name: name,
	})
	return errdefs.FromGRPC(err)
}

func tmpimageToProto(image *tmpimages.TmpImage) *tmpimagesapi.TmpImage {
	return &tmpimagesapi.TmpImage{
		Name:   image.Name,
		Target: descToProto(&image.Target),
	}
}

func tmpimageFromProto(tmpimagepd *tmpimagesapi.TmpImage) tmpimages.TmpImage {
	return tmpimages.TmpImage{
		Name:   tmpimagepd.Name,
		Target: descFromProto(tmpimagepd.Target),
	}
}

func tmpImageName(name string) string {
	// name: docker.io/qi0523/ubuntu:latest OR 192.168.1.1:3000/qi0523/ubuntu:latest
	return name[strings.Index(name, "/")+1:]
}
