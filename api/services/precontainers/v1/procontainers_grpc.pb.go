// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/qi0523/containerd/api/services/precontainers/v1/procontainers.proto

package precontainers

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

// PrecontainersClient is the client API for Precontainers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrecontainersClient interface {
	Preload(ctx context.Context, in *PreloadContainerRequest, opts ...grpc.CallOption) (*PreloadContainerResponse, error)
	Get(ctx context.Context, in *GetPreContainerRequest, opts ...grpc.CallOption) (*GetPreContainerResponse, error)
	Delete(ctx context.Context, in *DeletePreContainerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type precontainersClient struct {
	cc grpc.ClientConnInterface
}

func NewPrecontainersClient(cc grpc.ClientConnInterface) PrecontainersClient {
	return &precontainersClient{cc}
}

func (c *precontainersClient) Preload(ctx context.Context, in *PreloadContainerRequest, opts ...grpc.CallOption) (*PreloadContainerResponse, error) {
	out := new(PreloadContainerResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.precontainers.v1.Precontainers/Preload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *precontainersClient) Get(ctx context.Context, in *GetPreContainerRequest, opts ...grpc.CallOption) (*GetPreContainerResponse, error) {
	out := new(GetPreContainerResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.precontainers.v1.Precontainers/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *precontainersClient) Delete(ctx context.Context, in *DeletePreContainerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/containerd.services.precontainers.v1.Precontainers/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrecontainersServer is the server API for Precontainers service.
// All implementations must embed UnimplementedPrecontainersServer
// for forward compatibility
type PrecontainersServer interface {
	Preload(context.Context, *PreloadContainerRequest) (*PreloadContainerResponse, error)
	Get(context.Context, *GetPreContainerRequest) (*GetPreContainerResponse, error)
	Delete(context.Context, *DeletePreContainerRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPrecontainersServer()
}

// UnimplementedPrecontainersServer must be embedded to have forward compatible implementations.
type UnimplementedPrecontainersServer struct {
}

func (UnimplementedPrecontainersServer) Preload(context.Context, *PreloadContainerRequest) (*PreloadContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Preload not implemented")
}
func (UnimplementedPrecontainersServer) Get(context.Context, *GetPreContainerRequest) (*GetPreContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPrecontainersServer) Delete(context.Context, *DeletePreContainerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPrecontainersServer) mustEmbedUnimplementedPrecontainersServer() {}

// UnsafePrecontainersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrecontainersServer will
// result in compilation errors.
type UnsafePrecontainersServer interface {
	mustEmbedUnimplementedPrecontainersServer()
}

func RegisterPrecontainersServer(s grpc.ServiceRegistrar, srv PrecontainersServer) {
	s.RegisterService(&Precontainers_ServiceDesc, srv)
}

func _Precontainers_Preload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreloadContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrecontainersServer).Preload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.precontainers.v1.Precontainers/Preload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrecontainersServer).Preload(ctx, req.(*PreloadContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Precontainers_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrecontainersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.precontainers.v1.Precontainers/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrecontainersServer).Get(ctx, req.(*GetPreContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Precontainers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePreContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrecontainersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.precontainers.v1.Precontainers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrecontainersServer).Delete(ctx, req.(*DeletePreContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Precontainers_ServiceDesc is the grpc.ServiceDesc for Precontainers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Precontainers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.precontainers.v1.Precontainers",
	HandlerType: (*PrecontainersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Preload",
			Handler:    _Precontainers_Preload_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Precontainers_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Precontainers_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/qi0523/containerd/api/services/precontainers/v1/procontainers.proto",
}
