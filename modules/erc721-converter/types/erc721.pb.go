// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: erc721_converter/v1/erc721.proto

package types

import (
	fmt "fmt"
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

// Owner enumerates the ownership of a ERC721 contract.
type Owner int32

const (
	// OWNER_UNSPECIFIED defines an invalid/undefined owner.
	OWNER_UNSPECIFIED Owner = 0
	// OWNER_MODULE - ERC721 is owned by the ERC721 module account.
	OWNER_MODULE Owner = 1
	// OWNER_EXTERNAL - ERC721 is owned by an external account.
	OWNER_EXTERNAL Owner = 2
)

var Owner_name = map[int32]string{
	0: "OWNER_UNSPECIFIED",
	1: "OWNER_MODULE",
	2: "OWNER_EXTERNAL",
}

var Owner_value = map[string]int32{
	"OWNER_UNSPECIFIED": 0,
	"OWNER_MODULE":      1,
	"OWNER_EXTERNAL":    2,
}

func (x Owner) String() string {
	return proto.EnumName(Owner_name, int32(x))
}

func (Owner) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65233a5373a5dc6f, []int{0}
}

// TokenPair defines an instance that records a pairing consisting of a native
//  Cosmos Coin and an ERC721 token address.
type TokenPair struct {
	// erc721_address is the hex address of ERC721 contract token
	Erc721Address string `protobuf:"bytes,1,opt,name=erc721_address,json=erc721Address,proto3" json:"erc721_address,omitempty"`
	// classId is the class of Cosmos x/nft Class ID
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// enabled defines the token mapping enable status
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// contract_owner is the an ENUM specifying the type of ERC721 owner (0 invalid, 1 ModuleAccount, 2 external address)
	ContractOwner Owner `protobuf:"varint,4,opt,name=contract_owner,json=contractOwner,proto3,enum=irismod.erc721_converter.v1.Owner" json:"contract_owner,omitempty"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_65233a5373a5dc6f, []int{0}
}
func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return m.Size()
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetErc721Address() string {
	if m != nil {
		return m.Erc721Address
	}
	return ""
}

func (m *TokenPair) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *TokenPair) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *TokenPair) GetContractOwner() Owner {
	if m != nil {
		return m.ContractOwner
	}
	return OWNER_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("irismod.erc721_converter.v1.Owner", Owner_name, Owner_value)
	proto.RegisterType((*TokenPair)(nil), "irismod.erc721_converter.v1.TokenPair")
}

func init() { proto.RegisterFile("erc721_converter/v1/erc721.proto", fileDescriptor_65233a5373a5dc6f) }

var fileDescriptor_65233a5373a5dc6f = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2d, 0x4a, 0x36,
	0x37, 0x32, 0x8c, 0x4f, 0xce, 0xcf, 0x2b, 0x4b, 0x2d, 0x2a, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0xd4,
	0x87, 0x88, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x49, 0x67, 0x16, 0x65, 0x16, 0xe7, 0xe6,
	0xa7, 0xe8, 0xa1, 0xab, 0xd4, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd3,
	0x07, 0xb1, 0x20, 0x5a, 0x94, 0x76, 0x30, 0x72, 0x71, 0x86, 0xe4, 0x67, 0xa7, 0xe6, 0x05, 0x24,
	0x66, 0x16, 0x09, 0xa9, 0x72, 0xf1, 0x41, 0xb5, 0x26, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xf1, 0x42, 0x44, 0x1d, 0x21, 0x82, 0x42, 0x92, 0x5c, 0x1c,
	0xc9, 0x39, 0x89, 0xc5, 0xc5, 0xf1, 0x99, 0x29, 0x12, 0x4c, 0x60, 0x05, 0xec, 0x60, 0xbe, 0x67,
	0x8a, 0x90, 0x04, 0x17, 0x7b, 0x6a, 0x5e, 0x62, 0x52, 0x4e, 0x6a, 0x8a, 0x04, 0xb3, 0x02, 0xa3,
	0x06, 0x47, 0x10, 0x8c, 0x2b, 0xe4, 0xc9, 0xc5, 0x97, 0x9c, 0x9f, 0x57, 0x52, 0x94, 0x98, 0x5c,
	0x12, 0x9f, 0x5f, 0x9e, 0x97, 0x5a, 0x24, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x67, 0xa4, 0xa4, 0x87,
	0xc7, 0xd5, 0x7a, 0xfe, 0x20, 0x95, 0x41, 0xbc, 0x30, 0x9d, 0x60, 0xae, 0x15, 0xcb, 0x8b, 0x05,
	0xf2, 0x8c, 0x5a, 0x5e, 0x5c, 0xac, 0x60, 0xae, 0x90, 0x28, 0x97, 0xa0, 0x7f, 0xb8, 0x9f, 0x6b,
	0x50, 0x7c, 0xa8, 0x5f, 0x70, 0x80, 0xab, 0xb3, 0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00, 0x83, 0x90,
	0x00, 0x17, 0x0f, 0x44, 0xd8, 0xd7, 0xdf, 0x25, 0xd4, 0xc7, 0x55, 0x80, 0x51, 0x48, 0x88, 0x8b,
	0x0f, 0x22, 0xe2, 0x1a, 0x11, 0xe2, 0x1a, 0xe4, 0xe7, 0xe8, 0x23, 0xc0, 0x24, 0xc5, 0xd2, 0xb1,
	0x58, 0x8e, 0xc1, 0x29, 0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x2c,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x41, 0x0e, 0xcd, 0x4b, 0x2d,
	0xd1, 0x87, 0x3a, 0x58, 0x3f, 0x37, 0x3f, 0xa5, 0x34, 0x27, 0xb5, 0x18, 0x1a, 0x09, 0xba, 0x88,
	0x88, 0x29, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x07, 0xb1, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x86, 0xf6, 0xcf, 0xe3, 0xb9, 0x01, 0x00, 0x00,
}

func (this *TokenPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenPair)
	if !ok {
		that2, ok := that.(TokenPair)
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
	if this.Erc721Address != that1.Erc721Address {
		return false
	}
	if this.ClassId != that1.ClassId {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.ContractOwner != that1.ContractOwner {
		return false
	}
	return true
}
func (m *TokenPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ContractOwner != 0 {
		i = encodeVarintErc721(dAtA, i, uint64(m.ContractOwner))
		i--
		dAtA[i] = 0x20
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Erc721Address) > 0 {
		i -= len(m.Erc721Address)
		copy(dAtA[i:], m.Erc721Address)
		i = encodeVarintErc721(dAtA, i, uint64(len(m.Erc721Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErc721(dAtA []byte, offset int, v uint64) int {
	offset -= sovErc721(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Erc721Address)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovErc721(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if m.ContractOwner != 0 {
		n += 1 + sovErc721(uint64(m.ContractOwner))
	}
	return n
}

func sovErc721(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErc721(x uint64) (n int) {
	return sovErc721(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc721
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
			return fmt.Errorf("proto: TokenPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc721Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc721Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
				return ErrInvalidLengthErc721
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc721
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
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
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOwner", wireType)
			}
			m.ContractOwner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc721
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractOwner |= Owner(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipErc721(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc721
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
func skipErc721(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErc721
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
					return 0, ErrIntOverflowErc721
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
					return 0, ErrIntOverflowErc721
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
				return 0, ErrInvalidLengthErc721
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErc721
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErc721
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErc721        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErc721          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErc721 = fmt.Errorf("proto: unexpected end of group")
)
