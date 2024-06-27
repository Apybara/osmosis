// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/smartaccount/v1beta1/genesis.proto

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

// AuthenticatorData represents a genesis exported account with Authenticators.
// The address is used as the key, and the account authenticators are stored in
// the authenticators field.
type AuthenticatorData struct {
	// address is an account address, one address can have many authenticators
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// authenticators are the account's authenticators, these can be multiple
	// types including SignatureVerification, AllOfs, CosmWasmAuthenticators, etc
	Authenticators []AccountAuthenticator `protobuf:"bytes,2,rep,name=authenticators,proto3" json:"authenticators"`
}

func (m *AuthenticatorData) Reset()         { *m = AuthenticatorData{} }
func (m *AuthenticatorData) String() string { return proto.CompactTextString(m) }
func (*AuthenticatorData) ProtoMessage()    {}
func (*AuthenticatorData) Descriptor() ([]byte, []int) {
	return fileDescriptor_678d63c22c684b43, []int{0}
}
func (m *AuthenticatorData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticatorData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticatorData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticatorData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticatorData.Merge(m, src)
}
func (m *AuthenticatorData) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticatorData) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticatorData.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticatorData proto.InternalMessageInfo

func (m *AuthenticatorData) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AuthenticatorData) GetAuthenticators() []AccountAuthenticator {
	if m != nil {
		return m.Authenticators
	}
	return nil
}

