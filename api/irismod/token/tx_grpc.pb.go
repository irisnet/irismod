// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/token/tx.proto

package token

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
	Msg_IssueToken_FullMethodName         = "/irismod.token.Msg/IssueToken"
	Msg_EditToken_FullMethodName          = "/irismod.token.Msg/EditToken"
	Msg_MintToken_FullMethodName          = "/irismod.token.Msg/MintToken"
	Msg_BurnToken_FullMethodName          = "/irismod.token.Msg/BurnToken"
	Msg_TransferTokenOwner_FullMethodName = "/irismod.token.Msg/TransferTokenOwner"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// IssueToken defines a method for issuing a new token
	IssueToken(ctx context.Context, in *MsgIssueToken, opts ...grpc.CallOption) (*MsgIssueTokenResponse, error)
	// EditToken defines a method for editing a token
	EditToken(ctx context.Context, in *MsgEditToken, opts ...grpc.CallOption) (*MsgEditTokenResponse, error)
	// MintToken defines a method for minting some tokens
	MintToken(ctx context.Context, in *MsgMintToken, opts ...grpc.CallOption) (*MsgMintTokenResponse, error)
	// BurnToken defines a method for burning some tokens
	BurnToken(ctx context.Context, in *MsgBurnToken, opts ...grpc.CallOption) (*MsgBurnTokenResponse, error)
	// TransferTokenOwner defines a method for minting some tokens
	TransferTokenOwner(ctx context.Context, in *MsgTransferTokenOwner, opts ...grpc.CallOption) (*MsgTransferTokenOwnerResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IssueToken(ctx context.Context, in *MsgIssueToken, opts ...grpc.CallOption) (*MsgIssueTokenResponse, error) {
	out := new(MsgIssueTokenResponse)
	err := c.cc.Invoke(ctx, Msg_IssueToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditToken(ctx context.Context, in *MsgEditToken, opts ...grpc.CallOption) (*MsgEditTokenResponse, error) {
	out := new(MsgEditTokenResponse)
	err := c.cc.Invoke(ctx, Msg_EditToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintToken(ctx context.Context, in *MsgMintToken, opts ...grpc.CallOption) (*MsgMintTokenResponse, error) {
	out := new(MsgMintTokenResponse)
	err := c.cc.Invoke(ctx, Msg_MintToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnToken(ctx context.Context, in *MsgBurnToken, opts ...grpc.CallOption) (*MsgBurnTokenResponse, error) {
	out := new(MsgBurnTokenResponse)
	err := c.cc.Invoke(ctx, Msg_BurnToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferTokenOwner(ctx context.Context, in *MsgTransferTokenOwner, opts ...grpc.CallOption) (*MsgTransferTokenOwnerResponse, error) {
	out := new(MsgTransferTokenOwnerResponse)
	err := c.cc.Invoke(ctx, Msg_TransferTokenOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// IssueToken defines a method for issuing a new token
	IssueToken(context.Context, *MsgIssueToken) (*MsgIssueTokenResponse, error)
	// EditToken defines a method for editing a token
	EditToken(context.Context, *MsgEditToken) (*MsgEditTokenResponse, error)
	// MintToken defines a method for minting some tokens
	MintToken(context.Context, *MsgMintToken) (*MsgMintTokenResponse, error)
	// BurnToken defines a method for burning some tokens
	BurnToken(context.Context, *MsgBurnToken) (*MsgBurnTokenResponse, error)
	// TransferTokenOwner defines a method for minting some tokens
	TransferTokenOwner(context.Context, *MsgTransferTokenOwner) (*MsgTransferTokenOwnerResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) IssueToken(context.Context, *MsgIssueToken) (*MsgIssueTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (UnimplementedMsgServer) EditToken(context.Context, *MsgEditToken) (*MsgEditTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditToken not implemented")
}
func (UnimplementedMsgServer) MintToken(context.Context, *MsgMintToken) (*MsgMintTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintToken not implemented")
}
func (UnimplementedMsgServer) BurnToken(context.Context, *MsgBurnToken) (*MsgBurnTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnToken not implemented")
}
func (UnimplementedMsgServer) TransferTokenOwner(context.Context, *MsgTransferTokenOwner) (*MsgTransferTokenOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferTokenOwner not implemented")
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

func _Msg_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIssueToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_IssueToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IssueToken(ctx, req.(*MsgIssueToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EditToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditToken(ctx, req.(*MsgEditToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintToken(ctx, req.(*MsgMintToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BurnToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnToken(ctx, req.(*MsgBurnToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferTokenOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferTokenOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferTokenOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferTokenOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferTokenOwner(ctx, req.(*MsgTransferTokenOwner))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.token.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueToken",
			Handler:    _Msg_IssueToken_Handler,
		},
		{
			MethodName: "EditToken",
			Handler:    _Msg_EditToken_Handler,
		},
		{
			MethodName: "MintToken",
			Handler:    _Msg_MintToken_Handler,
		},
		{
			MethodName: "BurnToken",
			Handler:    _Msg_BurnToken_Handler,
		},
		{
			MethodName: "TransferTokenOwner",
			Handler:    _Msg_TransferTokenOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/token/tx.proto",
}
