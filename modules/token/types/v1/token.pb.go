// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: irismod/token/v1/token.proto

package v1

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Token defines a standard for the fungible token
type Token struct {
	Symbol        string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Scale         uint32 `protobuf:"varint,3,opt,name=scale,proto3" json:"scale,omitempty"`
	MinUnit       string `protobuf:"bytes,4,opt,name=min_unit,json=minUnit,proto3" json:"min_unit,omitempty"`
	InitialSupply uint64 `protobuf:"varint,5,opt,name=initial_supply,json=initialSupply,proto3" json:"initial_supply,omitempty"`
	MaxSupply     uint64 `protobuf:"varint,6,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	Mintable      bool   `protobuf:"varint,7,opt,name=mintable,proto3" json:"mintable,omitempty"`
	Owner         string `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Contract      string `protobuf:"bytes,9,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5b3436d30fd508a, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

// Params defines token module's parameters
type Params struct {
	TokenTaxRate      cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=token_tax_rate,json=tokenTaxRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"token_tax_rate"`
	IssueTokenBaseFee types.Coin                  `protobuf:"bytes,2,opt,name=issue_token_base_fee,json=issueTokenBaseFee,proto3" json:"issue_token_base_fee"`
	MintTokenFeeRatio cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=mint_token_fee_ratio,json=mintTokenFeeRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"mint_token_fee_ratio"`
	EnableErc20       bool                        `protobuf:"varint,4,opt,name=enable_erc20,json=enableErc20,proto3" json:"enable_erc20,omitempty"`
	Beacon            string                      `protobuf:"bytes,5,opt,name=beacon,proto3" json:"beacon,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5b3436d30fd508a, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Token)(nil), "irismod.token.v1.Token")
	proto.RegisterType((*Params)(nil), "irismod.token.v1.Params")
}

func init() { proto.RegisterFile("irismod/token/v1/token.proto", fileDescriptor_c5b3436d30fd508a) }

