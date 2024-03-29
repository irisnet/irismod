// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/oracle/query.proto

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
	Query_Feed_FullMethodName      = "/irismod.oracle.Query/Feed"
	Query_Feeds_FullMethodName     = "/irismod.oracle.Query/Feeds"
	Query_FeedValue_FullMethodName = "/irismod.oracle.Query/FeedValue"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Feed queries the feed
	Feed(ctx context.Context, in *QueryFeedRequest, opts ...grpc.CallOption) (*QueryFeedResponse, error)
	// Feeds queries the feed list
	Feeds(ctx context.Context, in *QueryFeedsRequest, opts ...grpc.CallOption) (*QueryFeedsResponse, error)
	// FeedValue queries the feed value
	FeedValue(ctx context.Context, in *QueryFeedValueRequest, opts ...grpc.CallOption) (*QueryFeedValueResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Feed(ctx context.Context, in *QueryFeedRequest, opts ...grpc.CallOption) (*QueryFeedResponse, error) {
	out := new(QueryFeedResponse)
	err := c.cc.Invoke(ctx, Query_Feed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Feeds(ctx context.Context, in *QueryFeedsRequest, opts ...grpc.CallOption) (*QueryFeedsResponse, error) {
	out := new(QueryFeedsResponse)
	err := c.cc.Invoke(ctx, Query_Feeds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FeedValue(ctx context.Context, in *QueryFeedValueRequest, opts ...grpc.CallOption) (*QueryFeedValueResponse, error) {
	out := new(QueryFeedValueResponse)
	err := c.cc.Invoke(ctx, Query_FeedValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Feed queries the feed
	Feed(context.Context, *QueryFeedRequest) (*QueryFeedResponse, error)
	// Feeds queries the feed list
	Feeds(context.Context, *QueryFeedsRequest) (*QueryFeedsResponse, error)
	// FeedValue queries the feed value
	FeedValue(context.Context, *QueryFeedValueRequest) (*QueryFeedValueResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Feed(context.Context, *QueryFeedRequest) (*QueryFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedQueryServer) Feeds(context.Context, *QueryFeedsRequest) (*QueryFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feeds not implemented")
}
func (UnimplementedQueryServer) FeedValue(context.Context, *QueryFeedValueRequest) (*QueryFeedValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedValue not implemented")
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

func _Query_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Feed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Feed(ctx, req.(*QueryFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Feeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Feeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Feeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Feeds(ctx, req.(*QueryFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FeedValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeedValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeedValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_FeedValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeedValue(ctx, req.(*QueryFeedValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.oracle.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Feed",
			Handler:    _Query_Feed_Handler,
		},
		{
			MethodName: "Feeds",
			Handler:    _Query_Feeds_Handler,
		},
		{
			MethodName: "FeedValue",
			Handler:    _Query_FeedValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/oracle/query.proto",
}
