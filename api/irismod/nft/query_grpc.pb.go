// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/nft/query.proto

package nft

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
	Query_Supply_FullMethodName     = "/irismod.nft.Query/Supply"
	Query_Owner_FullMethodName      = "/irismod.nft.Query/Owner"
	Query_Collection_FullMethodName = "/irismod.nft.Query/Collection"
	Query_Denom_FullMethodName      = "/irismod.nft.Query/Denom"
	Query_Denoms_FullMethodName     = "/irismod.nft.Query/Denoms"
	Query_NFT_FullMethodName        = "/irismod.nft.Query/NFT"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Supply queries the total supply of a given denom or owner
	Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error)
	// Owner queries the NFTs of the specified owner
	Owner(ctx context.Context, in *QueryOwnerRequest, opts ...grpc.CallOption) (*QueryOwnerResponse, error)
	// Collection queries the NFTs of the specified denom
	Collection(ctx context.Context, in *QueryCollectionRequest, opts ...grpc.CallOption) (*QueryCollectionResponse, error)
	// Denom queries the definition of a given denom
	Denom(ctx context.Context, in *QueryDenomRequest, opts ...grpc.CallOption) (*QueryDenomResponse, error)
	// Denoms queries all the denoms
	Denoms(ctx context.Context, in *QueryDenomsRequest, opts ...grpc.CallOption) (*QueryDenomsResponse, error)
	// NFT queries the NFT for the given denom and token ID
	NFT(ctx context.Context, in *QueryNFTRequest, opts ...grpc.CallOption) (*QueryNFTResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error) {
	out := new(QuerySupplyResponse)
	err := c.cc.Invoke(ctx, Query_Supply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Owner(ctx context.Context, in *QueryOwnerRequest, opts ...grpc.CallOption) (*QueryOwnerResponse, error) {
	out := new(QueryOwnerResponse)
	err := c.cc.Invoke(ctx, Query_Owner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Collection(ctx context.Context, in *QueryCollectionRequest, opts ...grpc.CallOption) (*QueryCollectionResponse, error) {
	out := new(QueryCollectionResponse)
	err := c.cc.Invoke(ctx, Query_Collection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Denom(ctx context.Context, in *QueryDenomRequest, opts ...grpc.CallOption) (*QueryDenomResponse, error) {
	out := new(QueryDenomResponse)
	err := c.cc.Invoke(ctx, Query_Denom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Denoms(ctx context.Context, in *QueryDenomsRequest, opts ...grpc.CallOption) (*QueryDenomsResponse, error) {
	out := new(QueryDenomsResponse)
	err := c.cc.Invoke(ctx, Query_Denoms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NFT(ctx context.Context, in *QueryNFTRequest, opts ...grpc.CallOption) (*QueryNFTResponse, error) {
	out := new(QueryNFTResponse)
	err := c.cc.Invoke(ctx, Query_NFT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Supply queries the total supply of a given denom or owner
	Supply(context.Context, *QuerySupplyRequest) (*QuerySupplyResponse, error)
	// Owner queries the NFTs of the specified owner
	Owner(context.Context, *QueryOwnerRequest) (*QueryOwnerResponse, error)
	// Collection queries the NFTs of the specified denom
	Collection(context.Context, *QueryCollectionRequest) (*QueryCollectionResponse, error)
	// Denom queries the definition of a given denom
	Denom(context.Context, *QueryDenomRequest) (*QueryDenomResponse, error)
	// Denoms queries all the denoms
	Denoms(context.Context, *QueryDenomsRequest) (*QueryDenomsResponse, error)
	// NFT queries the NFT for the given denom and token ID
	NFT(context.Context, *QueryNFTRequest) (*QueryNFTResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Supply(context.Context, *QuerySupplyRequest) (*QuerySupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
}
func (UnimplementedQueryServer) Owner(context.Context, *QueryOwnerRequest) (*QueryOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Owner not implemented")
}
func (UnimplementedQueryServer) Collection(context.Context, *QueryCollectionRequest) (*QueryCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collection not implemented")
}
func (UnimplementedQueryServer) Denom(context.Context, *QueryDenomRequest) (*QueryDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Denom not implemented")
}
func (UnimplementedQueryServer) Denoms(context.Context, *QueryDenomsRequest) (*QueryDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Denoms not implemented")
}
func (UnimplementedQueryServer) NFT(context.Context, *QueryNFTRequest) (*QueryNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NFT not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Supply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Supply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Supply(ctx, req.(*QuerySupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Owner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Owner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Owner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Owner(ctx, req.(*QueryOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Collection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Collection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Collection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Collection(ctx, req.(*QueryCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Denom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Denom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Denom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Denom(ctx, req.(*QueryDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Denoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Denoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Denoms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Denoms(ctx, req.(*QueryDenomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNFTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_NFT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NFT(ctx, req.(*QueryNFTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.nft.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Supply",
			Handler:    _Query_Supply_Handler,
		},
		{
			MethodName: "Owner",
			Handler:    _Query_Owner_Handler,
		},
		{
			MethodName: "Collection",
			Handler:    _Query_Collection_Handler,
		},
		{
			MethodName: "Denom",
			Handler:    _Query_Denom_Handler,
		},
		{
			MethodName: "Denoms",
			Handler:    _Query_Denoms_Handler,
		},
		{
			MethodName: "NFT",
			Handler:    _Query_NFT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/nft/query.proto",
}
