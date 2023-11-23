// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examplevess/examplevess/l1hash.proto

package types

import (
	fmt "fmt"
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

type L1Hash struct {
	Block uint64 `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	Hash  string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *L1Hash) Reset()         { *m = L1Hash{} }
func (m *L1Hash) String() string { return proto.CompactTextString(m) }
func (*L1Hash) ProtoMessage()    {}
func (*L1Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_da69b0976218de0e, []int{0}
}
func (m *L1Hash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *L1Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_L1Hash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *L1Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L1Hash.Merge(m, src)
}
func (m *L1Hash) XXX_Size() int {
	return m.Size()
}
func (m *L1Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_L1Hash.DiscardUnknown(m)
}

var xxx_messageInfo_L1Hash proto.InternalMessageInfo

func (m *L1Hash) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *L1Hash) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*L1Hash)(nil), "examplevess.examplevess.L1Hash")
}

func init() {
	proto.RegisterFile("examplevess/examplevess/l1hash.proto", fileDescriptor_da69b0976218de0e)
}

var fileDescriptor_da69b0976218de0e = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0x4b, 0x2d, 0x2e, 0xd6, 0x47, 0x66, 0xe7, 0x18, 0x66, 0x24, 0x16, 0x67,
	0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x23, 0xc9, 0xe8, 0x21, 0xb1, 0x95, 0x8c, 0xb8,
	0xd8, 0x7c, 0x0c, 0x3d, 0x12, 0x8b, 0x33, 0x84, 0x44, 0xb8, 0x58, 0x93, 0x72, 0xf2, 0x93, 0xb3,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x16, 0x90, 0x31, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x93, 0xe5, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x44, 0xc9, 0x23, 0x3b, 0xa0, 0x02, 0xc5, 0x39, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0x60, 0xe7, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x84, 0xc5, 0xa3, 0xb6,
	0x00, 0x00, 0x00,
}

func (m *L1Hash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *L1Hash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *L1Hash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintL1Hash(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Block != 0 {
		i = encodeVarintL1Hash(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintL1Hash(dAtA []byte, offset int, v uint64) int {
	offset -= sovL1Hash(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *L1Hash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != 0 {
		n += 1 + sovL1Hash(uint64(m.Block))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovL1Hash(uint64(l))
	}
	return n
}

func sovL1Hash(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozL1Hash(x uint64) (n int) {
	return sovL1Hash(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *L1Hash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowL1Hash
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
			return fmt.Errorf("proto: L1Hash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: L1Hash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowL1Hash
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowL1Hash
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
				return ErrInvalidLengthL1Hash
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthL1Hash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipL1Hash(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthL1Hash
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
func skipL1Hash(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowL1Hash
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
					return 0, ErrIntOverflowL1Hash
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
					return 0, ErrIntOverflowL1Hash
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
				return 0, ErrInvalidLengthL1Hash
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupL1Hash
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthL1Hash
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthL1Hash        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowL1Hash          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupL1Hash = fmt.Errorf("proto: unexpected end of group")
)
