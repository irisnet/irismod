// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nft/rental.proto

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

// RentalPlugin defines the message for a denom rental plugin config
type RentalPlugin struct {
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *RentalPlugin) Reset()         { *m = RentalPlugin{} }
func (m *RentalPlugin) String() string { return proto.CompactTextString(m) }
func (*RentalPlugin) ProtoMessage()    {}
func (*RentalPlugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_e57beeaf512eaaa2, []int{0}
}
func (m *RentalPlugin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RentalPlugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RentalPlugin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RentalPlugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RentalPlugin.Merge(m, src)
}
func (m *RentalPlugin) XXX_Size() int {
	return m.Size()
}
func (m *RentalPlugin) XXX_DiscardUnknown() {
	xxx_messageInfo_RentalPlugin.DiscardUnknown(m)
}

var xxx_messageInfo_RentalPlugin proto.InternalMessageInfo

func (m *RentalPlugin) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// RentalInfo defines the message for an nft rental info
type RentalInfo struct {
	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Expires int64  `protobuf:"varint,2,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (m *RentalInfo) Reset()         { *m = RentalInfo{} }
func (m *RentalInfo) String() string { return proto.CompactTextString(m) }
func (*RentalInfo) ProtoMessage()    {}
func (*RentalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e57beeaf512eaaa2, []int{1}
}
func (m *RentalInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RentalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RentalInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RentalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RentalInfo.Merge(m, src)
}
func (m *RentalInfo) XXX_Size() int {
	return m.Size()
}
func (m *RentalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RentalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RentalInfo proto.InternalMessageInfo

func (m *RentalInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *RentalInfo) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*RentalPlugin)(nil), "irismod.nft.RentalPlugin")
	proto.RegisterType((*RentalInfo)(nil), "irismod.nft.RentalInfo")
}

func init() { proto.RegisterFile("nft/rental.proto", fileDescriptor_e57beeaf512eaaa2) }

var fileDescriptor_e57beeaf512eaaa2 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x4b, 0x2b, 0xd1,
	0x2f, 0x4a, 0xcd, 0x2b, 0x49, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0x2c,
	0xca, 0x2c, 0xce, 0xcd, 0x4f, 0xd1, 0xcb, 0x4b, 0x2b, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x8b, 0xeb, 0x83, 0x58, 0x10, 0x25, 0x4a, 0x7a, 0x5c, 0x3c, 0x41, 0x60, 0x2d, 0x01, 0x39, 0xa5,
	0xe9, 0x99, 0x79, 0x42, 0x12, 0x5c, 0xec, 0xa9, 0x79, 0x89, 0x49, 0x39, 0xa9, 0x29, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x30, 0xae, 0x15, 0xcb, 0x8b, 0x05, 0xf2, 0x8c, 0x4a, 0x0e, 0x5c,
	0x5c, 0x10, 0xf5, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x42, 0x5c, 0x2c, 0xa5, 0xc5, 0xa9, 0x45, 0x60,
	0xa5, 0x9c, 0x41, 0x60, 0x36, 0xd8, 0x84, 0x8a, 0x82, 0xcc, 0xa2, 0xd4, 0x62, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xe6, 0x20, 0x18, 0x17, 0x62, 0x82, 0x93, 0xdb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0x83, 0x5c, 0x9e, 0x97, 0x5a, 0xa2, 0x0f, 0xf5, 0x81, 0x7e, 0x6e, 0x7e, 0x4a, 0x69, 0x4e, 0x6a,
	0xb1, 0x3e, 0xc8, 0x8f, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x0f, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xc4, 0xc4, 0x34, 0xf7, 0x00, 0x00, 0x00,
}

func (this *RentalPlugin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RentalPlugin)
	if !ok {
		that2, ok := that.(RentalPlugin)
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
	if this.Enabled != that1.Enabled {
		return false
	}
	return true
}
func (this *RentalInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RentalInfo)
	if !ok {
		that2, ok := that.(RentalInfo)
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
	if this.User != that1.User {
		return false
	}
	if this.Expires != that1.Expires {
		return false
	}
	return true
}
func (m *RentalPlugin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RentalPlugin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RentalPlugin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RentalInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RentalInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RentalInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expires != 0 {
		i = encodeVarintRental(dAtA, i, uint64(m.Expires))
		i--
		dAtA[i] = 0x10
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintRental(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRental(dAtA []byte, offset int, v uint64) int {
	offset -= sovRental(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RentalPlugin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *RentalInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovRental(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovRental(uint64(m.Expires))
	}
	return n
}

func sovRental(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRental(x uint64) (n int) {
	return sovRental(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RentalPlugin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRental
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
			return fmt.Errorf("proto: RentalPlugin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RentalPlugin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
		default:
			iNdEx = preIndex
			skippy, err := skipRental(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRental
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
func (m *RentalInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRental
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
			return fmt.Errorf("proto: RentalInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RentalInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
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
				return ErrInvalidLengthRental
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRental
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRental
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRental(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRental
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
func skipRental(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRental
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
					return 0, ErrIntOverflowRental
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
					return 0, ErrIntOverflowRental
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
				return 0, ErrInvalidLengthRental
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRental
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRental
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRental        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRental          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRental = fmt.Errorf("proto: unexpected end of group")
)