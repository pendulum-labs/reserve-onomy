// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/bond.proto

package types

import (
	fmt "fmt"
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

type Bond struct {
	Uid        string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner      string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Denom      string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	BondShares string `protobuf:"bytes,4,opt,name=bondShares,proto3" json:"bondShares,omitempty"`
}

func (m *Bond) Reset()         { *m = Bond{} }
func (m *Bond) String() string { return proto.CompactTextString(m) }
func (*Bond) ProtoMessage()    {}
func (*Bond) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b6795bb86e37479, []int{0}
}
func (m *Bond) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bond.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bond.Merge(m, src)
}
func (m *Bond) XXX_Size() int {
	return m.Size()
}
func (m *Bond) XXX_DiscardUnknown() {
	xxx_messageInfo_Bond.DiscardUnknown(m)
}

var xxx_messageInfo_Bond proto.InternalMessageInfo

func (m *Bond) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Bond) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Bond) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Bond) GetBondShares() string {
	if m != nil {
		return m.BondShares
	}
	return ""
}

func init() {
	proto.RegisterType((*Bond)(nil), "reserve.reserve.Bond")
}

func init() { proto.RegisterFile("reserve/bond.proto", fileDescriptor_5b6795bb86e37479) }

var fileDescriptor_5b6795bb86e37479 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4a, 0x2d, 0x4e,
	0x2d, 0x2a, 0x4b, 0xd5, 0x4f, 0xca, 0xcf, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x87, 0x8a, 0xe9, 0x41, 0x69, 0xa5, 0x14, 0x2e, 0x16, 0xa7, 0xfc, 0xbc, 0x14, 0x21, 0x01, 0x2e,
	0xe6, 0xd2, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x84, 0x8b,
	0x35, 0xbf, 0x3c, 0x2f, 0xb5, 0x48, 0x82, 0x09, 0x2c, 0x06, 0xe1, 0x80, 0x44, 0x53, 0x52, 0xf3,
	0xf2, 0x73, 0x25, 0x98, 0x21, 0xa2, 0x60, 0x8e, 0x90, 0x1c, 0x17, 0x17, 0xc8, 0x92, 0xe0, 0x8c,
	0xc4, 0xa2, 0xd4, 0x62, 0x09, 0x16, 0xb0, 0x14, 0x92, 0x88, 0x93, 0xe1, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x89, 0xc3, 0x1c, 0x59, 0xa1, 0x0f, 0x63, 0x95, 0x54, 0x16,
	0xa4, 0x16, 0x27, 0xb1, 0x81, 0x1d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x69, 0x11,
	0x78, 0xc6, 0x00, 0x00, 0x00,
}

func (m *Bond) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bond) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bond) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BondShares) > 0 {
		i -= len(m.BondShares)
		copy(dAtA[i:], m.BondShares)
		i = encodeVarintBond(dAtA, i, uint64(len(m.BondShares)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintBond(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintBond(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintBond(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBond(dAtA []byte, offset int, v uint64) int {
	offset -= sovBond(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bond) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovBond(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovBond(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovBond(uint64(l))
	}
	l = len(m.BondShares)
	if l > 0 {
		n += 1 + l + sovBond(uint64(l))
	}
	return n
}

func sovBond(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBond(x uint64) (n int) {
	return sovBond(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bond) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBond
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
			return fmt.Errorf("proto: Bond: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bond: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBond
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
				return ErrInvalidLengthBond
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBond
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
				return ErrInvalidLengthBond
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBond
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
				return ErrInvalidLengthBond
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBond
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
				return ErrInvalidLengthBond
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BondShares = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBond(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBond
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
func skipBond(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBond
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
					return 0, ErrIntOverflowBond
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
					return 0, ErrIntOverflowBond
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
				return 0, ErrInvalidLengthBond
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBond
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBond
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBond        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBond          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBond = fmt.Errorf("proto: unexpected end of group")
)
