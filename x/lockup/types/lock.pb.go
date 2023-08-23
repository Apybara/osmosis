// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/lockup/lock.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

// LockQueryType defines the type of the lock query that can
// either be by duration or start time of the lock.
type LockQueryType int32

const (
	ByDuration LockQueryType = 0
	ByTime     LockQueryType = 1
	NoLock     LockQueryType = 2
	ByGroup    LockQueryType = 3
)

var LockQueryType_name = map[int32]string{
	0: "ByDuration",
	1: "ByTime",
	2: "NoLock",
	3: "ByGroup",
}

var LockQueryType_value = map[string]int32{
	"ByDuration": 0,
	"ByTime":     1,
	"NoLock":     2,
	"ByGroup":    3,
}

func (x LockQueryType) String() string {
	return proto.EnumName(LockQueryType_name, int32(x))
}

func (LockQueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7e9d7527a237b489, []int{0}
}

// PeriodLock is a single lock unit by period defined by the x/lockup module.
// It's a record of a locked coin at a specific time. It stores owner, duration,
// unlock time and the number of coins locked. A state of a period lock is
// created upon lock creation, and deleted once the lock has been matured after
// the `duration` has passed since unbonding started.
type PeriodLock struct {
	// ID is the unique id of the lock.
	// The ID of the lock is decided upon lock creation, incrementing by 1 for
	// every lock.
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Owner is the account address of the lock owner.
	// Only the owner can modify the state of the lock.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// Duration is the time needed for a lock to mature after unlocking has
	// started.
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	// EndTime refers to the time at which the lock would mature and get deleted.
	// This value is first initialized when an unlock has started for the lock,
	// end time being block time + duration.
	EndTime time.Time `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	// Coins are the tokens locked within the lock, kept in the module account.
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// Reward Receiver Address is the address that would be receiving rewards for
	// the incentives for the lock. This is set to owner by default and can be
	// changed via separate msg.
	RewardReceiverAddress string `protobuf:"bytes,6,opt,name=reward_receiver_address,json=rewardReceiverAddress,proto3" json:"reward_receiver_address,omitempty" yaml:"reward_receiver_address"`
}

func (m *PeriodLock) Reset()         { *m = PeriodLock{} }
func (m *PeriodLock) String() string { return proto.CompactTextString(m) }
func (*PeriodLock) ProtoMessage()    {}
func (*PeriodLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9d7527a237b489, []int{0}
}
func (m *PeriodLock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodLock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeriodLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodLock.Merge(m, src)
}
func (m *PeriodLock) XXX_Size() int {
	return m.Size()
}
func (m *PeriodLock) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodLock.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodLock proto.InternalMessageInfo

func (m *PeriodLock) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PeriodLock) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PeriodLock) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PeriodLock) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *PeriodLock) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *PeriodLock) GetRewardReceiverAddress() string {
	if m != nil {
		return m.RewardReceiverAddress
	}
	return ""
}

// QueryCondition is a struct used for querying locks upon different conditions.
// Duration field and timestamp fields could be optional, depending on the
// LockQueryType.
type QueryCondition struct {
	// LockQueryType is a type of lock query, ByLockDuration | ByLockTime
	LockQueryType LockQueryType `protobuf:"varint,1,opt,name=lock_query_type,json=lockQueryType,proto3,enum=osmosis.lockup.LockQueryType" json:"lock_query_type,omitempty"`
	// Denom represents the token denomination we are looking to lock up
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	// Duration is used to query locks with longer duration than the specified
	// duration. Duration field must not be nil when the lock query type is
	// `ByLockDuration`.
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration" yaml:"duration"`
	// Timestamp is used by locks started before the specified duration.
	// Timestamp field must not be nil when the lock query type is `ByLockTime`.
	// Querying locks with timestamp is currently not implemented.
	Timestamp time.Time `protobuf:"bytes,4,opt,name=timestamp,proto3,stdtime" json:"timestamp" yaml:"timestamp"`
}

func (m *QueryCondition) Reset()         { *m = QueryCondition{} }
func (m *QueryCondition) String() string { return proto.CompactTextString(m) }
func (*QueryCondition) ProtoMessage()    {}
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9d7527a237b489, []int{1}
}
func (m *QueryCondition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCondition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCondition.Merge(m, src)
}
func (m *QueryCondition) XXX_Size() int {
	return m.Size()
}
func (m *QueryCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCondition.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCondition proto.InternalMessageInfo

func (m *QueryCondition) GetLockQueryType() LockQueryType {
	if m != nil {
		return m.LockQueryType
	}
	return ByDuration
}

func (m *QueryCondition) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *QueryCondition) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *QueryCondition) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

