// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/denom.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Denom struct {
	Base     string `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Display  string `protobuf:"bytes,2,opt,name=display,proto3" json:"display,omitempty"`
	InitTime int64  `protobuf:"varint,3,opt,name=init_time,json=initTime,proto3" json:"init_time,omitempty"`
	// Positive Interest rate on denom debt pool
	DebtInterestRate github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,5,opt,name=debt_interest_rate,json=debtInterestRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"debt_interest_rate"`
	// Total shares of the denom debt pool
	DebtShares github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,6,opt,name=debt_shares,json=debtShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"debt_shares"`
	// Negative Interest rate on Denoms Bonded
	BondInterestRate github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,7,opt,name=bond_interest_rate,json=bondInterestRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"bond_interest_rate"`
	// Total shares of the denom bond pool
	BondShares github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,8,opt,name=bond_shares,json=bondShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"bond_shares"`
	// Unbonding blocks
	BondBlocks github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,9,opt,name=bond_blocks,json=bondBlocks,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"bond_blocks"`
}

func (m *Denom) Reset()         { *m = Denom{} }
func (m *Denom) String() string { return proto.CompactTextString(m) }
func (*Denom) ProtoMessage()    {}
func (*Denom) Descriptor() ([]byte, []int) {
	return fileDescriptor_26550448cb73c9b3, []int{0}
}
func (m *Denom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Denom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Denom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Denom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Denom.Merge(m, src)
}
func (m *Denom) XXX_Size() int {
	return m.Size()
}
func (m *Denom) XXX_DiscardUnknown() {
	xxx_messageInfo_Denom.DiscardUnknown(m)
}

var xxx_messageInfo_Denom proto.InternalMessageInfo

func (m *Denom) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Denom) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Denom) GetInitTime() int64 {
	if m != nil {
		return m.InitTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Denom)(nil), "reserve.Denom")
}

func init() { proto.RegisterFile("reserve/denom.proto", fileDescriptor_26550448cb73c9b3) }

var fileDescriptor_26550448cb73c9b3 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd2, 0x3f, 0x4e, 0xf3, 0x30,
	0x18, 0x06, 0xf0, 0xf8, 0xeb, 0x7f, 0x7f, 0x0b, 0x32, 0x48, 0x58, 0x20, 0xb9, 0x15, 0x0b, 0x5d,
	0xa8, 0x07, 0xc4, 0x05, 0x2a, 0x16, 0x36, 0x14, 0x60, 0x41, 0x42, 0x55, 0xdc, 0xbc, 0x4a, 0xad,
	0x36, 0x71, 0x64, 0x1b, 0x44, 0x2f, 0xc0, 0xcc, 0x31, 0x38, 0x4a, 0xc7, 0x8e, 0x88, 0xa1, 0x42,
	0xc9, 0x45, 0x90, 0x4d, 0xb2, 0x74, 0x0c, 0x93, 0x5f, 0x3f, 0xb6, 0x7e, 0x7a, 0x86, 0x17, 0x1f,
	0x6a, 0x30, 0xa0, 0x5f, 0x80, 0xc7, 0x90, 0xa9, 0x74, 0x92, 0x6b, 0x65, 0x15, 0xe9, 0x55, 0xe1,
	0xc9, 0x51, 0xa2, 0x12, 0xe5, 0x33, 0xee, 0xa6, 0xdf, 0xe7, 0xb3, 0xb7, 0x36, 0xee, 0x5c, 0xbb,
	0xef, 0x84, 0xe0, 0xb6, 0x88, 0x0c, 0x50, 0x34, 0x42, 0xe3, 0x41, 0xe8, 0x67, 0x42, 0x71, 0x2f,
	0x96, 0x26, 0x5f, 0x45, 0x6b, 0xfa, 0xcf, 0xc7, 0xf5, 0x95, 0x9c, 0xe2, 0x81, 0xcc, 0xa4, 0x9d,
	0x59, 0x99, 0x02, 0x6d, 0x8d, 0xd0, 0xb8, 0x15, 0xf6, 0x5d, 0x70, 0x2f, 0x53, 0x20, 0x4f, 0x98,
	0xc4, 0x20, 0xec, 0x4c, 0x66, 0x16, 0x34, 0x18, 0x3b, 0xd3, 0x91, 0x05, 0xda, 0x71, 0xc2, 0x94,
	0x6f, 0x76, 0xc3, 0xe0, 0x6b, 0x37, 0x3c, 0x4f, 0xa4, 0x5d, 0x3c, 0x8b, 0xc9, 0x5c, 0xa5, 0x7c,
	0xae, 0x4c, 0xaa, 0x4c, 0x75, 0x5c, 0x98, 0x78, 0xc9, 0xed, 0x3a, 0x07, 0x33, 0x79, 0x90, 0x99,
	0x0d, 0x0f, 0x1c, 0x75, 0x53, 0x49, 0x61, 0x64, 0x81, 0xdc, 0xe2, 0xff, 0x9e, 0x37, 0x8b, 0x48,
	0x83, 0xa1, 0xdd, 0x66, 0x2e, 0x76, 0xc6, 0x9d, 0x27, 0x5c, 0x61, 0xa1, 0xb2, 0x78, 0xaf, 0x70,
	0xaf, 0x61, 0x61, 0x47, 0xed, 0x17, 0xf6, 0x7c, 0x55, 0xb8, 0xdf, 0xb0, 0xb0, 0x33, 0xaa, 0xc2,
	0xb5, 0x28, 0x56, 0x6a, 0xbe, 0x34, 0x74, 0xf0, 0x07, 0x71, 0xea, 0x89, 0xe9, 0xd5, 0x47, 0xc1,
	0xd0, 0xa6, 0x60, 0x68, 0x5b, 0x30, 0xf4, 0x5d, 0x30, 0xf4, 0x5e, 0xb2, 0x60, 0x5b, 0xb2, 0xe0,
	0xb3, 0x64, 0xc1, 0xe3, 0x71, 0xbd, 0x5a, 0xaf, 0xbc, 0x9e, 0xbc, 0x23, 0xba, 0x7e, 0x8d, 0x2e,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x4f, 0x3a, 0xdc, 0x7c, 0x02, 0x00, 0x00,
}

