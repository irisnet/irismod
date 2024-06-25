// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: irismod/farm/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the genesis information exported by the farm module
type GenesisState struct {
	Params    Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Pools     []FarmPool   `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools"`
	FarmInfos []FarmInfo   `protobuf:"bytes,3,rep,name=farm_infos,json=farmInfos,proto3" json:"farm_infos"`
	Sequence  uint64       `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Escrow    []EscrowInfo `protobuf:"bytes,5,rep,name=escrow,proto3" json:"escrow"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_93a576c7f1f8c765, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPools() []FarmPool {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *GenesisState) GetFarmInfos() []FarmInfo {
	if m != nil {
		return m.FarmInfos
	}
	return nil
}

func (m *GenesisState) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *GenesisState) GetEscrow() []EscrowInfo {
	if m != nil {
		return m.Escrow
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "irismod.farm.GenesisState")
}

func init() { proto.RegisterFile("irismod/farm/genesis.proto", fileDescriptor_93a576c7f1f8c765) }

var fileDescriptor_93a576c7f1f8c765 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x73, 0xfd, 0x87, 0x5e, 0x3b, 0x1d, 0x45, 0x8f, 0x0c, 0x67, 0xd1, 0x25, 0x53, 0x02,
	0x11, 0x5c, 0xc4, 0xa5, 0xa0, 0xe2, 0x56, 0xea, 0xe6, 0x22, 0xb1, 0xb9, 0x84, 0x40, 0x2e, 0x6f,
	0xbc, 0x37, 0x45, 0xfc, 0x04, 0xae, 0x7e, 0xac, 0x8e, 0x1d, 0x9d, 0x44, 0x92, 0x2f, 0x22, 0xb9,
	0x1c, 0xd2, 0x0c, 0x2e, 0xc7, 0xbd, 0xfc, 0x9e, 0xe7, 0x37, 0x3c, 0xd4, 0xcd, 0x74, 0x86, 0x0a,
	0xe2, 0x20, 0x89, 0xb4, 0x0a, 0x52, 0x59, 0x48, 0xcc, 0xd0, 0x2f, 0x35, 0x54, 0xc0, 0x66, 0x96,
	0xf9, 0x2d, 0x73, 0xe7, 0x29, 0xa4, 0x60, 0x40, 0xd0, 0xfe, 0xba, 0x8c, 0x7b, 0xda, 0xeb, 0xb7,
	0x4f, 0x07, 0xce, 0x3f, 0x06, 0x74, 0x76, 0xdf, 0xe9, 0x1e, 0xab, 0xa8, 0x92, 0x2c, 0xa4, 0x93,
	0x32, 0xd2, 0x91, 0x42, 0x4e, 0x16, 0xc4, 0x9b, 0x86, 0x73, 0xff, 0x50, 0xef, 0xaf, 0x0c, 0x5b,
	0x8e, 0x76, 0xdf, 0x67, 0xce, 0xda, 0x26, 0x59, 0x48, 0xc7, 0x25, 0x40, 0x8e, 0x7c, 0xb0, 0x18,
	0x7a, 0xd3, 0xf0, 0xa4, 0x5f, 0xb9, 0x8b, 0xb4, 0x5a, 0x01, 0xe4, 0xb6, 0xd4, 0x45, 0xd9, 0x35,
	0xa5, 0x2d, 0x7d, 0xce, 0x8a, 0x04, 0x90, 0x0f, 0xff, 0x2b, 0x3e, 0x14, 0x09, 0xd8, 0xe2, 0x71,
	0x62, 0x6f, 0x64, 0x2e, 0x3d, 0x42, 0xf9, 0xba, 0x95, 0xc5, 0x46, 0xf2, 0xd1, 0x82, 0x78, 0xa3,
	0xf5, 0xdf, 0xcd, 0xae, 0xe8, 0x44, 0xe2, 0x46, 0xc3, 0x1b, 0x1f, 0x1b, 0x29, 0xef, 0x4b, 0x6f,
	0x0d, 0x3b, 0xd0, 0xda, 0xf4, 0xf2, 0x66, 0x57, 0x0b, 0xb2, 0xaf, 0x05, 0xf9, 0xa9, 0x05, 0xf9,
	0x6c, 0x84, 0xb3, 0x6f, 0x84, 0xf3, 0xd5, 0x08, 0xe7, 0xe9, 0x42, 0x41, 0x8c, 0xc6, 0x52, 0xc8,
	0xca, 0x07, 0x9d, 0x06, 0x0a, 0xe2, 0x6d, 0x2e, 0xb1, 0x5b, 0xb3, 0x7a, 0x2f, 0x25, 0xbe, 0x4c,
	0xcc, 0x9e, 0x97, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0x33, 0xde, 0xbe, 0xaa, 0x01, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Escrow) > 0 {
		for iNdEx := len(m.Escrow) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Escrow[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Sequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x20
	}
	if len(m.FarmInfos) > 0 {
		for iNdEx := len(m.FarmInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FarmInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FarmInfos) > 0 {
		for _, e := range m.FarmInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Sequence != 0 {
		n += 1 + sovGenesis(uint64(m.Sequence))
	}
	if len(m.Escrow) > 0 {
		for _, e := range m.Escrow {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, FarmPool{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FarmInfos = append(m.FarmInfos, FarmInfo{})
			if err := m.FarmInfos[len(m.FarmInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Escrow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Escrow = append(m.Escrow, EscrowInfo{})
			if err := m.Escrow[len(m.Escrow)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
