// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: farm/farm.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type FarmPool struct {
	Name                   string                                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Creator                string                                  `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Description            string                                  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	BeginHeight            uint64                                  `protobuf:"varint,4,opt,name=begin_height,json=beginHeight,proto3" json:"begin_height,omitempty"`
	EndHeight              uint64                                  `protobuf:"varint,5,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	LastHeightDistrRewards uint64                                  `protobuf:"varint,6,opt,name=last_height_distr_rewards,json=lastHeightDistrRewards,proto3" json:"last_height_distr_rewards,omitempty"`
	Destructible           bool                                    `protobuf:"varint,7,opt,name=destructible,proto3" json:"destructible,omitempty"`
	TotalLpTokenLocked     github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,8,opt,name=total_lp_token_locked,json=totalLpTokenLocked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"total_lp_token_locked"`
	Rules                  []*RewardRule                           `protobuf:"bytes,9,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (m *FarmPool) Reset()         { *m = FarmPool{} }
func (m *FarmPool) String() string { return proto.CompactTextString(m) }
func (*FarmPool) ProtoMessage()    {}
func (*FarmPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85c74c264ccc821, []int{0}
}
func (m *FarmPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FarmPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FarmPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FarmPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FarmPool.Merge(m, src)
}
func (m *FarmPool) XXX_Size() int {
	return m.Size()
}
func (m *FarmPool) XXX_DiscardUnknown() {
	xxx_messageInfo_FarmPool.DiscardUnknown(m)
}

var xxx_messageInfo_FarmPool proto.InternalMessageInfo

type RewardRule struct {
	Reward          string                                 `protobuf:"bytes,1,opt,name=reward,proto3" json:"reward,omitempty"`
	TotalReward     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=total_reward,json=totalReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_reward"`
	RemainingReward github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=remaining_reward,json=remainingReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"remaining_reward"`
	RewardPerBlock  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=reward_per_block,json=rewardPerBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reward_per_block"`
	RewardPerShare  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reward_per_share,json=rewardPerShare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_per_share"`
}

func (m *RewardRule) Reset()         { *m = RewardRule{} }
func (m *RewardRule) String() string { return proto.CompactTextString(m) }
func (*RewardRule) ProtoMessage()    {}
func (*RewardRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85c74c264ccc821, []int{1}
}
func (m *RewardRule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardRule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardRule.Merge(m, src)
}
func (m *RewardRule) XXX_Size() int {
	return m.Size()
}
func (m *RewardRule) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardRule.DiscardUnknown(m)
}

var xxx_messageInfo_RewardRule proto.InternalMessageInfo