func (this *Denom) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Denom)
	if !ok {
		that2, ok := that.(Denom)
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
	if this.Base != that1.Base {
		return false
	}
	if this.Display != that1.Display {
		return false
	}
	if this.InitTime != that1.InitTime {
		return false
	}
	if !this.DebtInterestRate.Equal(that1.DebtInterestRate) {
		return false
	}
	if !this.DebtShares.Equal(that1.DebtShares) {
		return false
	}
	if !this.BondInterestRate.Equal(that1.BondInterestRate) {
		return false
	}
	if !this.BondShares.Equal(that1.BondShares) {
		return false
	}
	if !this.BondBlocks.Equal(that1.BondBlocks) {
		return false
	}
	return true
}
func (m *Denom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Denom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Denom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BondBlocks.Size()
		i -= size
		if _, err := m.BondBlocks.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDenom(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.BondShares.Size()
		i -= size
		if _, err := m.BondShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDenom(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.BondInterestRate.Size()
		i -= size
		if _, err := m.BondInterestRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDenom(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.DebtShares.Size()
		i -= size
		if _, err := m.DebtShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDenom(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.DebtInterestRate.Size()
		i -= size
		if _, err := m.DebtInterestRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDenom(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.InitTime != 0 {
		i = encodeVarintDenom(dAtA, i, uint64(m.InitTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Display) > 0 {
		i -= len(m.Display)
		copy(dAtA[i:], m.Display)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Display)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Base) > 0 {
		i -= len(m.Base)
		copy(dAtA[i:], m.Base)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Base)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDenom(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Denom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Base)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	l = len(m.Display)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	if m.InitTime != 0 {
		n += 1 + sovDenom(uint64(m.InitTime))
	}
	l = m.DebtInterestRate.Size()
	n += 1 + l + sovDenom(uint64(l))
	l = m.DebtShares.Size()
	n += 1 + l + sovDenom(uint64(l))
	l = m.BondInterestRate.Size()
	n += 1 + l + sovDenom(uint64(l))
	l = m.BondShares.Size()
	n += 1 + l + sovDenom(uint64(l))
	l = m.BondBlocks.Size()
	n += 1 + l + sovDenom(uint64(l))
	return n
}

func sovDenom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenom(x uint64) (n int) {
	return sovDenom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Denom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenom
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
			return fmt.Errorf("proto: Denom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Denom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Base = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Display", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Display = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitTime", wireType)
			}
			m.InitTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtInterestRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtInterestRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondInterestRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondInterestRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondBlocks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondBlocks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenom
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
func skipDenom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenom
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
					return 0, ErrIntOverflowDenom
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
					return 0, ErrIntOverflowDenom
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
				return 0, ErrInvalidLengthDenom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenom = fmt.Errorf("proto: unexpected end of group")
)
