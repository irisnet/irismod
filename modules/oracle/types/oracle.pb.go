// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/oracle.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Feed defines the feed standard
type Feed struct {
	FeedName         string `protobuf:"bytes,1,opt,name=feed_name,json=feedName,proto3" json:"feed_name,omitempty" yaml:"feed_name"`
	Description      string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AggregateFunc    string `protobuf:"bytes,3,opt,name=aggregate_func,json=aggregateFunc,proto3" json:"aggregate_func,omitempty" yaml:"aggregate_func"`
	ValueJsonPath    string `protobuf:"bytes,4,opt,name=value_json_path,json=valueJsonPath,proto3" json:"value_json_path,omitempty" yaml:"value_json_path"`
	LatestHistory    uint64 `protobuf:"varint,5,opt,name=latest_history,json=latestHistory,proto3" json:"latest_history,omitempty" yaml:"latest_history"`
	RequestContextID string `protobuf:"bytes,6,opt,name=request_context_id,json=requestContextId,proto3" json:"request_context_id,omitempty" yaml:"request_context_id"`
	Creator          string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Feed) Reset()         { *m = Feed{} }
func (m *Feed) String() string { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()    {}
func (*Feed) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc470b50b143d488, []int{0}
}
func (m *Feed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Feed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Feed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Feed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed.Merge(m, src)
}
func (m *Feed) XXX_Size() int {
	return m.Size()
}
func (m *Feed) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed.DiscardUnknown(m)
}

var xxx_messageInfo_Feed proto.InternalMessageInfo

func (m *Feed) GetFeedName() string {
	if m != nil {
		return m.FeedName
	}
	return ""
}

func (m *Feed) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Feed) GetAggregateFunc() string {
	if m != nil {
		return m.AggregateFunc
	}
	return ""
}

func (m *Feed) GetValueJsonPath() string {
	if m != nil {
		return m.ValueJsonPath
	}
	return ""
}

func (m *Feed) GetLatestHistory() uint64 {
	if m != nil {
		return m.LatestHistory
	}
	return 0
}

func (m *Feed) GetRequestContextID() string {
	if m != nil {
		return m.RequestContextID
	}
	return ""
}

