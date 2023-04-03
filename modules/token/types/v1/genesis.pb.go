// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/v1/genesis.proto

package v1

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the token module's genesis state
type GenesisState struct {
	Params      Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Tokens      []Token      `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens"`
	BurnedCoins []types.Coin `protobuf:"bytes,3,rep,name=burned_coins,json=burnedCoins,proto3" json:"burned_coins"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_25c9badd7ffb5250, []int{0}
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

func (m *GenesisState) GetTokens() []Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *GenesisState) GetBurnedCoins() []types.Coin {
	if m != nil {
		return m.BurnedCoins
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "irismod.token.v1.GenesisState")
}

func init() { proto.RegisterFile("token/v1/genesis.proto", fileDescriptor_25c9badd7ffb5250) }

var fileDescriptor_25c9badd7ffb5250 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xe3, 0xaf, 0x9f, 0x32, 0x24, 0x1d, 0x50, 0x54, 0x41, 0xe8, 0x60, 0x2a, 0xa6, 0x4e,
	0xb6, 0x52, 0x04, 0x17, 0x10, 0x06, 0x06, 0x16, 0x04, 0x4c, 0x2c, 0x28, 0x3f, 0x56, 0xb0, 0xc0,
	0x39, 0x51, 0x8e, 0x13, 0x89, 0xbb, 0xe0, 0x96, 0xd8, 0x3a, 0x76, 0x64, 0x42, 0x28, 0xb9, 0x11,
	0x64, 0x3b, 0x65, 0x80, 0xed, 0xe8, 0xbc, 0xef, 0xa3, 0xf3, 0xe8, 0x04, 0x87, 0x1a, 0x9e, 0x45,
	0xcd, 0xfb, 0x84, 0x57, 0xa2, 0x16, 0x28, 0x91, 0x35, 0x2d, 0x68, 0x88, 0x0e, 0x64, 0x2b, 0x51,
	0x41, 0xc9, 0x6c, 0xce, 0xfa, 0x64, 0xb9, 0xa8, 0xa0, 0x02, 0x1b, 0x72, 0x33, 0xb9, 0xde, 0x72,
	0xf1, 0xc3, 0xbb, 0xa2, 0xdb, 0xd2, 0x02, 0x50, 0x01, 0xf2, 0x3c, 0x43, 0xc1, 0xfb, 0x24, 0x17,
	0x3a, 0x4b, 0x78, 0x01, 0x72, 0xca, 0x4f, 0xdf, 0x49, 0x30, 0xbf, 0x72, 0xf7, 0xee, 0x74, 0xa6,
	0x45, 0x74, 0x11, 0xf8, 0x4d, 0xd6, 0x66, 0x0a, 0x63, 0xb2, 0x22, 0xeb, 0x70, 0x13, 0xb3, 0xdf,
	0xf7, 0xd9, 0x8d, 0xcd, 0xd3, 0xff, 0xdb, 0xcf, 0x13, 0xef, 0x76, 0x6a, 0x47, 0xe7, 0x81, 0x6f,
	0x0b, 0x18, 0xff, 0x5b, 0xcd, 0xd6, 0xe1, 0xe6, 0xe8, 0x2f, 0x77, 0x6f, 0x86, 0x3d, 0xe6, 0xca,
	0x51, 0x1a, 0xcc, 0xf3, 0xae, 0xad, 0x45, 0xf9, 0x68, 0xa4, 0x30, 0x9e, 0x59, 0xf8, 0x98, 0x39,
	0x6d, 0x66, 0xb4, 0xd9, 0xa4, 0xcd, 0x2e, 0x41, 0xee, 0xf1, 0xd0, 0x41, 0x66, 0x83, 0xe9, 0xf5,
	0x76, 0xa0, 0x64, 0x37, 0x50, 0xf2, 0x35, 0x50, 0xf2, 0x36, 0x52, 0x6f, 0x37, 0x52, 0xef, 0x63,
	0xa4, 0xde, 0x43, 0x52, 0x49, 0xfd, 0xd4, 0xe5, 0xac, 0x00, 0xc5, 0x8d, 0x4e, 0x2d, 0x34, 0x9f,
	0xb4, 0xb8, 0x82, 0xb2, 0x7b, 0x11, 0xe8, 0xbe, 0xc5, 0xf5, 0x6b, 0x23, 0xd0, 0x3c, 0xc8, 0xb7,
	0x7f, 0x39, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x23, 0x3e, 0x56, 0xcd, 0x8f, 0x01, 0x00, 0x00,
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
	if len(m.BurnedCoins) > 0 {
		for iNdEx := len(m.BurnedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BurnedCoins) > 0 {
		for _, e := range m.BurnedCoins {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
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
			m.Tokens = append(m.Tokens, Token{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnedCoins", wireType)
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
			m.BurnedCoins = append(m.BurnedCoins, types.Coin{})
			if err := m.BurnedCoins[len(m.BurnedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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