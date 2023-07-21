// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: irismod/htlc/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgCreateHTLC defines a message to create an HTLC
type MsgCreateHTLC struct {
	Sender               string                                   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	To                   string                                   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	ReceiverOnOtherChain string                                   `protobuf:"bytes,3,opt,name=receiver_on_other_chain,json=receiverOnOtherChain,proto3" json:"receiver_on_other_chain,omitempty" yaml:"receiver_on_other_chain"`
	SenderOnOtherChain   string                                   `protobuf:"bytes,4,opt,name=sender_on_other_chain,json=senderOnOtherChain,proto3" json:"sender_on_other_chain,omitempty" yaml:"sender_on_other_chain"`
	Amount               github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	HashLock             string                                   `protobuf:"bytes,6,opt,name=hash_lock,json=hashLock,proto3" json:"hash_lock,omitempty" yaml:"hash_lock"`
	Timestamp            uint64                                   `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TimeLock             uint64                                   `protobuf:"varint,8,opt,name=time_lock,json=timeLock,proto3" json:"time_lock,omitempty" yaml:"time_lock"`
	Transfer             bool                                     `protobuf:"varint,9,opt,name=transfer,proto3" json:"transfer,omitempty"`
}

func (m *MsgCreateHTLC) Reset()         { *m = MsgCreateHTLC{} }
func (m *MsgCreateHTLC) String() string { return proto.CompactTextString(m) }
func (*MsgCreateHTLC) ProtoMessage()    {}
func (*MsgCreateHTLC) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef42cbbbdd4c733d, []int{0}
}
func (m *MsgCreateHTLC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateHTLC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateHTLC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateHTLC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateHTLC.Merge(m, src)
}
func (m *MsgCreateHTLC) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateHTLC) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateHTLC.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateHTLC proto.InternalMessageInfo

// MsgCreateHTLCResponse defines the Msg/CreateHTLC response type
type MsgCreateHTLCResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateHTLCResponse) Reset()         { *m = MsgCreateHTLCResponse{} }
func (m *MsgCreateHTLCResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateHTLCResponse) ProtoMessage()    {}
func (*MsgCreateHTLCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef42cbbbdd4c733d, []int{1}
}
func (m *MsgCreateHTLCResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateHTLCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateHTLCResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateHTLCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateHTLCResponse.Merge(m, src)
}
func (m *MsgCreateHTLCResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateHTLCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateHTLCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateHTLCResponse proto.InternalMessageInfo

// MsgClaimHTLC defines a message to claim an HTLC
type MsgClaimHTLC struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Secret string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *MsgClaimHTLC) Reset()         { *m = MsgClaimHTLC{} }
func (m *MsgClaimHTLC) String() string { return proto.CompactTextString(m) }
func (*MsgClaimHTLC) ProtoMessage()    {}
func (*MsgClaimHTLC) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef42cbbbdd4c733d, []int{2}
}
func (m *MsgClaimHTLC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimHTLC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimHTLC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimHTLC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimHTLC.Merge(m, src)
}
func (m *MsgClaimHTLC) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimHTLC) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimHTLC.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimHTLC proto.InternalMessageInfo

// MsgClaimHTLCResponse defines the Msg/ClaimHTLC response type
type MsgClaimHTLCResponse struct {
}

func (m *MsgClaimHTLCResponse) Reset()         { *m = MsgClaimHTLCResponse{} }
func (m *MsgClaimHTLCResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimHTLCResponse) ProtoMessage()    {}
func (*MsgClaimHTLCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef42cbbbdd4c733d, []int{3}
}
func (m *MsgClaimHTLCResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimHTLCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimHTLCResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimHTLCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimHTLCResponse.Merge(m, src)
}
func (m *MsgClaimHTLCResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimHTLCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimHTLCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimHTLCResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateHTLC)(nil), "irismod.htlc.MsgCreateHTLC")
	proto.RegisterType((*MsgCreateHTLCResponse)(nil), "irismod.htlc.MsgCreateHTLCResponse")
	proto.RegisterType((*MsgClaimHTLC)(nil), "irismod.htlc.MsgClaimHTLC")
	proto.RegisterType((*MsgClaimHTLCResponse)(nil), "irismod.htlc.MsgClaimHTLCResponse")
}

func init() { proto.RegisterFile("irismod/htlc/tx.proto", fileDescriptor_ef42cbbbdd4c733d) }

var fileDescriptor_ef42cbbbdd4c733d = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x31, 0x73, 0xd3, 0x30,
	0x14, 0x8e, 0x93, 0x10, 0x62, 0xd1, 0x72, 0x9c, 0x2f, 0x29, 0xc6, 0x14, 0x3b, 0x27, 0x06, 0xb2,
	0xd4, 0x26, 0x65, 0xeb, 0x98, 0x2c, 0xdc, 0xb5, 0xa5, 0x77, 0x86, 0x05, 0x96, 0x9c, 0x22, 0x8b,
	0x58, 0x97, 0x58, 0xca, 0x49, 0x4a, 0x8f, 0xfe, 0x0b, 0x7e, 0x02, 0xc7, 0xc8, 0x2f, 0xc9, 0xd8,
	0x91, 0x29, 0x40, 0xb2, 0x30, 0xe7, 0x17, 0x70, 0xb2, 0x9c, 0x34, 0x01, 0xda, 0xc9, 0xef, 0x7d,
	0xdf, 0x7b, 0x9f, 0x9f, 0x9e, 0x3e, 0x81, 0x26, 0x15, 0x54, 0x66, 0x3c, 0x89, 0x52, 0x35, 0xc6,
	0x91, 0xfa, 0x14, 0x4e, 0x04, 0x57, 0xdc, 0xd9, 0x2b, 0xe0, 0x50, 0xc3, 0x9e, 0x8f, 0xb9, 0xcc,
	0xb8, 0x8c, 0x06, 0x48, 0x92, 0xe8, 0xb2, 0x33, 0x20, 0x0a, 0x75, 0x22, 0xcc, 0x29, 0x33, 0xd5,
	0x5e, 0x63, 0xc8, 0x87, 0x3c, 0x0f, 0x23, 0x1d, 0x19, 0x14, 0xae, 0x2a, 0x60, 0xff, 0x5c, 0x0e,
	0x7b, 0x82, 0x20, 0x45, 0x5e, 0xbf, 0x3b, 0xeb, 0x39, 0x07, 0xa0, 0x26, 0x09, 0x4b, 0x88, 0x70,
	0xad, 0x96, 0xd5, 0xb6, 0xe3, 0x22, 0x73, 0x1e, 0x82, 0xb2, 0xe2, 0x6e, 0x39, 0xc7, 0xca, 0x8a,
	0x3b, 0xef, 0xc1, 0x63, 0x41, 0x30, 0xa1, 0x97, 0x44, 0xf4, 0x39, 0xeb, 0x73, 0x95, 0x12, 0xd1,
	0xc7, 0x29, 0xa2, 0xcc, 0xad, 0xe8, 0xa2, 0x2e, 0x5c, 0xcd, 0x03, 0xff, 0x0a, 0x65, 0xe3, 0x13,
	0x78, 0x4b, 0x21, 0x8c, 0x1b, 0x6b, 0xe6, 0x82, 0x5d, 0x68, 0xbc, 0xa7, 0x61, 0xe7, 0x2d, 0x68,
	0x9a, 0x9f, 0xfe, 0x2d, 0x5c, 0xcd, 0x85, 0x5b, 0xab, 0x79, 0x70, 0x68, 0x84, 0xff, 0x5b, 0x06,
	0x63, 0xc7, 0xe0, 0x3b, 0xa2, 0x18, 0xd4, 0x50, 0xc6, 0xa7, 0x4c, 0xb9, 0xf7, 0x5a, 0x95, 0xf6,
	0x83, 0xe3, 0x27, 0xa1, 0x59, 0x58, 0xa8, 0x17, 0x16, 0x16, 0x0b, 0x0b, 0x7b, 0x9c, 0xb2, 0xee,
	0xcb, 0xd9, 0x3c, 0x28, 0x7d, 0xfb, 0x11, 0xb4, 0x87, 0x54, 0xa5, 0xd3, 0x41, 0x88, 0x79, 0x16,
	0x15, 0xdb, 0x35, 0x9f, 0x23, 0x99, 0x8c, 0x22, 0x75, 0x35, 0x21, 0x32, 0x6f, 0x90, 0x71, 0x21,
	0xed, 0x74, 0x80, 0x9d, 0x22, 0x99, 0xf6, 0xc7, 0x1c, 0x8f, 0xdc, 0x5a, 0x3e, 0x6d, 0x63, 0x35,
	0x0f, 0x1e, 0x99, 0x69, 0x37, 0x14, 0x8c, 0xeb, 0x3a, 0x3e, 0xe3, 0x78, 0xe4, 0x1c, 0x02, 0x5b,
	0xd1, 0x8c, 0x48, 0x85, 0xb2, 0x89, 0x7b, 0xbf, 0x65, 0xb5, 0xab, 0xf1, 0x0d, 0xa0, 0x05, 0x75,
	0x62, 0x04, 0xeb, 0x9a, 0xdd, 0x16, 0xdc, 0x50, 0x30, 0xae, 0xeb, 0x38, 0x17, 0xf4, 0x40, 0x5d,
	0x09, 0xc4, 0xe4, 0x47, 0x22, 0x5c, 0xbb, 0x65, 0xb5, 0xeb, 0xf1, 0x26, 0x3f, 0xa9, 0xfe, 0xfe,
	0x12, 0x58, 0xf0, 0x05, 0x68, 0xee, 0xdc, 0x79, 0x4c, 0xe4, 0x84, 0x33, 0x49, 0xf4, 0x1d, 0xd3,
	0xa4, 0xb8, 0xf7, 0x32, 0x4d, 0x20, 0x06, 0x7b, 0xba, 0x70, 0x8c, 0x68, 0x76, 0xa7, 0x37, 0x9e,
	0xe5, 0x7d, 0xb9, 0x37, 0xba, 0xfb, 0xab, 0x79, 0x60, 0x9b, 0xf1, 0x68, 0x02, 0xb5, 0x8c, 0x69,
	0xc3, 0x82, 0x28, 0xe3, 0x8c, 0xb8, 0xc8, 0x8a, 0x69, 0x0e, 0x40, 0x63, 0xfb, 0x27, 0xeb, 0x61,
	0x8e, 0xbf, 0x5a, 0xa0, 0x72, 0x2e, 0x87, 0xce, 0x1b, 0x00, 0xb6, 0xec, 0xf9, 0x34, 0xdc, 0x76,
	0x7d, 0xb8, 0x73, 0x0e, 0xef, 0xf9, 0x1d, 0xe4, 0xe6, 0x90, 0xa7, 0xc0, 0xbe, 0x39, 0x91, 0xf7,
	0x6f, 0xc7, 0x9a, 0xf3, 0xe0, 0xed, 0xdc, 0x5a, 0xac, 0x7b, 0x3a, 0xfb, 0xe5, 0x97, 0x66, 0x0b,
	0xdf, 0xba, 0x5e, 0xf8, 0xd6, 0xcf, 0x85, 0x6f, 0x7d, 0x5e, 0xfa, 0xa5, 0xeb, 0xa5, 0x5f, 0xfa,
	0xbe, 0xf4, 0x4b, 0x1f, 0x8e, 0xb6, 0x0c, 0xa4, 0xb5, 0x18, 0x51, 0xd1, 0xfa, 0x2d, 0x67, 0x3c,
	0x99, 0x8e, 0x89, 0x2c, 0xde, 0xb4, 0xf6, 0xd2, 0xa0, 0x96, 0xbf, 0xc9, 0x57, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x06, 0x16, 0xdf, 0xc4, 0xf0, 0x03, 0x00, 0x00,
}

func (this *MsgCreateHTLC) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgCreateHTLC)
	if !ok {
		that2, ok := that.(MsgCreateHTLC)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Sender != that1.Sender {
		return false
	}
	if this.To != that1.To {
		return false
	}
	if this.ReceiverOnOtherChain != that1.ReceiverOnOtherChain {
		return false
	}
	if this.SenderOnOtherChain != that1.SenderOnOtherChain {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(&that1.Amount[i]) {
			return false
		}
	}
	if this.HashLock != that1.HashLock {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if this.TimeLock != that1.TimeLock {
		return false
	}
	if this.Transfer != that1.Transfer {
		return false
	}
	return true
}
func (this *MsgClaimHTLC) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgClaimHTLC)
	if !ok {
		that2, ok := that.(MsgClaimHTLC)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Sender != that1.Sender {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Secret != that1.Secret {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateHTLC defines a method for creating a HTLC
	CreateHTLC(ctx context.Context, in *MsgCreateHTLC, opts ...grpc.CallOption) (*MsgCreateHTLCResponse, error)
	// ClaimHTLC defines a method for claiming a HTLC
	ClaimHTLC(ctx context.Context, in *MsgClaimHTLC, opts ...grpc.CallOption) (*MsgClaimHTLCResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateHTLC(ctx context.Context, in *MsgCreateHTLC, opts ...grpc.CallOption) (*MsgCreateHTLCResponse, error) {
	out := new(MsgCreateHTLCResponse)
	err := c.cc.Invoke(ctx, "/irismod.htlc.Msg/CreateHTLC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimHTLC(ctx context.Context, in *MsgClaimHTLC, opts ...grpc.CallOption) (*MsgClaimHTLCResponse, error) {
	out := new(MsgClaimHTLCResponse)
	err := c.cc.Invoke(ctx, "/irismod.htlc.Msg/ClaimHTLC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateHTLC defines a method for creating a HTLC
	CreateHTLC(context.Context, *MsgCreateHTLC) (*MsgCreateHTLCResponse, error)
	// ClaimHTLC defines a method for claiming a HTLC
	ClaimHTLC(context.Context, *MsgClaimHTLC) (*MsgClaimHTLCResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateHTLC(ctx context.Context, req *MsgCreateHTLC) (*MsgCreateHTLCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTLC not implemented")
}
func (*UnimplementedMsgServer) ClaimHTLC(ctx context.Context, req *MsgClaimHTLC) (*MsgClaimHTLCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimHTLC not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateHTLC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateHTLC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateHTLC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.htlc.Msg/CreateHTLC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateHTLC(ctx, req.(*MsgCreateHTLC))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimHTLC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimHTLC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimHTLC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.htlc.Msg/ClaimHTLC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimHTLC(ctx, req.(*MsgClaimHTLC))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.htlc.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHTLC",
			Handler:    _Msg_CreateHTLC_Handler,
		},
		{
			MethodName: "ClaimHTLC",
			Handler:    _Msg_ClaimHTLC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/htlc/tx.proto",
}

func (m *MsgCreateHTLC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateHTLC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateHTLC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Transfer {
		i--
		if m.Transfer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.TimeLock != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeLock))
		i--
		dAtA[i] = 0x40
	}
	if m.Timestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x38
	}
	if len(m.HashLock) > 0 {
		i -= len(m.HashLock)
		copy(dAtA[i:], m.HashLock)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HashLock)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.SenderOnOtherChain) > 0 {
		i -= len(m.SenderOnOtherChain)
		copy(dAtA[i:], m.SenderOnOtherChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SenderOnOtherChain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ReceiverOnOtherChain) > 0 {
		i -= len(m.ReceiverOnOtherChain)
		copy(dAtA[i:], m.ReceiverOnOtherChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ReceiverOnOtherChain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateHTLCResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateHTLCResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateHTLCResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimHTLC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimHTLC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimHTLC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimHTLCResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimHTLCResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimHTLCResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateHTLC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ReceiverOnOtherChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SenderOnOtherChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.HashLock)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTx(uint64(m.Timestamp))
	}
	if m.TimeLock != 0 {
		n += 1 + sovTx(uint64(m.TimeLock))
	}
	if m.Transfer {
		n += 2
	}
	return n
}

func (m *MsgCreateHTLCResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimHTLC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgClaimHTLCResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateHTLC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateHTLC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateHTLC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverOnOtherChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiverOnOtherChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderOnOtherChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderOnOtherChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashLock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashLock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeLock", wireType)
			}
			m.TimeLock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeLock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transfer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Transfer = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateHTLCResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateHTLCResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateHTLCResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgClaimHTLC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgClaimHTLC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimHTLC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgClaimHTLCResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgClaimHTLCResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimHTLCResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