// SyntheticLock is creating virtual lockup where new denom is combination of
// original denom and synthetic suffix. At the time of synthetic lockup creation
// and deletion, accumulation store is also being updated and on querier side,
// they can query as freely as native lockup.
type SyntheticLock struct {
	// Underlying Lock ID is the underlying native lock's id for this synthetic
	// lockup. A synthetic lock MUST have an underlying lock.
	UnderlyingLockId uint64 `protobuf:"varint,1,opt,name=underlying_lock_id,json=underlyingLockId,proto3" json:"underlying_lock_id,omitempty"`
	// SynthDenom is the synthetic denom that is a combination of
	// gamm share + bonding status + validator address.
	SynthDenom string `protobuf:"bytes,2,opt,name=synth_denom,json=synthDenom,proto3" json:"synth_denom,omitempty"`
	// used for unbonding synthetic lockups, for active synthetic lockups, this
	// value is set to uninitialized value
	EndTime time.Time `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	// Duration is the duration for a synthetic lock to mature
	// at the point of unbonding has started.
	Duration time.Duration `protobuf:"bytes,4,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
}

func (m *SyntheticLock) Reset()         { *m = SyntheticLock{} }
func (m *SyntheticLock) String() string { return proto.CompactTextString(m) }
func (*SyntheticLock) ProtoMessage()    {}
func (*SyntheticLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9d7527a237b489, []int{2}
}
func (m *SyntheticLock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyntheticLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyntheticLock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SyntheticLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyntheticLock.Merge(m, src)
}
func (m *SyntheticLock) XXX_Size() int {
	return m.Size()
}
func (m *SyntheticLock) XXX_DiscardUnknown() {
	xxx_messageInfo_SyntheticLock.DiscardUnknown(m)
}

var xxx_messageInfo_SyntheticLock proto.InternalMessageInfo

func (m *SyntheticLock) GetUnderlyingLockId() uint64 {
	if m != nil {
		return m.UnderlyingLockId
	}
	return 0
}

func (m *SyntheticLock) GetSynthDenom() string {
	if m != nil {
		return m.SynthDenom
	}
	return ""
}

func (m *SyntheticLock) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *SyntheticLock) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterEnum("osmosis.lockup.LockQueryType", LockQueryType_name, LockQueryType_value)
	proto.RegisterType((*PeriodLock)(nil), "osmosis.lockup.PeriodLock")
	proto.RegisterType((*QueryCondition)(nil), "osmosis.lockup.QueryCondition")
	proto.RegisterType((*SyntheticLock)(nil), "osmosis.lockup.SyntheticLock")
}

func init() { proto.RegisterFile("osmosis/lockup/lock.proto", fileDescriptor_7e9d7527a237b489) }

