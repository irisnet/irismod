// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: irismod/random/tx.proto

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

// MsgRequestRandom defines an sdk.Msg type that supports requesting a random
// number
type MsgRequestRandom struct {
	BlockInterval uint64                                   `protobuf:"varint,1,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty" yaml:"block_interval"`
	Consumer      string                                   `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Oracle        bool                                     `protobuf:"varint,3,opt,name=oracle,proto3" json:"oracle,omitempty"`
	ServiceFeeCap github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=service_fee_cap,json=serviceFeeCap,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"service_fee_cap" yaml:"service_fee_cap"`
}

func (m *MsgRequestRandom) Reset()         { *m = MsgRequestRandom{} }
func (m *MsgRequestRandom) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRandom) ProtoMessage()    {}
func (*MsgRequestRandom) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac3441cc610dfbbb, []int{0}
}
func (m *MsgRequestRandom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRandom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRandom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRandom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRandom.Merge(m, src)
}
func (m *MsgRequestRandom) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRandom) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRandom.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRandom proto.InternalMessageInfo

func (m *MsgRequestRandom) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

func (m *MsgRequestRandom) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *MsgRequestRandom) GetOracle() bool {
	if m != nil {
		return m.Oracle
	}
	return false
}

func (m *MsgRequestRandom) GetServiceFeeCap() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ServiceFeeCap
	}
	return nil
}

// MsgRequestRandomResponse defines the Msg/RequestRandom response type
type MsgRequestRandomResponse struct {
}

func (m *MsgRequestRandomResponse) Reset()         { *m = MsgRequestRandomResponse{} }
func (m *MsgRequestRandomResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRandomResponse) ProtoMessage()    {}
func (*MsgRequestRandomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac3441cc610dfbbb, []int{1}
}
func (m *MsgRequestRandomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRandomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRandomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRandomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRandomResponse.Merge(m, src)
}
func (m *MsgRequestRandomResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRandomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRandomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRandomResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequestRandom)(nil), "irismod.random.MsgRequestRandom")
	proto.RegisterType((*MsgRequestRandomResponse)(nil), "irismod.random.MsgRequestRandomResponse")
}

func init() { proto.RegisterFile("irismod/random/tx.proto", fileDescriptor_ac3441cc610dfbbb) }

