// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/gauge.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/osmosis-labs/osmosis/v17/x/lockup/types"
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

// SplittingPolicy determines the way we want to split incentives in groupGauges
type SplittingPolicy int32

const (
	Volume SplittingPolicy = 0
	Evenly SplittingPolicy = 1
)

var SplittingPolicy_name = map[int32]string{
	0: "Volume",
	1: "Evenly",
}

var SplittingPolicy_value = map[string]int32{
	"Volume": 0,
	"Evenly": 1,
}

func (x SplittingPolicy) String() string {
	return proto.EnumName(SplittingPolicy_name, int32(x))
}

func (SplittingPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0304e2bb0159901, []int{0}
}

// Gauge is an object that stores and distributes yields to recipients who
// satisfy certain conditions. Currently gauges support conditions around the
// duration for which a given denom is locked.
type Gauge struct {
	// id is the unique ID of a Gauge
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// is_perpetual is a flag to show if it's a perpetual or non-perpetual gauge
	// Non-perpetual gauges distribute their tokens equally per epoch while the
	// gauge is in the active period. Perpetual gauges distribute all their tokens
	// at a single time and only distribute their tokens again once the gauge is
	// refilled, Intended for use with incentives that get refilled daily.
	IsPerpetual bool `protobuf:"varint,2,opt,name=is_perpetual,json=isPerpetual,proto3" json:"is_perpetual,omitempty"`
	// distribute_to is where the gauge rewards are distributed to.
	// This is queried via lock duration or by timestamp
	DistributeTo types.QueryCondition `protobuf:"bytes,3,opt,name=distribute_to,json=distributeTo,proto3" json:"distribute_to"`
	// coins is the total amount of coins that have been in the gauge
	// Can distribute multiple coin denoms
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// start_time is the distribution start time
	StartTime time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// num_epochs_paid_over is the number of total epochs distribution will be
	// completed over
	NumEpochsPaidOver uint64 `protobuf:"varint,6,opt,name=num_epochs_paid_over,json=numEpochsPaidOver,proto3" json:"num_epochs_paid_over,omitempty"`
	// filled_epochs is the number of epochs distribution has been completed on
	// already
	FilledEpochs uint64 `protobuf:"varint,7,opt,name=filled_epochs,json=filledEpochs,proto3" json:"filled_epochs,omitempty"`
	// distributed_coins are coins that have been distributed already
	DistributedCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,8,rep,name=distributed_coins,json=distributedCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"distributed_coins"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}
func (*Gauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0304e2bb0159901, []int{0}
}
func (m *Gauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Gauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Gauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Gauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gauge.Merge(m, src)
}
func (m *Gauge) XXX_Size() int {
	return m.Size()
}
func (m *Gauge) XXX_DiscardUnknown() {
	xxx_messageInfo_Gauge.DiscardUnknown(m)
}

var xxx_messageInfo_Gauge proto.InternalMessageInfo

func (m *Gauge) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Gauge) GetIsPerpetual() bool {
	if m != nil {
		return m.IsPerpetual
	}
	return false
}

func (m *Gauge) GetDistributeTo() types.QueryCondition {
	if m != nil {
		return m.DistributeTo
	}
	return types.QueryCondition{}
}

func (m *Gauge) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Gauge) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Gauge) GetNumEpochsPaidOver() uint64 {
	if m != nil {
		return m.NumEpochsPaidOver
	}
	return 0
}

func (m *Gauge) GetFilledEpochs() uint64 {
	if m != nil {
		return m.FilledEpochs
	}
	return 0
}

func (m *Gauge) GetDistributedCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DistributedCoins
	}
	return nil
}

// Gauge is an object that stores GroupGaugeId as well as internalGaugeIds. We
// linked these two together so that we can distribute tokens from groupGauge to
// internalGauges.
type GroupGauge struct {
	GroupGaugeId    uint64          `protobuf:"varint,1,opt,name=group_gauge_id,json=groupGaugeId,proto3" json:"group_gauge_id,omitempty"`
	InternalIds     []uint64        `protobuf:"varint,2,rep,packed,name=internal_ids,json=internalIds,proto3" json:"internal_ids,omitempty"`
	SplittingPolicy SplittingPolicy `protobuf:"varint,3,opt,name=splitting_policy,json=splittingPolicy,proto3,enum=osmosis.incentives.SplittingPolicy" json:"splitting_policy,omitempty"`
}