var fileDescriptor_7e9d7527a237b489 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x1c, 0x8d, 0xf3, 0xa7, 0x7f, 0xae, 0x34, 0xb5, 0x4e, 0x45, 0xa4, 0x01, 0xec, 0xc8, 0x03, 0x8a,
	0x50, 0x6b, 0x93, 0x32, 0x20, 0xb1, 0xe1, 0x06, 0xa1, 0xa2, 0x0a, 0x81, 0xa9, 0x18, 0xba, 0x58,
	0x8e, 0xef, 0x48, 0x4f, 0x8d, 0x7d, 0xe6, 0xce, 0x6e, 0xf1, 0x37, 0x60, 0xec, 0x08, 0x12, 0x1b,
	0x1b, 0xdf, 0x82, 0xad, 0x63, 0x47, 0xa6, 0x14, 0xb5, 0x1b, 0x63, 0x3f, 0x01, 0xba, 0x3b, 0x3b,
	0x4d, 0x8b, 0x2a, 0x75, 0x80, 0xc9, 0xbe, 0x7b, 0xbf, 0xdf, 0xbb, 0x9f, 0xdf, 0x7b, 0x67, 0xb0,
	0x42, 0x79, 0x44, 0x39, 0xe1, 0xce, 0x88, 0x86, 0x7b, 0x59, 0x22, 0x1f, 0x76, 0xc2, 0x68, 0x4a,
	0x61, 0xb3, 0x80, 0x6c, 0x05, 0xb5, 0x97, 0x87, 0x74, 0x48, 0x25, 0xe4, 0x88, 0x37, 0x55, 0xd5,
	0x36, 0x86, 0x94, 0x0e, 0x47, 0xd8, 0x91, 0xab, 0x41, 0xf6, 0xde, 0x41, 0x19, 0x0b, 0x52, 0x42,
	0xe3, 0x02, 0x37, 0xaf, 0xe2, 0x29, 0x89, 0x30, 0x4f, 0x83, 0x28, 0x29, 0x09, 0x42, 0x79, 0x8e,
	0x33, 0x08, 0x38, 0x76, 0xf6, 0x7b, 0x03, 0x9c, 0x06, 0x3d, 0x27, 0xa4, 0xa4, 0x20, 0xb0, 0x7e,
	0xd4, 0x00, 0x78, 0x8d, 0x19, 0xa1, 0x68, 0x8b, 0x86, 0x7b, 0xb0, 0x09, 0xaa, 0x9b, 0xfd, 0x96,
	0xd6, 0xd1, 0xba, 0x75, 0xaf, 0xba, 0xd9, 0x87, 0x0f, 0x40, 0x83, 0x1e, 0xc4, 0x98, 0xb5, 0xaa,
	0x1d, 0xad, 0x3b, 0xef, 0xea, 0xe7, 0x63, 0xf3, 0x56, 0x1e, 0x44, 0xa3, 0xa7, 0x96, 0xdc, 0xb6,
	0x3c, 0x05, 0xc3, 0x5d, 0x30, 0x57, 0x4e, 0xd6, 0xaa, 0x75, 0xb4, 0xee, 0xc2, 0xfa, 0x8a, 0xad,
	0x46, 0xb3, 0xcb, 0xd1, 0xec, 0x7e, 0x51, 0xe0, 0xf6, 0x8e, 0xc6, 0x66, 0xe5, 0xf7, 0xd8, 0x84,
	0x65, 0xcb, 0x2a, 0x8d, 0x48, 0x8a, 0xa3, 0x24, 0xcd, 0xcf, 0xc7, 0xe6, 0x92, 0xe2, 0x2f, 0x31,
	0xeb, 0xf3, 0x89, 0xa9, 0x79, 0x13, 0x76, 0xe8, 0x81, 0x39, 0x1c, 0x23, 0x5f, 0x7c, 0x67, 0xab,
	0x2e, 0x4f, 0x6a, 0xff, 0x75, 0xd2, 0x76, 0x29, 0x82, 0x7b, 0x57, 0x1c, 0x75, 0x41, 0x5a, 0x76,
	0x5a, 0x87, 0x82, 0x74, 0x16, 0xc7, 0x48, 0x94, 0xc2, 0x00, 0x34, 0x84, 0x24, 0xbc, 0xd5, 0xe8,
	0xd4, 0xe4, 0xe8, 0x4a, 0x34, 0x5b, 0x88, 0x66, 0x17, 0xa2, 0xd9, 0x1b, 0x94, 0xc4, 0xee, 0x23,
	0xc1, 0xf7, 0xfd, 0xc4, 0xec, 0x0e, 0x49, 0xba, 0x9b, 0x0d, 0xec, 0x90, 0x46, 0x4e, 0xa1, 0xb0,
	0x7a, 0xac, 0x71, 0xb4, 0xe7, 0xa4, 0x79, 0x82, 0xb9, 0x6c, 0xe0, 0x9e, 0x62, 0x86, 0x3b, 0xe0,
	0x0e, 0xc3, 0x07, 0x01, 0x43, 0x3e, 0xc3, 0x21, 0x26, 0xfb, 0x98, 0xf9, 0x01, 0x42, 0x0c, 0x73,
	0xde, 0x9a, 0x91, 0xd2, 0x5a, 0xe7, 0x63, 0xd3, 0x50, 0x53, 0x5e, 0x53, 0x68, 0x79, 0xb7, 0x15,
	0xe2, 0x15, 0xc0, 0xb3, 0x62, 0xff, 0x4b, 0x15, 0x34, 0xdf, 0x64, 0x98, 0xe5, 0x1b, 0x34, 0x46,
	0x44, 0xaa, 0xf4, 0x1c, 0x2c, 0x89, 0x5c, 0xf9, 0x1f, 0xc4, 0xb6, 0x2f, 0xe6, 0x91, 0xa6, 0x36,
	0xd7, 0xef, 0xdb, 0x97, 0x73, 0x67, 0x0b, 0xdb, 0x65, 0xf3, 0x76, 0x9e, 0x60, 0x6f, 0x71, 0x34,
	0xbd, 0x84, 0xcb, 0xa0, 0x81, 0x70, 0x4c, 0x23, 0x65, 0xbf, 0xa7, 0x16, 0xc2, 0x82, 0x9b, 0x9b,
	0x7d, 0xc5, 0x81, 0xeb, 0x6c, 0x7d, 0x07, 0xe6, 0x27, 0xd1, 0xbd, 0x81, 0xaf, 0xf7, 0x0a, 0x56,
	0x5d, 0xb1, 0x4e, 0x5a, 0x95, 0xb1, 0x17, 0x54, 0xd6, 0xd7, 0x2a, 0x58, 0x7c, 0x9b, 0xc7, 0xe9,
	0x2e, 0x4e, 0x49, 0x28, 0x23, 0xbe, 0x0a, 0x60, 0x16, 0x23, 0xcc, 0x46, 0x39, 0x89, 0x87, 0xbe,
	0x54, 0x89, 0xa0, 0x22, 0xf2, 0xfa, 0x05, 0x22, 0x6a, 0x37, 0x11, 0x34, 0xc1, 0x02, 0x17, 0xed,
	0xfe, 0xb4, 0x0e, 0x40, 0x6e, 0xf5, 0x4b, 0x31, 0x26, 0x79, 0xac, 0xfd, 0xa3, 0x3c, 0x4e, 0xdf,
	0xa6, 0xfa, 0xff, 0xbc, 0x4d, 0x0f, 0x5f, 0x82, 0xc5, 0x4b, 0x01, 0x80, 0x4d, 0x00, 0xdc, 0xbc,
	0xe4, 0xd6, 0x2b, 0x10, 0x80, 0x19, 0x37, 0x17, 0x43, 0xe9, 0x9a, 0x78, 0x7f, 0x45, 0x45, 0xb9,
	0x5e, 0x85, 0x0b, 0x60, 0xd6, 0xcd, 0x5f, 0x30, 0x9a, 0x25, 0x7a, 0xad, 0x5d, 0xff, 0xf4, 0xcd,
	0xa8, 0xb8, 0x5b, 0x47, 0xa7, 0x86, 0x76, 0x7c, 0x6a, 0x68, 0xbf, 0x4e, 0x0d, 0xed, 0xf0, 0xcc,
	0xa8, 0x1c, 0x9f, 0x19, 0x95, 0x9f, 0x67, 0x46, 0x65, 0x67, 0x7d, 0xea, 0xb6, 0x14, 0xf1, 0x5b,
	0x1b, 0x05, 0x03, 0x5e, 0x2e, 0x9c, 0xfd, 0xde, 0x13, 0xe7, 0x63, 0xf9, 0x93, 0x94, 0xb7, 0x67,
	0x30, 0x23, 0xbf, 0xf4, 0xf1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x87, 0x80, 0xb1, 0x43,
	0x05, 0x00, 0x00,
}