// GenesisState defines the authenticator module's genesis state.
type GenesisState struct {
	// params define the parameters for the authenticator module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// next_authenticator_id is the next available authenticator ID.
	NextAuthenticatorId uint64 `protobuf:"varint,2,opt,name=next_authenticator_id,json=nextAuthenticatorId,proto3" json:"next_authenticator_id,omitempty"`
	// authenticator_data contains the data for multiple accounts, each with their
	// authenticators.
	AuthenticatorData []AuthenticatorData `protobuf:"bytes,3,rep,name=authenticator_data,json=authenticatorData,proto3" json:"authenticator_data"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_678d63c22c684b43, []int{1}
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

func (m *GenesisState) GetNextAuthenticatorId() uint64 {
	if m != nil {
		return m.NextAuthenticatorId
	}
	return 0
}

func (m *GenesisState) GetAuthenticatorData() []AuthenticatorData {
	if m != nil {
		return m.AuthenticatorData
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticatorData)(nil), "osmosis.smartaccount.v1beta1.AuthenticatorData")
	proto.RegisterType((*GenesisState)(nil), "osmosis.smartaccount.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/smartaccount/v1beta1/genesis.proto", fileDescriptor_678d63c22c684b43)
}

var fileDescriptor_678d63c22c684b43 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x77, 0x54, 0x8c, 0xc6, 0x08, 0xdc, 0x0a, 0x16, 0x89, 0x4d, 0xa4, 0x83, 0x05, 0xee,
	0xe0, 0x46, 0x87, 0x8e, 0x4a, 0x10, 0xdd, 0x42, 0x6f, 0x5d, 0xec, 0xed, 0xce, 0xb0, 0x2e, 0xb8,
	0x3b, 0xb2, 0xf3, 0x14, 0xfb, 0x14, 0xf5, 0xb1, 0x3c, 0x7a, 0xec, 0x14, 0xa1, 0xd7, 0x3e, 0x44,
	0xb8, 0x3b, 0x82, 0x6b, 0x60, 0xdd, 0xe6, 0x31, 0xbf, 0xf7, 0xfe, 0xef, 0xff, 0xfe, 0xf4, 0x5a,
	0xaa, 0x48, 0xaa, 0x50, 0x31, 0x15, 0x41, 0x82, 0xe0, 0xfb, 0x72, 0x12, 0x23, 0x9b, 0xb6, 0x3d,
	0x81, 0xd0, 0x66, 0x81, 0x88, 0x85, 0x0a, 0x95, 0x33, 0x4e, 0x24, 0x4a, 0xf3, 0x5c, 0xb3, 0xce,
	0x36, 0xeb, 0x68, 0xb6, 0x76, 0x1a, 0xc8, 0x40, 0xa6, 0x20, 0x5b, 0xbf, 0xb2, 0x9e, 0xda, 0xd5,
	0xde, 0xf9, 0x63, 0x48, 0x20, 0x52, 0xff, 0x42, 0x23, 0xc9, 0xc5, 0x48, 0xa3, 0x8d, 0x37, 0x42,
	0xab, 0x9d, 0x09, 0x0e, 0x45, 0x8c, 0xa1, 0x0f, 0x28, 0x93, 0x7b, 0x40, 0x30, 0x2d, 0x7a, 0x00,
	0x9c, 0x27, 0x42, 0x29, 0x8b, 0xd4, 0x49, 0xf3, 0xb0, 0xb7, 0x29, 0xcd, 0x17, 0x7a, 0x0c, 0xdb,
	0xb8, 0xb2, 0x0a, 0xf5, 0x62, 0xb3, 0xe2, 0xba, 0xce, 0x3e, 0x4b, 0x4e, 0x27, 0xab, 0x73, 0x4a,
	0xdd, 0xd2, 0xfc, 0xf3, 0xc2, 0xe8, 0xed, 0xcc, 0x6b, 0x7c, 0x13, 0x7a, 0xf4, 0x90, 0x5d, 0xab,
	0x8f, 0x80, 0xc2, 0xec, 0xd2, 0x72, 0xe6, 0x2e, 0xdd, 0xa5, 0xe2, 0x5e, 0xee, 0x97, 0x7a, 0x4a,
	0x59, 0x3d, 0x5c, 0x77, 0x9a, 0x2e, 0x3d, 0x8b, 0xc5, 0x0c, 0x07, 0x39, 0xad, 0x41, 0xc8, 0xad,
	0x42, 0x9d, 0x34, 0x4b, 0xbd, 0x93, 0xf5, 0x67, 0x6e, 0xb9, 0x47, 0x6e, 0x72, 0x6a, 0xe6, 0x71,
	0x0e, 0x08, 0x56, 0x31, 0xb5, 0xcb, 0xfe, 0xb0, 0xbb, 0x7b, 0x51, 0xbd, 0x4e, 0x15, 0x7e, 0x7d,
	0xf4, 0xe7, 0x4b, 0x9b, 0x2c, 0x96, 0x36, 0xf9, 0x5a, 0xda, 0xe4, 0x7d, 0x65, 0x1b, 0x8b, 0x95,
	0x6d, 0x7c, 0xac, 0x6c, 0xe3, 0xf9, 0x2e, 0x08, 0x71, 0x38, 0xf1, 0x1c, 0x5f, 0x46, 0x4c, 0xab,
	0xb5, 0x46, 0xe0, 0xa9, 0x4d, 0xc1, 0xa6, 0xee, 0x2d, 0x9b, 0x65, 0x19, 0xb7, 0x36, 0x21, 0xe3,
	0xeb, 0x58, 0x28, 0xaf, 0x9c, 0x86, 0x7b, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x8f, 0xf2,
	0xe4, 0x94, 0x02, 0x00, 0x00,
}

func (m *AuthenticatorData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticatorData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthenticatorData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authenticators) > 0 {
		for iNdEx := len(m.Authenticators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Authenticators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.AuthenticatorData) > 0 {
		for iNdEx := len(m.AuthenticatorData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AuthenticatorData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.NextAuthenticatorId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextAuthenticatorId))
		i--
		dAtA[i] = 0x10
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
func (m *AuthenticatorData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Authenticators) > 0 {
		for _, e := range m.Authenticators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.NextAuthenticatorId != 0 {
		n += 1 + sovGenesis(uint64(m.NextAuthenticatorId))
	}
	if len(m.AuthenticatorData) > 0 {
		for _, e := range m.AuthenticatorData {
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
func (m *AuthenticatorData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AuthenticatorData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticatorData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticators", wireType)
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
			m.Authenticators = append(m.Authenticators, AccountAuthenticator{})
			if err := m.Authenticators[len(m.Authenticators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextAuthenticatorId", wireType)
			}
			m.NextAuthenticatorId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextAuthenticatorId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticatorData", wireType)
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
			m.AuthenticatorData = append(m.AuthenticatorData, AuthenticatorData{})
			if err := m.AuthenticatorData[len(m.AuthenticatorData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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