var fileDescriptor_c5b3436d30fd508a = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x6e, 0x13, 0x3f,
	0x10, 0xde, 0xed, 0x6f, 0x93, 0x26, 0xee, 0x1f, 0xfd, 0x6a, 0x45, 0x68, 0x1b, 0x60, 0x13, 0x8a,
	0x10, 0x39, 0x79, 0x49, 0xb9, 0xf5, 0x18, 0x68, 0x25, 0x24, 0x0e, 0x95, 0x09, 0x17, 0x2e, 0x2b,
	0xef, 0x66, 0x1a, 0xac, 0xae, 0xed, 0x68, 0xed, 0x84, 0xe4, 0x0d, 0x38, 0xf2, 0x08, 0x7d, 0x10,
	0x1e, 0x20, 0xc7, 0x1e, 0x11, 0x87, 0x0a, 0x92, 0x0b, 0x4f, 0x81, 0x90, 0xed, 0xa5, 0x67, 0x6e,
	0xf3, 0x7d, 0x9e, 0xf9, 0x66, 0x3c, 0xdf, 0xa0, 0x47, 0xbc, 0xe2, 0x5a, 0xa8, 0x49, 0x6a, 0xd4,
	0x35, 0xc8, 0x74, 0x31, 0xf4, 0x01, 0x99, 0x55, 0xca, 0x28, 0xfc, 0x7f, 0xfd, 0x4a, 0x3c, 0xb9,
	0x18, 0x76, 0x93, 0x42, 0x69, 0xa1, 0x74, 0x9a, 0x33, 0x0d, 0xe9, 0x62, 0x98, 0x83, 0x61, 0xc3,
	0xb4, 0x50, 0xbc, 0xae, 0xe8, 0x76, 0xa6, 0x6a, 0xaa, 0x5c, 0x98, 0xda, 0xc8, 0xb3, 0x27, 0xbf,
	0x43, 0xd4, 0x18, 0x5b, 0x09, 0xfc, 0x00, 0x35, 0xf5, 0x4a, 0xe4, 0xaa, 0x8c, 0xc3, 0x7e, 0x38,
	0x68, 0xd3, 0x1a, 0x61, 0x8c, 0x22, 0xc9, 0x04, 0xc4, 0x3b, 0x8e, 0x75, 0x31, 0xee, 0xa0, 0x86,
	0x2e, 0x58, 0x09, 0xf1, 0x7f, 0xfd, 0x70, 0x70, 0x40, 0x3d, 0xc0, 0xc7, 0xa8, 0x25, 0xb8, 0xcc,
	0xe6, 0x92, 0x9b, 0x38, 0x72, 0xd9, 0xbb, 0x82, 0xcb, 0xf7, 0x92, 0x1b, 0xfc, 0x0c, 0x1d, 0x72,
	0xc9, 0x0d, 0x67, 0x65, 0xa6, 0xe7, 0xb3, 0x59, 0xb9, 0x8a, 0x1b, 0xfd, 0x70, 0x10, 0xd1, 0x83,
	0x9a, 0x7d, 0xe7, 0x48, 0xfc, 0x18, 0x21, 0xc1, 0x96, 0x7f, 0x53, 0x9a, 0x2e, 0xa5, 0x2d, 0xd8,
	0xb2, 0x7e, 0xee, 0xba, 0x06, 0x86, 0xe5, 0x25, 0xc4, 0xbb, 0xfd, 0x70, 0xd0, 0xa2, 0xf7, 0xd8,
	0x8e, 0xa4, 0x3e, 0x49, 0xa8, 0xe2, 0x96, 0xeb, 0xec, 0x81, 0xad, 0x28, 0x94, 0x34, 0x15, 0x2b,
	0x4c, 0xdc, 0x76, 0x0f, 0xf7, 0xf8, 0x2c, 0xfa, 0x7c, 0xd3, 0x0b, 0x4e, 0xbe, 0xee, 0xa0, 0xe6,
	0x25, 0xab, 0x98, 0xd0, 0xf8, 0x0d, 0x3a, 0x74, 0xdb, 0xcc, 0x0c, 0x5b, 0x66, 0x15, 0x33, 0xe0,
	0x37, 0x31, 0x7a, 0xba, 0xbe, 0xeb, 0x05, 0xdf, 0xef, 0x7a, 0x0f, 0xfd, 0x86, 0xf5, 0xe4, 0x9a,
	0x70, 0x95, 0x0a, 0x66, 0x3e, 0x92, 0xb7, 0x30, 0x65, 0xc5, 0xea, 0x35, 0x14, 0x74, 0xdf, 0x95,
	0x8e, 0xd9, 0x92, 0x32, 0x03, 0xf8, 0x12, 0x75, 0xb8, 0xd6, 0x73, 0xc8, 0xbc, 0xa0, 0xf5, 0x24,
	0xbb, 0x02, 0xbf, 0xc4, 0xbd, 0xd3, 0x63, 0xe2, 0x95, 0x88, 0xe5, 0x49, 0xed, 0x15, 0x79, 0xa5,
	0xb8, 0x1c, 0x45, 0xb6, 0x17, 0x3d, 0x72, 0xc5, 0xce, 0x97, 0x11, 0xd3, 0x70, 0x01, 0x80, 0xc7,
	0xa8, 0x63, 0xff, 0x5a, 0x0b, 0x5e, 0x01, 0xd8, 0x09, 0xb9, 0x72, 0x0e, 0xfc, 0xe3, 0x88, 0x47,
	0x56, 0xc0, 0x89, 0x5e, 0x00, 0x50, 0x5b, 0x8d, 0x9f, 0xa0, 0x7d, 0x90, 0x76, 0x7f, 0x19, 0x54,
	0xc5, 0xe9, 0x0b, 0x67, 0x5b, 0x8b, 0xee, 0x79, 0xee, 0xdc, 0x52, 0xf6, 0x2e, 0x72, 0x60, 0x85,
	0x92, 0xce, 0xb2, 0x36, 0xad, 0xd1, 0x59, 0xf4, 0xeb, 0xa6, 0x17, 0x8e, 0xce, 0xd7, 0x3f, 0x93,
	0x60, 0xbd, 0x49, 0xc2, 0xdb, 0x4d, 0x12, 0xfe, 0xd8, 0x24, 0xe1, 0x97, 0x6d, 0x12, 0xdc, 0x6e,
	0x93, 0xe0, 0xdb, 0x36, 0x09, 0x3e, 0x3c, 0x17, 0x6a, 0xa2, 0x89, 0x3d, 0x55, 0x09, 0x86, 0xa8,
	0x6a, 0x9a, 0x0a, 0x35, 0x99, 0x97, 0xa0, 0xeb, 0xa3, 0x36, 0xab, 0x19, 0x68, 0x7b, 0xac, 0x4d,
	0x77, 0x8d, 0x2f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x0c, 0xff, 0x53, 0xf5, 0x02, 0x00,
	0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if !this.TokenTaxRate.Equal(that1.TokenTaxRate) {
		return false
	}
	if !this.IssueTokenBaseFee.Equal(&that1.IssueTokenBaseFee) {
		return false
	}
	if !this.MintTokenFeeRatio.Equal(that1.MintTokenFeeRatio) {
		return false
	}
	if this.EnableErc20 != that1.EnableErc20 {
		return false
	}
	if this.Beacon != that1.Beacon {
		return false
	}
	return true
}
func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x42
	}
	if m.Mintable {
		i--
		if m.Mintable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.MaxSupply != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x30
	}
	if m.InitialSupply != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.InitialSupply))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MinUnit) > 0 {
		i -= len(m.MinUnit)
		copy(dAtA[i:], m.MinUnit)
		i = encodeVarintToken(dAtA, i, uint64(len(m.MinUnit)))
		i--
		dAtA[i] = 0x22
	}
	if m.Scale != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.Scale))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Beacon) > 0 {
		i -= len(m.Beacon)
		copy(dAtA[i:], m.Beacon)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Beacon)))
		i--
		dAtA[i] = 0x2a
	}
	if m.EnableErc20 {
		i--
		if m.EnableErc20 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.MintTokenFeeRatio.Size()
		i -= size
		if _, err := m.MintTokenFeeRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.IssueTokenBaseFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.TokenTaxRate.Size()
		i -= size
		if _, err := m.TokenTaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintToken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.Scale != 0 {
		n += 1 + sovToken(uint64(m.Scale))
	}
	l = len(m.MinUnit)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	if m.InitialSupply != 0 {
		n += 1 + sovToken(uint64(m.InitialSupply))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovToken(uint64(m.MaxSupply))
	}
	if m.Mintable {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenTaxRate.Size()
	n += 1 + l + sovToken(uint64(l))
	l = m.IssueTokenBaseFee.Size()
	n += 1 + l + sovToken(uint64(l))
	l = m.MintTokenFeeRatio.Size()
	n += 1 + l + sovToken(uint64(l))
	if m.EnableErc20 {
		n += 2
	}
	l = len(m.Beacon)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scale", wireType)
			}
			m.Scale = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Scale |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinUnit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSupply", wireType)
			}
			m.InitialSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
			m.Mintable = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenTaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenTaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueTokenBaseFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IssueTokenBaseFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintTokenFeeRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintTokenFeeRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableErc20", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
			m.EnableErc20 = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beacon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Beacon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