func (m *Feed) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// FeedValue defines the feed result standard
type FeedValue struct {
	Data      string    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp time.Time `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *FeedValue) Reset()         { *m = FeedValue{} }
func (m *FeedValue) String() string { return proto.CompactTextString(m) }
func (*FeedValue) ProtoMessage()    {}
func (*FeedValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc470b50b143d488, []int{1}
}
func (m *FeedValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeedValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeedValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeedValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedValue.Merge(m, src)
}
func (m *FeedValue) XXX_Size() int {
	return m.Size()
}
func (m *FeedValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedValue.DiscardUnknown(m)
}

var xxx_messageInfo_FeedValue proto.InternalMessageInfo

func (m *FeedValue) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *FeedValue) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Feed)(nil), "irismod.oracle.Feed")
	proto.RegisterType((*FeedValue)(nil), "irismod.oracle.FeedValue")
}

func init() { proto.RegisterFile("oracle/oracle.proto", fileDescriptor_dc470b50b143d488) }

var fileDescriptor_dc470b50b143d488 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x69, 0x68, 0x9b, 0xad, 0x52, 0xa2, 0xa5, 0x20, 0x37, 0x07, 0x3b, 0xf2, 0xa9, 0x27,
	0xaf, 0x0a, 0x37, 0x4e, 0xc8, 0xa0, 0x8a, 0x72, 0x40, 0x68, 0x85, 0x38, 0x70, 0xb1, 0x36, 0xf6,
	0xc4, 0x31, 0xb2, 0xbd, 0x66, 0x77, 0x8c, 0xc8, 0x5f, 0xf4, 0x0b, 0xf8, 0x9e, 0x1e, 0x7b, 0xe4,
	0x64, 0x50, 0xf2, 0x07, 0xf9, 0x02, 0xe4, 0xdd, 0x26, 0x24, 0xf4, 0xe4, 0x99, 0xf7, 0xde, 0x3c,
	0x8f, 0xfd, 0x86, 0x3c, 0x95, 0x4a, 0x24, 0x05, 0x30, 0xfb, 0x08, 0x6b, 0x25, 0x51, 0xd2, 0xd3,
	0x5c, 0xe5, 0xba, 0x94, 0x69, 0x68, 0xd1, 0xf1, 0x59, 0x26, 0x33, 0x69, 0x28, 0xd6, 0x55, 0x56,
	0x35, 0xf6, 0x33, 0x29, 0xb3, 0x02, 0x98, 0xe9, 0xa6, 0xcd, 0x8c, 0x61, 0x5e, 0x82, 0x46, 0x51,
	0xd6, 0x56, 0x10, 0xfc, 0x3c, 0x20, 0xfd, 0x2b, 0x80, 0x94, 0x5e, 0x92, 0xc1, 0x0c, 0x20, 0x8d,
	0x2b, 0x51, 0x82, 0xeb, 0x4c, 0x9c, 0x8b, 0x41, 0x74, 0xb6, 0x6e, 0xfd, 0xd1, 0x42, 0x94, 0xc5,
	0xab, 0x60, 0x4b, 0x05, 0xfc, 0xb8, 0xab, 0x3f, 0x88, 0x12, 0xe8, 0x84, 0x9c, 0xa4, 0xa0, 0x13,
	0x95, 0xd7, 0x98, 0xcb, 0xca, 0x7d, 0xd4, 0x0d, 0xf1, 0x5d, 0x88, 0xbe, 0x26, 0xa7, 0x22, 0xcb,
	0x14, 0x64, 0x02, 0x21, 0x9e, 0x35, 0x55, 0xe2, 0x1e, 0x18, 0xe7, 0xf3, 0x75, 0xeb, 0x3f, 0xb3,
	0xce, 0xfb, 0x7c, 0xc0, 0x87, 0x5b, 0xe0, 0xaa, 0xa9, 0x12, 0x1a, 0x91, 0x27, 0xdf, 0x45, 0xd1,
	0x40, 0xfc, 0x55, 0xcb, 0x2a, 0xae, 0x05, 0xce, 0xdd, 0xbe, 0xb1, 0x18, 0xaf, 0x5b, 0xff, 0xb9,
	0xb5, 0xf8, 0x4f, 0x10, 0xf0, 0xa1, 0x41, 0xde, 0x6b, 0x59, 0x7d, 0x14, 0x38, 0xef, 0xb6, 0x28,
	0x04, 0x82, 0xc6, 0x78, 0x9e, 0x6b, 0x94, 0x6a, 0xe1, 0x3e, 0x9e, 0x38, 0x17, 0xfd, 0xdd, 0x2d,
	0xf6, 0xf9, 0x80, 0x0f, 0x2d, 0xf0, 0xce, 0xf6, 0x34, 0x26, 0x54, 0xc1, 0xb7, 0xa6, 0x93, 0x24,
	0xb2, 0x42, 0xf8, 0x81, 0x71, 0x9e, 0xba, 0x87, 0x66, 0x91, 0xcb, 0x65, 0xeb, 0x8f, 0xb8, 0x65,
	0xdf, 0x58, 0xf2, 0xfa, 0xed, 0xba, 0xf5, 0xcf, 0xad, 0xf3, 0xc3, 0xb9, 0x80, 0x8f, 0xd4, 0xbe,
	0x3c, 0xa5, 0x2e, 0x39, 0x4a, 0x14, 0x08, 0x94, 0xca, 0x3d, 0x32, 0xbf, 0x71, 0xd3, 0x06, 0x09,
	0x19, 0x74, 0xf9, 0x7c, 0xee, 0xbe, 0x88, 0x52, 0xd2, 0x4f, 0x05, 0x0a, 0x9b, 0x0f, 0x37, 0x35,
	0x8d, 0xc8, 0x60, 0x1b, 0xaa, 0xc9, 0xe0, 0xe4, 0xc5, 0x38, 0xb4, 0xb1, 0x87, 0x9b, 0xd8, 0xc3,
	0x4f, 0x1b, 0x45, 0x74, 0x7c, 0xdb, 0xfa, 0xbd, 0x9b, 0xdf, 0xbe, 0xc3, 0xff, 0x8d, 0x45, 0xd7,
	0xb7, 0x4b, 0xcf, 0xb9, 0x5b, 0x7a, 0xce, 0x9f, 0xa5, 0xe7, 0xdc, 0xac, 0xbc, 0xde, 0xdd, 0xca,
	0xeb, 0xfd, 0x5a, 0x79, 0xbd, 0x2f, 0x2c, 0xcb, 0x71, 0xde, 0x4c, 0xc3, 0x44, 0x96, 0xac, 0xbb,
	0xb8, 0x0a, 0x90, 0xdd, 0x5f, 0x1e, 0x2b, 0x65, 0xda, 0x14, 0xa0, 0xef, 0xef, 0x92, 0xe1, 0xa2,
	0x06, 0x3d, 0x3d, 0x34, 0xef, 0x7c, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x02, 0xb9, 0x7c, 0x2e,
	0xb5, 0x02, 0x00, 0x00,
}

func (m *Feed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Feed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Feed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.RequestContextID) > 0 {
		i -= len(m.RequestContextID)
		copy(dAtA[i:], m.RequestContextID)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.RequestContextID)))
		i--
		dAtA[i] = 0x32
	}
	if m.LatestHistory != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.LatestHistory))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ValueJsonPath) > 0 {
		i -= len(m.ValueJsonPath)
		copy(dAtA[i:], m.ValueJsonPath)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ValueJsonPath)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AggregateFunc) > 0 {
		i -= len(m.AggregateFunc)
		copy(dAtA[i:], m.AggregateFunc)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.AggregateFunc)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FeedName) > 0 {
		i -= len(m.FeedName)
		copy(dAtA[i:], m.FeedName)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.FeedName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FeedValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeedValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeedValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintOracle(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Feed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeedName)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.AggregateFunc)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.ValueJsonPath)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.LatestHistory != 0 {
		n += 1 + sovOracle(uint64(m.LatestHistory))
	}
	l = len(m.RequestContextID)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *FeedValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovOracle(uint64(l))
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Feed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Feed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Feed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeedName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeedName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregateFunc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregateFunc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueJsonPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueJsonPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHistory", wireType)
			}
			m.LatestHistory = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestHistory |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestContextID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestContextID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *FeedValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: FeedValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeedValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
