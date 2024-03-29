// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/farm/tx.proto

package farm

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
	Msg_CreatePool_FullMethodName                  = "/irismod.farm.Msg/CreatePool"
	Msg_CreatePoolWithCommunityPool_FullMethodName = "/irismod.farm.Msg/CreatePoolWithCommunityPool"
	Msg_DestroyPool_FullMethodName                 = "/irismod.farm.Msg/DestroyPool"
	Msg_AdjustPool_FullMethodName                  = "/irismod.farm.Msg/AdjustPool"
	Msg_Stake_FullMethodName                       = "/irismod.farm.Msg/Stake"
	Msg_Unstake_FullMethodName                     = "/irismod.farm.Msg/Unstake"
	Msg_Harvest_FullMethodName                     = "/irismod.farm.Msg/Harvest"
	Msg_UpdateParams_FullMethodName                = "/irismod.farm.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreatePool defines a method for creating a new farm pool
	CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
	// CreatePoolWithCommunityPool defines a method for creating a new farm pool
	CreatePoolWithCommunityPool(ctx context.Context, in *MsgCreatePoolWithCommunityPool, opts ...grpc.CallOption) (*MsgCreatePoolWithCommunityPoolResponse, error)
	// DestroyPool defines a method for destroying a existed farm pool
	DestroyPool(ctx context.Context, in *MsgDestroyPool, opts ...grpc.CallOption) (*MsgDestroyPoolResponse, error)
	// AdjustPool defines a method for adjusting the farm pool params
	AdjustPool(ctx context.Context, in *MsgAdjustPool, opts ...grpc.CallOption) (*MsgAdjustPoolResponse, error)
	// Stake defines a method for staking some lp token to a farm pool
	Stake(ctx context.Context, in *MsgStake, opts ...grpc.CallOption) (*MsgStakeResponse, error)
	// Unstake defines a method for unstaking some lp token from a farm pool and
	// withdraw some reward
	Unstake(ctx context.Context, in *MsgUnstake, opts ...grpc.CallOption) (*MsgUnstakeResponse, error)
	// Harvest defines a method withdraw some reward from a farm pool
	Harvest(ctx context.Context, in *MsgHarvest, opts ...grpc.CallOption) (*MsgHarvestResponse, error)
	// UpdateParams defines a governance operation for updating the x/coinswap
	// module parameters. The authority is defined in the keeper.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePoolWithCommunityPool(ctx context.Context, in *MsgCreatePoolWithCommunityPool, opts ...grpc.CallOption) (*MsgCreatePoolWithCommunityPoolResponse, error) {
	out := new(MsgCreatePoolWithCommunityPoolResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePoolWithCommunityPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DestroyPool(ctx context.Context, in *MsgDestroyPool, opts ...grpc.CallOption) (*MsgDestroyPoolResponse, error) {
	out := new(MsgDestroyPoolResponse)
	err := c.cc.Invoke(ctx, Msg_DestroyPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AdjustPool(ctx context.Context, in *MsgAdjustPool, opts ...grpc.CallOption) (*MsgAdjustPoolResponse, error) {
	out := new(MsgAdjustPoolResponse)
	err := c.cc.Invoke(ctx, Msg_AdjustPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Stake(ctx context.Context, in *MsgStake, opts ...grpc.CallOption) (*MsgStakeResponse, error) {
	out := new(MsgStakeResponse)
	err := c.cc.Invoke(ctx, Msg_Stake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Unstake(ctx context.Context, in *MsgUnstake, opts ...grpc.CallOption) (*MsgUnstakeResponse, error) {
	out := new(MsgUnstakeResponse)
	err := c.cc.Invoke(ctx, Msg_Unstake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Harvest(ctx context.Context, in *MsgHarvest, opts ...grpc.CallOption) (*MsgHarvestResponse, error) {
	out := new(MsgHarvestResponse)
	err := c.cc.Invoke(ctx, Msg_Harvest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreatePool defines a method for creating a new farm pool
	CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error)
	// CreatePoolWithCommunityPool defines a method for creating a new farm pool
	CreatePoolWithCommunityPool(context.Context, *MsgCreatePoolWithCommunityPool) (*MsgCreatePoolWithCommunityPoolResponse, error)
	// DestroyPool defines a method for destroying a existed farm pool
	DestroyPool(context.Context, *MsgDestroyPool) (*MsgDestroyPoolResponse, error)
	// AdjustPool defines a method for adjusting the farm pool params
	AdjustPool(context.Context, *MsgAdjustPool) (*MsgAdjustPoolResponse, error)
	// Stake defines a method for staking some lp token to a farm pool
	Stake(context.Context, *MsgStake) (*MsgStakeResponse, error)
	// Unstake defines a method for unstaking some lp token from a farm pool and
	// withdraw some reward
	Unstake(context.Context, *MsgUnstake) (*MsgUnstakeResponse, error)
	// Harvest defines a method withdraw some reward from a farm pool
	Harvest(context.Context, *MsgHarvest) (*MsgHarvestResponse, error)
	// UpdateParams defines a governance operation for updating the x/coinswap
	// module parameters. The authority is defined in the keeper.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedMsgServer) CreatePoolWithCommunityPool(context.Context, *MsgCreatePoolWithCommunityPool) (*MsgCreatePoolWithCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePoolWithCommunityPool not implemented")
}
func (UnimplementedMsgServer) DestroyPool(context.Context, *MsgDestroyPool) (*MsgDestroyPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyPool not implemented")
}
func (UnimplementedMsgServer) AdjustPool(context.Context, *MsgAdjustPool) (*MsgAdjustPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdjustPool not implemented")
}
func (UnimplementedMsgServer) Stake(context.Context, *MsgStake) (*MsgStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stake not implemented")
}
func (UnimplementedMsgServer) Unstake(context.Context, *MsgUnstake) (*MsgUnstakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unstake not implemented")
}
func (UnimplementedMsgServer) Harvest(context.Context, *MsgHarvest) (*MsgHarvestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Harvest not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
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

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*MsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePoolWithCommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePoolWithCommunityPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePoolWithCommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePoolWithCommunityPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePoolWithCommunityPool(ctx, req.(*MsgCreatePoolWithCommunityPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DestroyPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDestroyPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DestroyPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DestroyPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DestroyPool(ctx, req.(*MsgDestroyPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AdjustPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAdjustPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AdjustPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AdjustPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AdjustPool(ctx, req.(*MsgAdjustPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Stake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Stake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Stake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Stake(ctx, req.(*MsgStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Unstake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnstake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unstake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Unstake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unstake(ctx, req.(*MsgUnstake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Harvest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgHarvest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Harvest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Harvest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Harvest(ctx, req.(*MsgHarvest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.farm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
		{
			MethodName: "CreatePoolWithCommunityPool",
			Handler:    _Msg_CreatePoolWithCommunityPool_Handler,
		},
		{
			MethodName: "DestroyPool",
			Handler:    _Msg_DestroyPool_Handler,
		},
		{
			MethodName: "AdjustPool",
			Handler:    _Msg_AdjustPool_Handler,
		},
		{
			MethodName: "Stake",
			Handler:    _Msg_Stake_Handler,
		},
		{
			MethodName: "Unstake",
			Handler:    _Msg_Unstake_Handler,
		},
		{
			MethodName: "Harvest",
			Handler:    _Msg_Harvest_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/farm/tx.proto",
}
