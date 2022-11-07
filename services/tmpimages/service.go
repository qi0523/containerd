package tmpimages

import (
	"context"
	"errors"

	tmpimagesapi "github.com/containerd/containerd/api/services/tmpimages/v1"
	"github.com/containerd/containerd/plugin"
	ptypes "github.com/containerd/containerd/protobuf/types"
	"github.com/containerd/containerd/services"
	"google.golang.org/grpc"
)

func init() {
	plugin.Register(&plugin.Registration{
		Type: plugin.GRPCPlugin,
		ID:   "tmpimages",
		Requires: []plugin.Type{
			plugin.ServicePlugin,
		},
		InitFn: func(ic *plugin.InitContext) (interface{}, error) {
			plugins, err := ic.GetByType(plugin.ServicePlugin)
			if err != nil {
				return nil, err
			}
			p, ok := plugins[services.TmpImagesService]
			if !ok {
				return nil, errors.New("tmpimages service not found")
			}
			i, err := p.Instance()
			if err != nil {
				return nil, err
			}
			return &service{local: i.(tmpimagesapi.TmpImagesClient)}, nil
		},
	})
}

type service struct {
	local tmpimagesapi.TmpImagesClient
	tmpimagesapi.UnimplementedTmpImagesServer
}

var _ tmpimagesapi.TmpImagesServer = &service{}

func (s *service) Register(server *grpc.Server) error {
	tmpimagesapi.RegisterTmpImagesServer(server, s)
	return nil
}

func (s *service) InsertTmpImage(ctx context.Context, req *tmpimagesapi.CreateTmpImageRequest) (*tmpimagesapi.CreateTmpImageResponse, error) {
	return s.local.InsertTmpImage(ctx, req)
}

func (s *service) GetTmpImage(ctx context.Context, req *tmpimagesapi.GetTmpImageRequest) (*tmpimagesapi.GetTmpImageResponse, error) {
	return s.local.GetTmpImage(ctx, req)
}

func (s *service) Delete(ctx context.Context, req *tmpimagesapi.DeleteTmpImageRequest) (*ptypes.Empty, error) {
	return s.local.Delete(ctx, req)
}
