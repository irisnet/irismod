// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/oracle/tx.proto

package oracle

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_CreateFeed_FullMethodName = "/irismod.oracle.Msg/CreateFeed"
	Msg_EditFeed_FullMethodName   = "/irismod.oracle.Msg/EditFeed"
	Msg_StartFeed_FullMethodName  = "/irismod.oracle.Msg/StartFeed"
	Msg_PauseFeed_FullMethodName  = "/irismod.oracle.Msg/PauseFeed"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateFeed defines a method for creating a new feed
	CreateFeed(ctx context.Context, in *MsgCreateFeed, opts ...grpc.CallOption) (*MsgCreateFeedResponse, error)
	// EditFeed defines a method for editing a feed
	EditFeed(ctx context.Context, in *MsgEditFeed, opts ...grpc.CallOption) (*MsgEditFeedResponse, error)
	// StartFeed defines a method for starting a feed
	StartFeed(ctx context.Context, in *MsgStartFeed, opts ...grpc.CallOption) (*MsgStartFeedResponse, error)
	// PauseFeed defines a method for pausing a feed
	PauseFeed(ctx context.Context, in *MsgPauseFeed, opts ...grpc.CallOption) (*MsgPauseFeedResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateFeed(ctx context.Context, in *MsgCreateFeed, opts ...grpc.CallOption) (*MsgCreateFeedResponse, error) {
	out := new(MsgCreateFeedResponse)
	err := c.cc.Invoke(ctx, Msg_CreateFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditFeed(ctx context.Context, in *MsgEditFeed, opts ...grpc.CallOption) (*MsgEditFeedResponse, error) {
	out := new(MsgEditFeedResponse)
	err := c.cc.Invoke(ctx, Msg_EditFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StartFeed(ctx context.Context, in *MsgStartFeed, opts ...grpc.CallOption) (*MsgStartFeedResponse, error) {
	out := new(MsgStartFeedResponse)
	err := c.cc.Invoke(ctx, Msg_StartFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PauseFeed(ctx context.Context, in *MsgPauseFeed, opts ...grpc.CallOption) (*MsgPauseFeedResponse, error) {
	out := new(MsgPauseFeedResponse)
	err := c.cc.Invoke(ctx, Msg_PauseFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateFeed defines a method for creating a new feed
	CreateFeed(context.Context, *MsgCreateFeed) (*MsgCreateFeedResponse, error)
	// EditFeed defines a method for editing a feed
	EditFeed(context.Context, *MsgEditFeed) (*MsgEditFeedResponse, error)
	// StartFeed defines a method for starting a feed
	StartFeed(context.Context, *MsgStartFeed) (*MsgStartFeedResponse, error)
	// PauseFeed defines a method for pausing a feed
	PauseFeed(context.Context, *MsgPauseFeed) (*MsgPauseFeedResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateFeed(context.Context, *MsgCreateFeed) (*MsgCreateFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeed not implemented")
}
func (UnimplementedMsgServer) EditFeed(context.Context, *MsgEditFeed) (*MsgEditFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditFeed not implemented")
}
func (UnimplementedMsgServer) StartFeed(context.Context, *MsgStartFeed) (*MsgStartFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFeed not implemented")
}
func (UnimplementedMsgServer) PauseFeed(context.Context, *MsgPauseFeed) (*MsgPauseFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseFeed not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateFeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateFeed(ctx, req.(*MsgCreateFeed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditFeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EditFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditFeed(ctx, req.(*MsgEditFeed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StartFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartFeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StartFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StartFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StartFeed(ctx, req.(*MsgStartFeed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PauseFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPauseFeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PauseFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_PauseFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PauseFeed(ctx, req.(*MsgPauseFeed))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.oracle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeed",
			Handler:    _Msg_CreateFeed_Handler,
		},
		{
			MethodName: "EditFeed",
			Handler:    _Msg_EditFeed_Handler,
		},
		{
			MethodName: "StartFeed",
			Handler:    _Msg_StartFeed_Handler,
		},
		{
			MethodName: "PauseFeed",
			Handler:    _Msg_PauseFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/oracle/tx.proto",
}