var fileDescriptor_ac3441cc610dfbbb = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x5d, 0x77, 0xab, 0xaa, 0x18, 0x6d, 0x41, 0x11, 0x94, 0x34, 0x87, 0x6c, 0x14, 0x09, 0x29,
	0x17, 0x6c, 0xb5, 0xdc, 0xb8, 0x80, 0xb6, 0x12, 0x12, 0x48, 0xbd, 0xe4, 0x08, 0x87, 0x95, 0xe3,
	0x0c, 0x21, 0x6a, 0x9c, 0x09, 0x1e, 0x67, 0x45, 0x3f, 0x02, 0x89, 0x23, 0xdf, 0xc0, 0x97, 0xf4,
	0xd8, 0x23, 0xa7, 0x05, 0xed, 0xfe, 0x41, 0xbf, 0x00, 0x6d, 0x12, 0x56, 0xda, 0xbd, 0x70, 0xb2,
	0x9f, 0xdf, 0xcc, 0xb3, 0xdf, 0x1b, 0xf3, 0x67, 0xa5, 0x2d, 0xc9, 0x60, 0x2e, 0xad, 0xaa, 0x73,
	0x34, 0xd2, 0x7d, 0x15, 0x8d, 0x45, 0x87, 0xde, 0xc9, 0x40, 0x88, 0x9e, 0x08, 0x9e, 0x14, 0x58,
	0x60, 0x47, 0xc9, 0xcd, 0xae, 0xaf, 0x0a, 0x42, 0x8d, 0x64, 0x90, 0x64, 0xa6, 0x08, 0xe4, 0xe2,
	0x3c, 0x03, 0xa7, 0xce, 0xa5, 0xc6, 0xb2, 0xee, 0xf9, 0xf8, 0xc7, 0x01, 0x7f, 0x7c, 0x45, 0x45,
	0x0a, 0x5f, 0x5a, 0x20, 0x97, 0x76, 0x52, 0xde, 0x1b, 0x7e, 0x92, 0x55, 0xa8, 0xaf, 0xe7, 0x65,
	0xed, 0xc0, 0x2e, 0x54, 0xe5, 0xb3, 0x88, 0x25, 0x87, 0xb3, 0xb3, 0xfb, 0xe5, 0xf4, 0xe9, 0x8d,
	0x32, 0xd5, 0xab, 0x78, 0x97, 0x8f, 0xd3, 0x49, 0x77, 0xf0, 0x6e, 0xc0, 0x5e, 0xc0, 0x8f, 0x35,
	0xd6, 0xd4, 0x1a, 0xb0, 0xfe, 0x41, 0xc4, 0x92, 0x07, 0xe9, 0x16, 0x7b, 0xa7, 0xfc, 0x08, 0xad,
	0xd2, 0x15, 0xf8, 0xe3, 0x88, 0x25, 0xc7, 0xe9, 0x80, 0xbc, 0x6f, 0x8c, 0x3f, 0x22, 0xb0, 0x8b,
	0x52, 0xc3, 0xfc, 0x13, 0xc0, 0x5c, 0xab, 0xc6, 0x3f, 0x8c, 0xc6, 0xc9, 0xc3, 0x8b, 0x33, 0xd1,
	0xbb, 0x10, 0x1b, 0x17, 0x62, 0x70, 0x21, 0x2e, 0xb1, 0xac, 0x67, 0xef, 0x6f, 0x97, 0xd3, 0xd1,
	0xfd, 0x72, 0x7a, 0xda, 0x3f, 0x6b, 0xaf, 0x3f, 0xfe, 0xf9, 0x7b, 0x9a, 0x14, 0xa5, 0xfb, 0xdc,
	0x66, 0x42, 0xa3, 0x91, 0x43, 0x18, 0xfd, 0xf2, 0x82, 0xf2, 0x6b, 0xe9, 0x6e, 0x1a, 0xa0, 0x4e,
	0x8a, 0xd2, 0xc9, 0xd0, 0xfd, 0x16, 0xe0, 0x52, 0x35, 0x71, 0xc0, 0xfd, 0xfd, 0x64, 0x52, 0xa0,
	0x06, 0x6b, 0x82, 0x8b, 0x8c, 0x8f, 0xaf, 0xa8, 0xf0, 0x3e, 0xf2, 0xc9, 0x6e, 0x72, 0x91, 0xd8,
	0x9d, 0x8a, 0xd8, 0x57, 0x08, 0x92, 0xff, 0x55, 0xfc, 0xbb, 0x63, 0xf6, 0xfa, 0x76, 0x15, 0xb2,
	0xbb, 0x55, 0xc8, 0xfe, 0xac, 0x42, 0xf6, 0x7d, 0x1d, 0x8e, 0xee, 0xd6, 0xe1, 0xe8, 0xd7, 0x3a,
	0x1c, 0x7d, 0x78, 0x6e, 0x30, 0xa7, 0x4e, 0xa7, 0x06, 0x27, 0xd0, 0x16, 0xd2, 0x60, 0xde, 0x56,
	0x40, 0xdb, 0x4f, 0xb2, 0x71, 0x95, 0x1d, 0x75, 0x23, 0x7e, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x01, 0x1d, 0x05, 0x40, 0x43, 0x02, 0x00, 0x00,
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
	// RequestRandom defines a method for requesting a new random number
	RequestRandom(ctx context.Context, in *MsgRequestRandom, opts ...grpc.CallOption) (*MsgRequestRandomResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestRandom(ctx context.Context, in *MsgRequestRandom, opts ...grpc.CallOption) (*MsgRequestRandomResponse, error) {
	out := new(MsgRequestRandomResponse)
	err := c.cc.Invoke(ctx, "/irismod.random.Msg/RequestRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RequestRandom defines a method for requesting a new random number
	RequestRandom(context.Context, *MsgRequestRandom) (*MsgRequestRandomResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RequestRandom(ctx context.Context, req *MsgRequestRandom) (*MsgRequestRandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRandom not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RequestRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestRandom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.random.Msg/RequestRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestRandom(ctx, req.(*MsgRequestRandom))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.random.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestRandom",
			Handler:    _Msg_RequestRandom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/random/tx.proto",
}

func (m *MsgRequestRandom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRandom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRandom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceFeeCap) > 0 {
		for iNdEx := len(m.ServiceFeeCap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceFeeCap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Oracle {
		i--
		if m.Oracle {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x12
	}
	if m.BlockInterval != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlockInterval))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestRandomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRandomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRandomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRequestRandom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockInterval != 0 {
		n += 1 + sovTx(uint64(m.BlockInterval))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Oracle {
		n += 2
	}
	if len(m.ServiceFeeCap) > 0 {
		for _, e := range m.ServiceFeeCap {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgRequestRandomResponse) Size() (n int) {
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
func (m *MsgRequestRandom) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRequestRandom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRandom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockInterval", wireType)
			}
			m.BlockInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
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
			m.Consumer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
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
			m.Oracle = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceFeeCap", wireType)
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
			m.ServiceFeeCap = append(m.ServiceFeeCap, types.Coin{})
			if err := m.ServiceFeeCap[len(m.ServiceFeeCap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgRequestRandomResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRequestRandomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRandomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
