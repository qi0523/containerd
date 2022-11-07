package tmpimages

import (
	"context"

	tmpimagesapi "github.com/containerd/containerd/api/services/tmpimages/v1"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/metadata"
	"github.com/containerd/containerd/plugin"
	ptypes "github.com/containerd/containerd/protobuf/types"
	"github.com/containerd/containerd/services"
	"github.com/containerd/containerd/tmpimages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	plugin.Register(&plugin.Registration{
		Type: plugin.ServicePlugin,
		ID:   services.TmpImagesService,
		Requires: []plugin.Type{
			plugin.MetadataPlugin,
		},
		InitFn: func(ic *plugin.InitContext) (interface{}, error) {
			m, err := ic.Get(plugin.MetadataPlugin)
			if err != nil {
				return nil, err
			}
			return &local{
				store: metadata.NewTmpImageStore(m.(*metadata.DB)),
			}, nil
		},
	})
}

type local struct {
	store tmpimages.Store
}

var _ tmpimagesapi.TmpImagesClient = &local{}

func (l *local) InsertTmpImage(ctx context.Context, req *tmpimagesapi.CreateTmpImageRequest, _ ...grpc.CallOption) (*tmpimagesapi.CreateTmpImageResponse, error) {
	if req.Image.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "TmpImage.Name required")
	}
	var (
		tmpimage = tmpImageFromProto(req.Image)
		resp     tmpimagesapi.CreateTmpImageResponse
	)

	created, err := l.store.InsertTmpImage(ctx, tmpimage)
	if err != nil {
		return nil, errdefs.ToGRPC(err)
	}
	resp.Image = tmpImageToProto(&created)
	return &resp, nil
}

func (l *local) GetTmpImage(ctx context.Context, req *tmpimagesapi.GetTmpImageRequest, _ ...grpc.CallOption) (*tmpimagesapi.GetTmpImageResponse, error) {
	tmpimage, err := l.store.GetTmpImage(ctx, req.Name)
	if err != nil {
		return nil, errdefs.ToGRPC(err)
	}
	tmpimagepb := tmpImageToProto(&tmpimage)
	return &tmpimagesapi.GetTmpImageResponse{
		Image: tmpimagepb,
	}, nil
}

func (l *local) Delete(ctx context.Context, req *tmpimagesapi.DeleteTmpImageRequest, opts ...grpc.CallOption) (*ptypes.Empty, error) {
	if err := l.store.Delete(ctx, req.Name); err != nil {
		return nil, errdefs.ToGRPC(err)
	}
	return &ptypes.Empty{}, nil
}
