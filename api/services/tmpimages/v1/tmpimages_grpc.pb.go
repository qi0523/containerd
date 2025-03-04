// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/services/tmpimages/v1/tmpimages.proto

package tmpimages

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TmpImagesClient is the client API for TmpImages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TmpImagesClient interface {
	// Insert returns an image by name.
	InsertTmpImage(ctx context.Context, in *CreateTmpImageRequest, opts ...grpc.CallOption) (*CreateTmpImageResponse, error)
	// GetTmpImage
	GetTmpImage(ctx context.Context, in *GetTmpImageRequest, opts ...grpc.CallOption) (*GetTmpImageResponse, error)
	// Delete deletes the image by name.
	Delete(ctx context.Context, in *DeleteTmpImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tmpImagesClient struct {
	cc grpc.ClientConnInterface
}

func NewTmpImagesClient(cc grpc.ClientConnInterface) TmpImagesClient {
	return &tmpImagesClient{cc}
}

func (c *tmpImagesClient) InsertTmpImage(ctx context.Context, in *CreateTmpImageRequest, opts ...grpc.CallOption) (*CreateTmpImageResponse, error) {
	out := new(CreateTmpImageResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.tmpimages.v1.TmpImages/InsertTmpImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tmpImagesClient) GetTmpImage(ctx context.Context, in *GetTmpImageRequest, opts ...grpc.CallOption) (*GetTmpImageResponse, error) {
	out := new(GetTmpImageResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.tmpimages.v1.TmpImages/GetTmpImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tmpImagesClient) Delete(ctx context.Context, in *DeleteTmpImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.tmpimages.v1.TmpImages/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TmpImagesServer is the server API for TmpImages service.
// All implementations must embed UnimplementedTmpImagesServer
// for forward compatibility
type TmpImagesServer interface {
	// Insert returns an image by name.
	InsertTmpImage(context.Context, *CreateTmpImageRequest) (*CreateTmpImageResponse, error)
	// GetTmpImage
	GetTmpImage(context.Context, *GetTmpImageRequest) (*GetTmpImageResponse, error)
	// Delete deletes the image by name.
	Delete(context.Context, *DeleteTmpImageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTmpImagesServer()
}

// UnimplementedTmpImagesServer must be embedded to have forward compatible implementations.
type UnimplementedTmpImagesServer struct {
}

func (UnimplementedTmpImagesServer) InsertTmpImage(context.Context, *CreateTmpImageRequest) (*CreateTmpImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertTmpImage not implemented")
}
func (UnimplementedTmpImagesServer) GetTmpImage(context.Context, *GetTmpImageRequest) (*GetTmpImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTmpImage not implemented")
}
func (UnimplementedTmpImagesServer) Delete(context.Context, *DeleteTmpImageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTmpImagesServer) mustEmbedUnimplementedTmpImagesServer() {}

// UnsafeTmpImagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TmpImagesServer will
// result in compilation errors.
type UnsafeTmpImagesServer interface {
	mustEmbedUnimplementedTmpImagesServer()
}

func RegisterTmpImagesServer(s grpc.ServiceRegistrar, srv TmpImagesServer) {
	s.RegisterService(&TmpImages_ServiceDesc, srv)
}

func _TmpImages_InsertTmpImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTmpImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TmpImagesServer).InsertTmpImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.tmpimages.v1.TmpImages/InsertTmpImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TmpImagesServer).InsertTmpImage(ctx, req.(*CreateTmpImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TmpImages_GetTmpImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTmpImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TmpImagesServer).GetTmpImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.tmpimages.v1.TmpImages/GetTmpImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TmpImagesServer).GetTmpImage(ctx, req.(*GetTmpImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TmpImages_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTmpImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TmpImagesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.tmpimages.v1.TmpImages/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TmpImagesServer).Delete(ctx, req.(*DeleteTmpImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TmpImages_ServiceDesc is the grpc.ServiceDesc for TmpImages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TmpImages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.tmpimages.v1.TmpImages",
	HandlerType: (*TmpImagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertTmpImage",
			Handler:    _TmpImages_InsertTmpImage_Handler,
		},
		{
			MethodName: "GetTmpImage",
			Handler:    _TmpImages_GetTmpImage_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TmpImages_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/tmpimages/v1/tmpimages.proto",
}
