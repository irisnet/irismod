// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/mt/tx.proto

package mt

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
	Msg_IssueDenom_FullMethodName    = "/irismod.mt.Msg/IssueDenom"
	Msg_TransferDenom_FullMethodName = "/irismod.mt.Msg/TransferDenom"
	Msg_MintMT_FullMethodName        = "/irismod.mt.Msg/MintMT"
	Msg_EditMT_FullMethodName        = "/irismod.mt.Msg/EditMT"
	Msg_TransferMT_FullMethodName    = "/irismod.mt.Msg/TransferMT"
	Msg_BurnMT_FullMethodName        = "/irismod.mt.Msg/BurnMT"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// IssueDenom defines a method for issuing a denom.
	IssueDenom(ctx context.Context, in *MsgIssueDenom, opts ...grpc.CallOption) (*MsgIssueDenomResponse, error)
	// TransferDenom defines a method for transferring a denom.
	TransferDenom(ctx context.Context, in *MsgTransferDenom, opts ...grpc.CallOption) (*MsgTransferDenomResponse, error)
	// MintMT defines a method for creating a new MT or minting amounts of an existing MT
	MintMT(ctx context.Context, in *MsgMintMT, opts ...grpc.CallOption) (*MsgMintMTResponse, error)
	// EditMT defines a method for editing an MT.
	EditMT(ctx context.Context, in *MsgEditMT, opts ...grpc.CallOption) (*MsgEditMTResponse, error)
	// TransferMT defines a method for transferring an MT.
	TransferMT(ctx context.Context, in *MsgTransferMT, opts ...grpc.CallOption) (*MsgTransferMTResponse, error)
	// BurnMT defines a method for burning an MT.
	BurnMT(ctx context.Context, in *MsgBurnMT, opts ...grpc.CallOption) (*MsgBurnMTResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IssueDenom(ctx context.Context, in *MsgIssueDenom, opts ...grpc.CallOption) (*MsgIssueDenomResponse, error) {
	out := new(MsgIssueDenomResponse)
	err := c.cc.Invoke(ctx, Msg_IssueDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferDenom(ctx context.Context, in *MsgTransferDenom, opts ...grpc.CallOption) (*MsgTransferDenomResponse, error) {
	out := new(MsgTransferDenomResponse)
	err := c.cc.Invoke(ctx, Msg_TransferDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintMT(ctx context.Context, in *MsgMintMT, opts ...grpc.CallOption) (*MsgMintMTResponse, error) {
	out := new(MsgMintMTResponse)
	err := c.cc.Invoke(ctx, Msg_MintMT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditMT(ctx context.Context, in *MsgEditMT, opts ...grpc.CallOption) (*MsgEditMTResponse, error) {
	out := new(MsgEditMTResponse)
	err := c.cc.Invoke(ctx, Msg_EditMT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferMT(ctx context.Context, in *MsgTransferMT, opts ...grpc.CallOption) (*MsgTransferMTResponse, error) {
	out := new(MsgTransferMTResponse)
	err := c.cc.Invoke(ctx, Msg_TransferMT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnMT(ctx context.Context, in *MsgBurnMT, opts ...grpc.CallOption) (*MsgBurnMTResponse, error) {
	out := new(MsgBurnMTResponse)
	err := c.cc.Invoke(ctx, Msg_BurnMT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// IssueDenom defines a method for issuing a denom.
	IssueDenom(context.Context, *MsgIssueDenom) (*MsgIssueDenomResponse, error)
	// TransferDenom defines a method for transferring a denom.
	TransferDenom(context.Context, *MsgTransferDenom) (*MsgTransferDenomResponse, error)
	// MintMT defines a method for creating a new MT or minting amounts of an existing MT
	MintMT(context.Context, *MsgMintMT) (*MsgMintMTResponse, error)
	// EditMT defines a method for editing an MT.
	EditMT(context.Context, *MsgEditMT) (*MsgEditMTResponse, error)
	// TransferMT defines a method for transferring an MT.
	TransferMT(context.Context, *MsgTransferMT) (*MsgTransferMTResponse, error)
	// BurnMT defines a method for burning an MT.
	BurnMT(context.Context, *MsgBurnMT) (*MsgBurnMTResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) IssueDenom(context.Context, *MsgIssueDenom) (*MsgIssueDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueDenom not implemented")
}
func (UnimplementedMsgServer) TransferDenom(context.Context, *MsgTransferDenom) (*MsgTransferDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferDenom not implemented")
}
func (UnimplementedMsgServer) MintMT(context.Context, *MsgMintMT) (*MsgMintMTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintMT not implemented")
}
func (UnimplementedMsgServer) EditMT(context.Context, *MsgEditMT) (*MsgEditMTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMT not implemented")
}
func (UnimplementedMsgServer) TransferMT(context.Context, *MsgTransferMT) (*MsgTransferMTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferMT not implemented")
}
func (UnimplementedMsgServer) BurnMT(context.Context, *MsgBurnMT) (*MsgBurnMTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnMT not implemented")
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

func _Msg_IssueDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIssueDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IssueDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_IssueDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IssueDenom(ctx, req.(*MsgIssueDenom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferDenom(ctx, req.(*MsgTransferDenom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintMT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintMT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintMT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintMT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintMT(ctx, req.(*MsgMintMT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditMT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditMT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditMT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EditMT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditMT(ctx, req.(*MsgEditMT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferMT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferMT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferMT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferMT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferMT(ctx, req.(*MsgTransferMT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnMT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnMT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnMT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BurnMT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnMT(ctx, req.(*MsgBurnMT))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.mt.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueDenom",
			Handler:    _Msg_IssueDenom_Handler,
		},
		{
			MethodName: "TransferDenom",
			Handler:    _Msg_TransferDenom_Handler,
		},
		{
			MethodName: "MintMT",
			Handler:    _Msg_MintMT_Handler,
		},
		{
			MethodName: "EditMT",
			Handler:    _Msg_EditMT_Handler,
		},
		{
			MethodName: "TransferMT",
			Handler:    _Msg_TransferMT_Handler,
		},
		{
			MethodName: "BurnMT",
			Handler:    _Msg_BurnMT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/mt/tx.proto",
}