func (m *GroupGauge) Reset()         { *m = GroupGauge{} }
func (m *GroupGauge) String() string { return proto.CompactTextString(m) }
func (*GroupGauge) ProtoMessage()    {}
func (*GroupGauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0304e2bb0159901, []int{1}
}
func (m *GroupGauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupGauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupGauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupGauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupGauge.Merge(m, src)
}
func (m *GroupGauge) XXX_Size() int {
	return m.Size()
}
func (m *GroupGauge) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupGauge.DiscardUnknown(m)
}

var xxx_messageInfo_GroupGauge proto.InternalMessageInfo

func (m *GroupGauge) GetGroupGaugeId() uint64 {
	if m != nil {
		return m.GroupGaugeId
	}
	return 0
}

func (m *GroupGauge) GetInternalIds() []uint64 {
	if m != nil {
		return m.InternalIds
	}
	return nil
}

func (m *GroupGauge) GetSplittingPolicy() SplittingPolicy {
	if m != nil {
		return m.SplittingPolicy
	}
	return Volume
}

type LockableDurationsInfo struct {
	// List of incentivised durations that gauges will pay out to
	LockableDurations []time.Duration `protobuf:"bytes,1,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
}

func (m *LockableDurationsInfo) Reset()         { *m = LockableDurationsInfo{} }
func (m *LockableDurationsInfo) String() string { return proto.CompactTextString(m) }
func (*LockableDurationsInfo) ProtoMessage()    {}
func (*LockableDurationsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0304e2bb0159901, []int{2}
}
func (m *LockableDurationsInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LockableDurationsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockableDurationsInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LockableDurationsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockableDurationsInfo.Merge(m, src)
}
func (m *LockableDurationsInfo) XXX_Size() int {
	return m.Size()
}
func (m *LockableDurationsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LockableDurationsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LockableDurationsInfo proto.InternalMessageInfo

func (m *LockableDurationsInfo) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

func init() {
	proto.RegisterEnum("osmosis.incentives.SplittingPolicy", SplittingPolicy_name, SplittingPolicy_value)
	proto.RegisterType((*Gauge)(nil), "osmosis.incentives.Gauge")
	proto.RegisterType((*GroupGauge)(nil), "osmosis.incentives.GroupGauge")
	proto.RegisterType((*LockableDurationsInfo)(nil), "osmosis.incentives.LockableDurationsInfo")
}

func init() { proto.RegisterFile("osmosis/incentives/gauge.proto", fileDescriptor_c0304e2bb0159901) }

var fileDescriptor_c0304e2bb0159901 = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x4f, 0xdb, 0x3e,
	0x18, 0x6e, 0x4a, 0xe1, 0xc7, 0xcf, 0x2d, 0xd0, 0x5a, 0x4c, 0x0a, 0x95, 0x96, 0x76, 0x65, 0x93,
	0xaa, 0x49, 0xd8, 0x83, 0x49, 0x9b, 0xb4, 0x63, 0x19, 0x42, 0x95, 0xa6, 0xad, 0xeb, 0xd0, 0x34,
	0xed, 0x12, 0x39, 0x89, 0x09, 0x16, 0x4e, 0x1c, 0xc5, 0x4e, 0x45, 0xbf, 0x01, 0x47, 0xb4, 0xd3,
	0xee, 0xd3, 0x2e, 0xfb, 0x24, 0x1c, 0x39, 0xee, 0x04, 0x13, 0x7c, 0x83, 0x7d, 0x82, 0x29, 0x76,
	0x42, 0x59, 0xb9, 0xee, 0x14, 0xfb, 0x7d, 0xde, 0x7f, 0xcf, 0xf3, 0xbe, 0x0e, 0x70, 0x84, 0x8c,
	0x84, 0x64, 0x12, 0xb3, 0xd8, 0xa7, 0xb1, 0x62, 0x13, 0x2a, 0x71, 0x48, 0xb2, 0x90, 0xa2, 0x24,
	0x15, 0x4a, 0x40, 0x58, 0xe0, 0x68, 0x86, 0xb7, 0xd7, 0x43, 0x11, 0x0a, 0x0d, 0xe3, 0xfc, 0x64,
	0x3c, 0xdb, 0x4e, 0x28, 0x44, 0xc8, 0x29, 0xd6, 0x37, 0x2f, 0x3b, 0xc4, 0x41, 0x96, 0x12, 0xc5,
	0x44, 0x5c, 0xe0, 0x9d, 0x79, 0x5c, 0xb1, 0x88, 0x4a, 0x45, 0xa2, 0xa4, 0x4c, 0xe0, 0xeb, 0x5a,
	0xd8, 0x23, 0x92, 0xe2, 0xc9, 0xb6, 0x47, 0x15, 0xd9, 0xc6, 0xbe, 0x60, 0x65, 0x82, 0x8d, 0xb2,
	0x55, 0x2e, 0xfc, 0xe3, 0x2c, 0xd1, 0x1f, 0x03, 0xf5, 0xbe, 0xd4, 0xc0, 0xe2, 0x7e, 0xde, 0x35,
	0x5c, 0x05, 0x55, 0x16, 0xd8, 0x56, 0xd7, 0xea, 0xd7, 0xc6, 0x55, 0x16, 0xc0, 0x47, 0xa0, 0xc1,
	0xa4, 0x9b, 0xd0, 0x34, 0xa1, 0x2a, 0x23, 0xdc, 0xae, 0x76, 0xad, 0xfe, 0xf2, 0xb8, 0xce, 0xe4,
	0xa8, 0x34, 0xc1, 0x21, 0x58, 0x09, 0x98, 0x54, 0x29, 0xf3, 0x32, 0x45, 0x5d, 0x25, 0xec, 0x85,
	0xae, 0xd5, 0xaf, 0xef, 0x38, 0xa8, 0xa4, 0x6e, 0xea, 0xa1, 0xf7, 0x19, 0x4d, 0xa7, 0xbb, 0x22,
	0x0e, 0x58, 0xce, 0x6a, 0x50, 0x3b, 0xbf, 0xec, 0x54, 0xc6, 0x8d, 0x59, 0xe8, 0x81, 0x80, 0x04,
	0x2c, 0xe6, 0x0d, 0x4b, 0xbb, 0xd6, 0x5d, 0xe8, 0xd7, 0x77, 0x36, 0x90, 0xa1, 0x84, 0x72, 0x4a,
	0xa8, 0xa0, 0x84, 0x76, 0x05, 0x8b, 0x07, 0xcf, 0xf2, 0xe8, 0x1f, 0x57, 0x9d, 0x7e, 0xc8, 0xd4,
	0x51, 0xe6, 0x21, 0x5f, 0x44, 0xb8, 0xe0, 0x6f, 0x3e, 0x5b, 0x32, 0x38, 0xc6, 0x6a, 0x9a, 0x50,
	0xa9, 0x03, 0xe4, 0xd8, 0x64, 0x86, 0x9f, 0x00, 0x90, 0x8a, 0xa4, 0xca, 0xcd, 0xe5, 0xb3, 0x17,
	0x75, 0xab, 0x6d, 0x64, 0xb4, 0x45, 0xa5, 0xb6, 0xe8, 0xa0, 0xd4, 0x76, 0xf0, 0x30, 0x2f, 0xf4,
	0xfb, 0xb2, 0xd3, 0x9a, 0x92, 0x88, 0xbf, 0xea, 0xcd, 0x62, 0x7b, 0x67, 0x57, 0x1d, 0x6b, 0xfc,
	0xbf, 0x36, 0xe4, 0xee, 0x10, 0x83, 0xf5, 0x38, 0x8b, 0x5c, 0x9a, 0x08, 0xff, 0x48, 0xba, 0x09,
	0x61, 0x81, 0x2b, 0x26, 0x34, 0xb5, 0x97, 0xb4, 0x98, 0xad, 0x38, 0x8b, 0xf6, 0x34, 0x34, 0x22,
	0x2c, 0x78, 0x37, 0xa1, 0x29, 0xdc, 0x04, 0x2b, 0x87, 0x8c, 0x73, 0x1a, 0x14, 0x31, 0xf6, 0x7f,
	0xda, 0xb3, 0x61, 0x8c, 0xc6, 0x19, 0x9e, 0x80, 0xd6, 0x4c, 0xa2, 0xc0, 0x35, 0xf2, 0x2c, 0xff,
	0x7b, 0x79, 0x9a, 0x77, 0xaa, 0x68, 0x4b, 0xef, 0xbb, 0x05, 0xc0, 0x7e, 0x2a, 0xb2, 0xc4, 0x6c,
	0xc6, 0x63, 0xb0, 0x1a, 0xe6, 0x37, 0x57, 0xaf, 0xb7, 0x7b, 0xbb, 0x25, 0x8d, 0xf0, 0xd6, 0x67,
	0x68, 0xf6, 0x25, 0x56, 0x34, 0x8d, 0x09, 0x77, 0x59, 0x20, 0xed, 0x6a, 0x77, 0xa1, 0x5f, 0x1b,
	0xd7, 0x4b, 0xdb, 0x30, 0x90, 0xf0, 0x2d, 0x68, 0xca, 0x84, 0x33, 0xa5, 0x58, 0x1c, 0xba, 0x89,
	0xe0, 0xcc, 0x9f, 0xea, 0x95, 0x59, 0xdd, 0xd9, 0x44, 0xf7, 0x5f, 0x0b, 0xfa, 0x50, 0xfa, 0x8e,
	0xb4, 0xeb, 0x78, 0x4d, 0xfe, 0x6d, 0xe8, 0x9d, 0x5a, 0xe0, 0xc1, 0x1b, 0xe1, 0x1f, 0x13, 0x8f,
	0xd3, 0xd7, 0xc5, 0x9b, 0x91, 0xc3, 0xf8, 0x50, 0x40, 0x01, 0x20, 0x2f, 0x00, 0xb7, 0x7c, 0x4d,
	0xd2, 0xb6, 0x0a, 0xf1, 0xe6, 0x67, 0x5e, 0xc6, 0x0e, 0x9e, 0x14, 0x23, 0xdf, 0x30, 0x23, 0xbf,
	0x9f, 0xa2, 0xf7, 0x35, 0x1f, 0x7d, 0x8b, 0xcf, 0x17, 0x7d, 0x8a, 0xc1, 0xda, 0x5c, 0xbb, 0x10,
	0x80, 0xa5, 0x8f, 0x82, 0x67, 0x11, 0x6d, 0x56, 0xf2, 0xf3, 0xde, 0x84, 0xc6, 0x7c, 0xda, 0xb4,
	0xda, 0xb5, 0xd3, 0x6f, 0x4e, 0x65, 0x30, 0x3a, 0xbf, 0x76, 0xac, 0x8b, 0x6b, 0xc7, 0xfa, 0x75,
	0xed, 0x58, 0x67, 0x37, 0x4e, 0xe5, 0xe2, 0xc6, 0xa9, 0xfc, 0xbc, 0x71, 0x2a, 0x9f, 0x5f, 0xdc,
	0x99, 0x5c, 0xa1, 0xca, 0x16, 0x27, 0x9e, 0x2c, 0x2f, 0x78, 0xb2, 0xfd, 0x12, 0x9f, 0xdc, 0xfd,
	0xed, 0xe8, 0x69, 0x7a, 0x4b, 0x9a, 0xcf, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0x40,
	0x27, 0xeb, 0x99, 0x04, 0x00, 0x00,
}

func (m *Gauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Gauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DistributedCoins) > 0 {
		for iNdEx := len(m.DistributedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DistributedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGauge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.FilledEpochs != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.FilledEpochs))
		i--
		dAtA[i] = 0x38
	}
	if m.NumEpochsPaidOver != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.NumEpochsPaidOver))
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGauge(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGauge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.DistributeTo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGauge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.IsPerpetual {
		i--
		if m.IsPerpetual {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GroupGauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupGauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupGauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SplittingPolicy != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.SplittingPolicy))
		i--
		dAtA[i] = 0x18
	}
	if len(m.InternalIds) > 0 {
		dAtA4 := make([]byte, len(m.InternalIds)*10)
		var j3 int
		for _, num := range m.InternalIds {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintGauge(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if m.GroupGaugeId != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.GroupGaugeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LockableDurationsInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockableDurationsInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockableDurationsInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LockableDurations) > 0 {
		for iNdEx := len(m.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockableDurations[iNdEx], dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockableDurations[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintGauge(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGauge(dAtA []byte, offset int, v uint64) int {
	offset -= sovGauge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Gauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGauge(uint64(m.Id))
	}
	if m.IsPerpetual {
		n += 2
	}
	l = m.DistributeTo.Size()
	n += 1 + l + sovGauge(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovGauge(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovGauge(uint64(l))
	if m.NumEpochsPaidOver != 0 {
		n += 1 + sovGauge(uint64(m.NumEpochsPaidOver))
	}
	if m.FilledEpochs != 0 {
		n += 1 + sovGauge(uint64(m.FilledEpochs))
	}
	if len(m.DistributedCoins) > 0 {
		for _, e := range m.DistributedCoins {
			l = e.Size()
			n += 1 + l + sovGauge(uint64(l))
		}
	}
	return n
}

func (m *GroupGauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupGaugeId != 0 {
		n += 1 + sovGauge(uint64(m.GroupGaugeId))
	}
	if len(m.InternalIds) > 0 {
		l = 0
		for _, e := range m.InternalIds {
			l += sovGauge(uint64(e))
		}
		n += 1 + sovGauge(uint64(l)) + l
	}
	if m.SplittingPolicy != 0 {
		n += 1 + sovGauge(uint64(m.SplittingPolicy))
	}
	return n
}

func (m *LockableDurationsInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovGauge(uint64(l))
		}
	}
	return n
}

func sovGauge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGauge(x uint64) (n int) {
	return sovGauge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Gauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGauge
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
			return fmt.Errorf("proto: Gauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPerpetual", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
			m.IsPerpetual = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeTo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributeTo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types1.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEpochsPaidOver", wireType)
			}
			m.NumEpochsPaidOver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumEpochsPaidOver |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilledEpochs", wireType)
			}
			m.FilledEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FilledEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributedCoins = append(m.DistributedCoins, types1.Coin{})
			if err := m.DistributedCoins[len(m.DistributedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGauge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGauge
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
func (m *GroupGauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGauge
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
			return fmt.Errorf("proto: GroupGauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupGauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupGaugeId", wireType)
			}
			m.GroupGaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupGaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGauge
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.InternalIds = append(m.InternalIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGauge
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGauge
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGauge
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.InternalIds) == 0 {
					m.InternalIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGauge
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.InternalIds = append(m.InternalIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalIds", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplittingPolicy", wireType)
			}
			m.SplittingPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SplittingPolicy |= SplittingPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGauge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGauge
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
func (m *LockableDurationsInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGauge
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
			return fmt.Errorf("proto: LockableDurationsInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockableDurationsInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockableDurations = append(m.LockableDurations, time.Duration(0))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&(m.LockableDurations[len(m.LockableDurations)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGauge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGauge
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
func skipGauge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGauge
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
					return 0, ErrIntOverflowGauge
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
					return 0, ErrIntOverflowGauge
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
				return 0, ErrInvalidLengthGauge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGauge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGauge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGauge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGauge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGauge = fmt.Errorf("proto: unexpected end of group")
)