func (m *PeriodLock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodLock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeriodLock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RewardReceiverAddress) > 0 {
		i -= len(m.RewardReceiverAddress)
		copy(dAtA[i:], m.RewardReceiverAddress)
		i = encodeVarintLock(dAtA, i, uint64(len(m.RewardReceiverAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLock(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLock(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCondition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCondition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCondition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintLock(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintLock(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x1a
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.LockQueryType != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.LockQueryType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SyntheticLock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyntheticLock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SyntheticLock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n5, err5 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintLock(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x22
	n6, err6 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintLock(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0x1a
	if len(m.SynthDenom) > 0 {
		i -= len(m.SynthDenom)
		copy(dAtA[i:], m.SynthDenom)
		i = encodeVarintLock(dAtA, i, uint64(len(m.SynthDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.UnderlyingLockId != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.UnderlyingLockId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLock(dAtA []byte, offset int, v uint64) int {
	offset -= sovLock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PeriodLock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLock(uint64(m.ID))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLock(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovLock(uint64(l))
		}
	}
	l = len(m.RewardReceiverAddress)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	return n
}

func (m *QueryCondition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LockQueryType != 0 {
		n += 1 + sovLock(uint64(m.LockQueryType))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovLock(uint64(l))
	return n
}

func (m *SyntheticLock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnderlyingLockId != 0 {
		n += 1 + sovLock(uint64(m.UnderlyingLockId))
	}
	l = len(m.SynthDenom)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	return n
}

func sovLock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLock(x uint64) (n int) {
	return sovLock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PeriodLock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: PeriodLock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodLock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types1.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardReceiverAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardReceiverAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func (m *QueryCondition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: QueryCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockQueryType", wireType)
			}
			m.LockQueryType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockQueryType |= LockQueryType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func (m *SyntheticLock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: SyntheticLock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyntheticLock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnderlyingLockId", wireType)
			}
			m.UnderlyingLockId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnderlyingLockId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SynthDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SynthDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func skipLock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLock
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
					return 0, ErrIntOverflowLock
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
					return 0, ErrIntOverflowLock
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
				return 0, ErrInvalidLengthLock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLock = fmt.Errorf("proto: unexpected end of group")
)
