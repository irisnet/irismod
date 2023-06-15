// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/genesis.proto

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

// GenesisState defines the service module's genesis state
type GenesisState struct {
	Params            Params                     `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Definitions       []ServiceDefinition        `protobuf:"bytes,2,rep,name=definitions,proto3" json:"definitions"`
	Bindings          []ServiceBinding           `protobuf:"bytes,3,rep,name=bindings,proto3" json:"bindings"`
	WithdrawAddresses map[string]string          `protobuf:"bytes,4,rep,name=withdraw_addresses,json=withdrawAddresses,proto3" json:"withdraw_addresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RequestContexts   map[string]*RequestContext `protobuf:"bytes,5,rep,name=request_contexts,json=requestContexts,proto3" json:"request_contexts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_de239081d7e59627, []int{0}
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

func (m *GenesisState) GetDefinitions() []ServiceDefinition {
	if m != nil {
		return m.Definitions
	}
	return nil
}

func (m *GenesisState) GetBindings() []ServiceBinding {
	if m != nil {
		return m.Bindings
	}
	return nil
}

func (m *GenesisState) GetWithdrawAddresses() map[string]string {
	if m != nil {
		return m.WithdrawAddresses
	}
	return nil
}

func (m *GenesisState) GetRequestContexts() map[string]*RequestContext {
	if m != nil {
		return m.RequestContexts
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "irismod.service.GenesisState")
	proto.RegisterMapType((map[string]*RequestContext)(nil), "irismod.service.GenesisState.RequestContextsEntry")
	proto.RegisterMapType((map[string]string)(nil), "irismod.service.GenesisState.WithdrawAddressesEntry")
}

func init() { proto.RegisterFile("service/genesis.proto", fileDescriptor_de239081d7e59627) }

var fileDescriptor_de239081d7e59627 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0xcf, 0xd2, 0x30,
	0x18, 0xc7, 0x37, 0xe0, 0x7d, 0xa3, 0xc5, 0xe4, 0x7d, 0x6d, 0x50, 0x97, 0x1d, 0x06, 0xe1, 0xc4,
	0x69, 0x33, 0x28, 0x89, 0xf1, 0x06, 0x62, 0x4c, 0x38, 0x99, 0x71, 0x30, 0x31, 0x31, 0x64, 0x6c,
	0x8f, 0xa3, 0x91, 0xb5, 0xd8, 0x76, 0x20, 0xdf, 0xc2, 0xef, 0xe4, 0x85, 0x23, 0x47, 0x4f, 0xc6,
	0xc0, 0x17, 0x31, 0xb4, 0x85, 0x0c, 0x98, 0xef, 0xa9, 0xdd, 0xf3, 0x3c, 0xbf, 0xdf, 0x7f, 0x69,
	0x8b, 0x9e, 0x09, 0xe0, 0x4b, 0x12, 0x43, 0x90, 0x02, 0x05, 0x41, 0x84, 0xbf, 0xe0, 0x4c, 0x32,
	0x7c, 0x47, 0x38, 0x11, 0x19, 0x4b, 0x7c, 0xd3, 0x76, 0x1b, 0x29, 0x4b, 0x99, 0xea, 0x05, 0x87,
	0x9d, 0x1e, 0x73, 0x4f, 0xb4, 0x59, 0x75, 0xb9, 0xfd, 0xab, 0x86, 0x9e, 0x7c, 0xd0, 0xbe, 0xb1,
	0x8c, 0x24, 0xe0, 0x1e, 0xba, 0x5d, 0x44, 0x3c, 0xca, 0x84, 0x63, 0xb7, 0xec, 0x4e, 0xbd, 0xfb,
	0xc2, 0xbf, 0xf0, 0xfb, 0x1f, 0x55, 0x7b, 0x50, 0xdb, 0xfc, 0x69, 0x5a, 0xa1, 0x19, 0xc6, 0x23,
	0x54, 0x4f, 0xe0, 0x2b, 0xa1, 0x44, 0x12, 0x46, 0x85, 0x53, 0x69, 0x55, 0x3b, 0xf5, 0x6e, 0xfb,
	0x8a, 0x1d, 0xeb, 0x75, 0x78, 0x1a, 0x35, 0x9a, 0x22, 0x8c, 0xfb, 0xe8, 0xd1, 0x94, 0xd0, 0x84,
	0xd0, 0x54, 0x38, 0x55, 0x25, 0x6a, 0xfe, 0x4f, 0x34, 0xd0, 0x73, 0xc6, 0x72, 0xc2, 0x70, 0x8c,
	0xf0, 0x8a, 0xc8, 0x59, 0xc2, 0xa3, 0xd5, 0x24, 0x4a, 0x12, 0x0e, 0x42, 0x80, 0x70, 0x6a, 0x4a,
	0xf6, 0xfa, 0x4a, 0x56, 0x3c, 0x00, 0xff, 0x93, 0xe1, 0xfa, 0x47, 0xec, 0x3d, 0x95, 0x7c, 0x1d,
	0x3e, 0x5d, 0x5d, 0xd6, 0xf1, 0x17, 0x74, 0xcf, 0xe1, 0x7b, 0x0e, 0x42, 0x4e, 0x62, 0x46, 0x25,
	0xfc, 0x90, 0xc2, 0xb9, 0x51, 0x11, 0xdd, 0x87, 0x23, 0x42, 0x4d, 0xbd, 0x33, 0x90, 0x0e, 0xb8,
	0xe3, 0xe7, 0x55, 0x77, 0x88, 0x9e, 0x97, 0xff, 0x0b, 0xbe, 0x47, 0xd5, 0x6f, 0xb0, 0x56, 0x17,
	0xf4, 0x38, 0x3c, 0x6c, 0x71, 0x03, 0xdd, 0x2c, 0xa3, 0x79, 0x0e, 0x4e, 0x45, 0xd5, 0xf4, 0xc7,
	0xdb, 0xca, 0x1b, 0xdb, 0x8d, 0x51, 0xa3, 0x2c, 0xae, 0xc4, 0xd1, 0x2b, 0x3a, 0xca, 0xce, 0xfc,
	0xdc, 0x53, 0x08, 0x19, 0x8c, 0x36, 0x3b, 0xcf, 0xde, 0xee, 0x3c, 0xfb, 0xef, 0xce, 0xb3, 0x7f,
	0xee, 0x3d, 0x6b, 0xbb, 0xf7, 0xac, 0xdf, 0x7b, 0xcf, 0xfa, 0xfc, 0x32, 0x25, 0x72, 0x96, 0x4f,
	0xfd, 0x98, 0x65, 0xc1, 0xc1, 0x47, 0x41, 0x06, 0xc6, 0x1b, 0x64, 0x2c, 0xc9, 0xe7, 0x20, 0x8e,
	0x2f, 0x32, 0x90, 0xeb, 0x05, 0x88, 0xe9, 0xad, 0x7a, 0x98, 0xaf, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xee, 0x63, 0x04, 0x43, 0xef, 0x02, 0x00, 0x00,
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
	if len(m.RequestContexts) > 0 {
		for k := range m.RequestContexts {
			v := m.RequestContexts[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintGenesis(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.WithdrawAddresses) > 0 {
		for k := range m.WithdrawAddresses {
			v := m.WithdrawAddresses[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenesis(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Bindings) > 0 {
		for iNdEx := len(m.Bindings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bindings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Definitions) > 0 {
		for iNdEx := len(m.Definitions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Definitions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Definitions) > 0 {
		for _, e := range m.Definitions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Bindings) > 0 {
		for _, e := range m.Bindings {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WithdrawAddresses) > 0 {
		for k, v := range m.WithdrawAddresses {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + len(v) + sovGenesis(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	if len(m.RequestContexts) > 0 {
		for k, v := range m.RequestContexts {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovGenesis(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Definitions", wireType)
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
			m.Definitions = append(m.Definitions, ServiceDefinition{})
			if err := m.Definitions[len(m.Definitions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bindings", wireType)
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
			m.Bindings = append(m.Bindings, ServiceBinding{})
			if err := m.Bindings[len(m.Bindings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddresses", wireType)
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
			if m.WithdrawAddresses == nil {
				m.WithdrawAddresses = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.WithdrawAddresses[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestContexts", wireType)
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
			if m.RequestContexts == nil {
				m.RequestContexts = make(map[string]*RequestContext)
			}
			var mapkey string
			var mapvalue *RequestContext
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &RequestContext{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.RequestContexts[mapkey] = mapvalue
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