type FarmInfo struct {
	PoolName   string                                   `protobuf:"bytes,1,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	Address    string                                   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Locked     github_com_cosmos_cosmos_sdk_types.Int   `protobuf:"bytes,3,opt,name=locked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"locked"`
	RewardDebt github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=reward_debt,json=rewardDebt,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"reward_debt"`
}

func (m *FarmInfo) Reset()         { *m = FarmInfo{} }
func (m *FarmInfo) String() string { return proto.CompactTextString(m) }
func (*FarmInfo) ProtoMessage()    {}
func (*FarmInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a85c74c264ccc821, []int{2}
}
func (m *FarmInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FarmInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FarmInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FarmInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FarmInfo.Merge(m, src)
}
func (m *FarmInfo) XXX_Size() int {
	return m.Size()
}
func (m *FarmInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FarmInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FarmInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FarmPool)(nil), "irismod.farm.FarmPool")
	proto.RegisterType((*RewardRule)(nil), "irismod.farm.RewardRule")
	proto.RegisterType((*FarmInfo)(nil), "irismod.farm.FarmInfo")
}

func init() { proto.RegisterFile("farm/farm.proto", fileDescriptor_a85c74c264ccc821) }

var fileDescriptor_a85c74c264ccc821 = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xbd, 0x6e, 0x13, 0x41,
	0x10, 0xf6, 0x61, 0xc7, 0xb1, 0xd7, 0x16, 0x41, 0x2b, 0x88, 0x2e, 0x41, 0x9c, 0x4d, 0x0a, 0x70,
	0x93, 0x3b, 0x12, 0x2a, 0x28, 0x4d, 0x14, 0x11, 0x11, 0xa1, 0x70, 0x50, 0x00, 0xcd, 0xe9, 0xee,
	0x76, 0x72, 0x5e, 0xe5, 0x6e, 0xf7, 0xb4, 0xbb, 0x06, 0x51, 0xf0, 0x06, 0x14, 0x3c, 0x02, 0x35,
	0x15, 0x8f, 0x91, 0x32, 0x25, 0xa2, 0x08, 0x90, 0x34, 0xbc, 0x04, 0x12, 0xda, 0x1f, 0x83, 0xa1,
	0x40, 0x49, 0x1a, 0x7b, 0xf6, 0x9b, 0xb9, 0x6f, 0x76, 0x66, 0xbe, 0x59, 0xb4, 0xb4, 0x9f, 0x8a,
	0x2a, 0xd2, 0x3f, 0x61, 0x2d, 0xb8, 0xe2, 0xb8, 0x4f, 0x05, 0x95, 0x15, 0x27, 0xa1, 0xc6, 0x56,
	0x83, 0x9c, 0xcb, 0x8a, 0xcb, 0x28, 0x4b, 0x25, 0x44, 0xaf, 0x36, 0x32, 0x50, 0xe9, 0x46, 0x94,
	0x73, 0xca, 0x6c, 0xf4, 0xea, 0xd5, 0x82, 0x17, 0xdc, 0x98, 0x91, 0xb6, 0x2c, 0xba, 0xf6, 0xa9,
	0x89, 0x3a, 0xdb, 0xa9, 0xa8, 0xf6, 0x38, 0x2f, 0x31, 0x46, 0x2d, 0x96, 0x56, 0xe0, 0x7b, 0x43,
	0x6f, 0xd4, 0x8d, 0x8d, 0x8d, 0x7d, 0xb4, 0x98, 0x0b, 0x48, 0x15, 0x17, 0xfe, 0x25, 0x03, 0xcf,
	0x8e, 0x78, 0x88, 0x7a, 0x04, 0x64, 0x2e, 0x68, 0xad, 0x28, 0x67, 0x7e, 0xd3, 0x78, 0xe7, 0x21,
	0x7c, 0x13, 0xf5, 0x33, 0x28, 0x28, 0x4b, 0x26, 0x40, 0x8b, 0x89, 0xf2, 0x5b, 0x43, 0x6f, 0xd4,
	0x8a, 0x7b, 0x06, 0x7b, 0x68, 0x20, 0x7c, 0x03, 0x21, 0x60, 0x64, 0x16, 0xb0, 0x60, 0x02, 0xba,
	0xc0, 0x88, 0x73, 0xdf, 0x43, 0x2b, 0x65, 0x2a, 0x95, 0xf3, 0x27, 0x84, 0x4a, 0x25, 0x12, 0x01,
	0xaf, 0x53, 0x41, 0xa4, 0xdf, 0x36, 0xd1, 0xcb, 0x3a, 0xc0, 0x86, 0x6f, 0x69, 0x77, 0x6c, 0xbd,
	0x78, 0x0d, 0xf5, 0x09, 0x48, 0x25, 0xa6, 0xb9, 0xa2, 0x59, 0x09, 0xfe, 0xe2, 0xd0, 0x1b, 0x75,
	0xe2, 0xbf, 0x30, 0xfc, 0x16, 0x5d, 0x53, 0x5c, 0xa5, 0x65, 0x52, 0xd6, 0x89, 0xe2, 0x07, 0xc0,
	0x92, 0x92, 0xe7, 0x07, 0x40, 0xfc, 0xce, 0xd0, 0x1b, 0xf5, 0x36, 0x57, 0x42, 0xdb, 0xd3, 0x50,
	0xf7, 0x34, 0x74, 0x3d, 0x0d, 0x1f, 0x70, 0xca, 0xc6, 0xd1, 0xe1, 0xf1, 0xa0, 0xf1, 0xe5, 0x78,
	0x70, 0xbb, 0xa0, 0x6a, 0x32, 0xcd, 0xc2, 0x9c, 0x57, 0x91, 0x1b, 0x80, 0xfd, 0x5b, 0x97, 0xe4,
	0x20, 0x52, 0x6f, 0x6a, 0x90, 0xe6, 0x83, 0x18, 0x9b, 0x44, 0xbb, 0xf5, 0x33, 0x9d, 0x66, 0xd7,
	0x64, 0xc1, 0x21, 0x5a, 0x10, 0xd3, 0x12, 0xa4, 0xdf, 0x1d, 0x36, 0x47, 0xbd, 0x4d, 0x3f, 0x9c,
	0x1f, 0x68, 0x68, 0x0b, 0x89, 0xa7, 0x25, 0xc4, 0x36, 0xec, 0x7e, 0xeb, 0xc7, 0x87, 0x81, 0xb7,
	0xf6, 0xae, 0x89, 0xd0, 0x1f, 0x1f, 0x5e, 0x46, 0x6d, 0xdb, 0x10, 0x37, 0x36, 0x77, 0xc2, 0x4f,
	0x50, 0xdf, 0xd6, 0xe6, 0xbc, 0x66, 0x7a, 0xe3, 0xd0, 0xdd, 0xfb, 0xd6, 0x19, 0xee, 0xbd, 0xc3,
	0x54, 0xdc, 0x33, 0x1c, 0x36, 0x1d, 0x7e, 0x81, 0xae, 0x08, 0xa8, 0x52, 0xca, 0x28, 0x2b, 0x66,
	0xb4, 0xcd, 0x0b, 0xd1, 0x2e, 0xfd, 0xe6, 0x71, 0xd4, 0xcf, 0x35, 0xb5, 0xb6, 0x92, 0x1a, 0x44,
	0x92, 0xe9, 0x31, 0x18, 0xb9, 0x9c, 0x9f, 0xfa, 0xb2, 0xe5, 0xd9, 0x03, 0x31, 0xd6, 0x2c, 0xff,
	0x30, 0xcb, 0x49, 0x2a, 0xc0, 0xe8, 0xec, 0x7c, 0xcc, 0x5b, 0x90, 0xcf, 0x31, 0x3f, 0xd5, 0x2c,
	0x6e, 0x1c, 0x3f, 0x3d, 0xbb, 0x41, 0x3b, 0x6c, 0x9f, 0xe3, 0xeb, 0xa8, 0x5b, 0x73, 0x5e, 0x26,
	0x73, 0x6b, 0xd4, 0xd1, 0xc0, 0x63, 0xb7, 0x4a, 0x29, 0x21, 0x02, 0xa4, 0x9c, 0xad, 0x92, 0x3b,
	0xe2, 0x6d, 0xd4, 0x76, 0xc2, 0xbb, 0x58, 0x3b, 0xdd, 0xd7, 0xb8, 0x44, 0x3d, 0x57, 0x2b, 0x81,
	0x4c, 0xef, 0x5b, 0xf3, 0xff, 0x2a, 0xbe, 0xa3, 0xf3, 0x7c, 0xfc, 0x3a, 0x18, 0x9d, 0x51, 0xc5,
	0x32, 0x46, 0x96, 0x7f, 0x0b, 0x32, 0x65, 0xeb, 0x1f, 0x3f, 0x3a, 0xfc, 0x1e, 0x34, 0x0e, 0x4f,
	0x02, 0xef, 0xe8, 0x24, 0xf0, 0xbe, 0x9d, 0x04, 0xde, 0xfb, 0xd3, 0xa0, 0x71, 0x74, 0x1a, 0x34,
	0x3e, 0x9f, 0x06, 0x8d, 0x97, 0xeb, 0x73, 0xcc, 0x5a, 0xdd, 0x0c, 0x54, 0xe4, 0x54, 0x1e, 0x55,
	0x9c, 0x68, 0x59, 0x9b, 0x27, 0xcd, 0x26, 0xc9, 0xda, 0xe6, 0x55, 0xba, 0xfb, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0xe5, 0x53, 0xf6, 0xca, 0xec, 0x04, 0x00, 0x00,
}

func (this *FarmPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FarmPool)
	if !ok {
		that2, ok := that.(FarmPool)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.BeginHeight != that1.BeginHeight {
		return false
	}
	if this.EndHeight != that1.EndHeight {
		return false
	}
	if this.LastHeightDistrRewards != that1.LastHeightDistrRewards {
		return false
	}
	if this.Destructible != that1.Destructible {
		return false
	}
	if !this.TotalLpTokenLocked.Equal(that1.TotalLpTokenLocked) {
		return false
	}
	if len(this.Rules) != len(that1.Rules) {
		return false
	}
	for i := range this.Rules {
		if !this.Rules[i].Equal(that1.Rules[i]) {
			return false
		}
	}
	return true
}
func (this *RewardRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RewardRule)
	if !ok {
		that2, ok := that.(RewardRule)
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
	if this.Reward != that1.Reward {
		return false
	}
	if !this.TotalReward.Equal(that1.TotalReward) {
		return false
	}
	if !this.RemainingReward.Equal(that1.RemainingReward) {
		return false
	}
	if !this.RewardPerBlock.Equal(that1.RewardPerBlock) {
		return false
	}
	if !this.RewardPerShare.Equal(that1.RewardPerShare) {
		return false
	}
	return true
}
func (this *FarmInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FarmInfo)
	if !ok {
		that2, ok := that.(FarmInfo)
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
	if this.PoolName != that1.PoolName {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if !this.Locked.Equal(that1.Locked) {
		return false
	}
	if len(this.RewardDebt) != len(that1.RewardDebt) {
		return false
	}
	for i := range this.RewardDebt {
		if !this.RewardDebt[i].Equal(&that1.RewardDebt[i]) {
			return false
		}
	}
	return true
}
func (m *FarmPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FarmPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FarmPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for iNdEx := len(m.Rules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFarm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	{
		size := m.TotalLpTokenLocked.Size()
		i -= size
		if _, err := m.TotalLpTokenLocked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFarm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.Destructible {
		i--
		if m.Destructible {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.LastHeightDistrRewards != 0 {
		i = encodeVarintFarm(dAtA, i, uint64(m.LastHeightDistrRewards))
		i--
		dAtA[i] = 0x30
	}
	if m.EndHeight != 0 {
		i = encodeVarintFarm(dAtA, i, uint64(m.EndHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.BeginHeight != 0 {
		i = encodeVarintFarm(dAtA, i, uint64(m.BeginHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintFarm(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintFarm(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintFarm(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RewardRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardRule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardRule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RewardPerShare.Size()
		i -= size
		if _, err := m.RewardPerShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFarm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.RewardPerBlock.Size()
		i -= size
		if _, err := m.RewardPerBlock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFarm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.RemainingReward.Size()
		i -= size
		if _, err := m.RemainingReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFarm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.TotalReward.Size()
		i -= size
		if _, err := m.TotalReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFarm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Reward) > 0 {
		i -= len(m.Reward)
		copy(dAtA[i:], m.Reward)
		i = encodeVarintFarm(dAtA, i, uint64(len(m.Reward)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FarmInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FarmInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FarmInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RewardDebt) > 0 {
		for iNdEx := len(m.RewardDebt) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RewardDebt[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFarm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.Locked.Size()
		i -= size
		if _, err := m.Locked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFarm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintFarm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolName) > 0 {
		i -= len(m.PoolName)
		copy(dAtA[i:], m.PoolName)
		i = encodeVarintFarm(dAtA, i, uint64(len(m.PoolName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFarm(dAtA []byte, offset int, v uint64) int {
	offset -= sovFarm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FarmPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFarm(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovFarm(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovFarm(uint64(l))
	}
	if m.BeginHeight != 0 {
		n += 1 + sovFarm(uint64(m.BeginHeight))
	}
	if m.EndHeight != 0 {
		n += 1 + sovFarm(uint64(m.EndHeight))
	}
	if m.LastHeightDistrRewards != 0 {
		n += 1 + sovFarm(uint64(m.LastHeightDistrRewards))
	}
	if m.Destructible {
		n += 2
	}
	l = m.TotalLpTokenLocked.Size()
	n += 1 + l + sovFarm(uint64(l))
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovFarm(uint64(l))
		}
	}
	return n
}

func (m *RewardRule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reward)
	if l > 0 {
		n += 1 + l + sovFarm(uint64(l))
	}
	l = m.TotalReward.Size()
	n += 1 + l + sovFarm(uint64(l))
	l = m.RemainingReward.Size()
	n += 1 + l + sovFarm(uint64(l))
	l = m.RewardPerBlock.Size()
	n += 1 + l + sovFarm(uint64(l))
	l = m.RewardPerShare.Size()
	n += 1 + l + sovFarm(uint64(l))
	return n
}

func (m *FarmInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolName)
	if l > 0 {
		n += 1 + l + sovFarm(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovFarm(uint64(l))
	}
	l = m.Locked.Size()
	n += 1 + l + sovFarm(uint64(l))
	if len(m.RewardDebt) > 0 {
		for _, e := range m.RewardDebt {
			l = e.Size()
			n += 1 + l + sovFarm(uint64(l))
		}
	}
	return n
}

func sovFarm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFarm(x uint64) (n int) {
	return sovFarm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FarmPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFarm
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
			return fmt.Errorf("proto: FarmPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FarmPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginHeight", wireType)
			}
			m.BeginHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BeginHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndHeight", wireType)
			}
			m.EndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeightDistrRewards", wireType)
			}
			m.LastHeightDistrRewards = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeightDistrRewards |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destructible", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
			m.Destructible = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalLpTokenLocked", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalLpTokenLocked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, &RewardRule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFarm
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFarm
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
func (m *RewardRule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFarm
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
			return fmt.Errorf("proto: RewardRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reward = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPerBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPerBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPerShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPerShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFarm
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFarm
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
func (m *FarmInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFarm
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
			return fmt.Errorf("proto: FarmInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FarmInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Locked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDebt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFarm
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
				return ErrInvalidLengthFarm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFarm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardDebt = append(m.RewardDebt, types.Coin{})
			if err := m.RewardDebt[len(m.RewardDebt)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFarm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFarm
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFarm
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
func skipFarm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFarm
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
					return 0, ErrIntOverflowFarm
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
					return 0, ErrIntOverflowFarm
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
				return 0, ErrInvalidLengthFarm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFarm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFarm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFarm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFarm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFarm = fmt.Errorf("proto: unexpected end of group")
)
