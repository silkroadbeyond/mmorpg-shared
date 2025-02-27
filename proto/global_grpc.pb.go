// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/global.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GlobalService_RegisterServer_FullMethodName = "/global.GlobalService/RegisterServer"
	GlobalService_GetServerList_FullMethodName  = "/global.GlobalService/GetServerList"
)

// GlobalServiceClient is the client API for GlobalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalServiceClient interface {
	RegisterServer(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*Response, error)
	GetServerList(ctx context.Context, in *Region, opts ...grpc.CallOption) (*ServerList, error)
}

type globalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalServiceClient(cc grpc.ClientConnInterface) GlobalServiceClient {
	return &globalServiceClient{cc}
}

func (c *globalServiceClient) RegisterServer(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, GlobalService_RegisterServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalServiceClient) GetServerList(ctx context.Context, in *Region, opts ...grpc.CallOption) (*ServerList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerList)
	err := c.cc.Invoke(ctx, GlobalService_GetServerList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlobalServiceServer is the server API for GlobalService service.
// All implementations must embed UnimplementedGlobalServiceServer
// for forward compatibility.
type GlobalServiceServer interface {
	RegisterServer(context.Context, *ServerInfo) (*Response, error)
	GetServerList(context.Context, *Region) (*ServerList, error)
	mustEmbedUnimplementedGlobalServiceServer()
}

// UnimplementedGlobalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGlobalServiceServer struct{}

func (UnimplementedGlobalServiceServer) RegisterServer(context.Context, *ServerInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterServer not implemented")
}
func (UnimplementedGlobalServiceServer) GetServerList(context.Context, *Region) (*ServerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerList not implemented")
}
func (UnimplementedGlobalServiceServer) mustEmbedUnimplementedGlobalServiceServer() {}
func (UnimplementedGlobalServiceServer) testEmbeddedByValue()                       {}

// UnsafeGlobalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalServiceServer will
// result in compilation errors.
type UnsafeGlobalServiceServer interface {
	mustEmbedUnimplementedGlobalServiceServer()
}

func RegisterGlobalServiceServer(s grpc.ServiceRegistrar, srv GlobalServiceServer) {
	// If the following call pancis, it indicates UnimplementedGlobalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GlobalService_ServiceDesc, srv)
}

func _GlobalService_RegisterServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalServiceServer).RegisterServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GlobalService_RegisterServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalServiceServer).RegisterServer(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalService_GetServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Region)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalServiceServer).GetServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GlobalService_GetServerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalServiceServer).GetServerList(ctx, req.(*Region))
	}
	return interceptor(ctx, in, info, handler)
}

// GlobalService_ServiceDesc is the grpc.ServiceDesc for GlobalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GlobalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "global.GlobalService",
	HandlerType: (*GlobalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterServer",
			Handler:    _GlobalService_RegisterServer_Handler,
		},
		{
			MethodName: "GetServerList",
			Handler:    _GlobalService_GetServerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/global.proto",
}
