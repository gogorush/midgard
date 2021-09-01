// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MidgardClient is the client API for Midgard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MidgardClient interface {
	Ping(ctx context.Context, in *PingInput, opts ...grpc.CallOption) (*PingOutput, error)
	AllocateURL(ctx context.Context, in *AllocateURLInput, opts ...grpc.CallOption) (*AllocateURLOutput, error)
	CodeToImage(ctx context.Context, in *CodeToImageInput, opts ...grpc.CallOption) (*CodeToImageOutput, error)
	ListDaemons(ctx context.Context, in *ListDaemonsInput, opts ...grpc.CallOption) (*ListDaemonsOutput, error)
}

type midgardClient struct {
	cc grpc.ClientConnInterface
}

func NewMidgardClient(cc grpc.ClientConnInterface) MidgardClient {
	return &midgardClient{cc}
}

func (c *midgardClient) Ping(ctx context.Context, in *PingInput, opts ...grpc.CallOption) (*PingOutput, error) {
	out := new(PingOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) AllocateURL(ctx context.Context, in *AllocateURLInput, opts ...grpc.CallOption) (*AllocateURLOutput, error) {
	out := new(AllocateURLOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/AllocateURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) CodeToImage(ctx context.Context, in *CodeToImageInput, opts ...grpc.CallOption) (*CodeToImageOutput, error) {
	out := new(CodeToImageOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/CodeToImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) ListDaemons(ctx context.Context, in *ListDaemonsInput, opts ...grpc.CallOption) (*ListDaemonsOutput, error) {
	out := new(ListDaemonsOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/ListDaemons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MidgardServer is the server API for Midgard service.
// All implementations must embed UnimplementedMidgardServer
// for forward compatibility
type MidgardServer interface {
	Ping(context.Context, *PingInput) (*PingOutput, error)
	AllocateURL(context.Context, *AllocateURLInput) (*AllocateURLOutput, error)
	CodeToImage(context.Context, *CodeToImageInput) (*CodeToImageOutput, error)
	ListDaemons(context.Context, *ListDaemonsInput) (*ListDaemonsOutput, error)
	mustEmbedUnimplementedMidgardServer()
}

// UnimplementedMidgardServer must be embedded to have forward compatible implementations.
type UnimplementedMidgardServer struct {
}

func (UnimplementedMidgardServer) Ping(context.Context, *PingInput) (*PingOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMidgardServer) AllocateURL(context.Context, *AllocateURLInput) (*AllocateURLOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateURL not implemented")
}
func (UnimplementedMidgardServer) CodeToImage(context.Context, *CodeToImageInput) (*CodeToImageOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeToImage not implemented")
}
func (UnimplementedMidgardServer) ListDaemons(context.Context, *ListDaemonsInput) (*ListDaemonsOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDaemons not implemented")
}
func (UnimplementedMidgardServer) mustEmbedUnimplementedMidgardServer() {}

// UnsafeMidgardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MidgardServer will
// result in compilation errors.
type UnsafeMidgardServer interface {
	mustEmbedUnimplementedMidgardServer()
}

func RegisterMidgardServer(s grpc.ServiceRegistrar, srv MidgardServer) {
	s.RegisterService(&_Midgard_serviceDesc, srv)
}

func _Midgard_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).Ping(ctx, req.(*PingInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_AllocateURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateURLInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).AllocateURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/AllocateURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).AllocateURL(ctx, req.(*AllocateURLInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_CodeToImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeToImageInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).CodeToImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/CodeToImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).CodeToImage(ctx, req.(*CodeToImageInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_ListDaemons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDaemonsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).ListDaemons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/ListDaemons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).ListDaemons(ctx, req.(*ListDaemonsInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Midgard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Midgard",
	HandlerType: (*MidgardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Midgard_Ping_Handler,
		},
		{
			MethodName: "AllocateURL",
			Handler:    _Midgard_AllocateURL_Handler,
		},
		{
			MethodName: "CodeToImage",
			Handler:    _Midgard_CodeToImage_Handler,
		},
		{
			MethodName: "ListDaemons",
			Handler:    _Midgard_ListDaemons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "midgard.proto",
}
