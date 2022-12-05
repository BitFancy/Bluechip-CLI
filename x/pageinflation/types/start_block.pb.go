// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pageinflation/start_block.proto

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

type StartBlock struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StartBlock) Reset()         { *m = StartBlock{} }
func (m *StartBlock) String() string { return proto.CompactTextString(m) }
func (*StartBlock) ProtoMessage()    {}
func (*StartBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8099d24dde06966, []int{0}
}
func (m *StartBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StartBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StartBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StartBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBlock.Merge(m, src)
}
func (m *StartBlock) XXX_Size() int {
	return m.Size()
}
func (m *StartBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBlock.DiscardUnknown(m)
}

var xxx_messageInfo_StartBlock proto.InternalMessageInfo

func (m *StartBlock) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StartBlock)(nil), "smartdev0328.bluechip.pageinflation.StartBlock")
}

func init() { proto.RegisterFile("pageinflation/start_block.proto", fileDescriptor_d8099d24dde06966) }

var fileDescriptor_d8099d24dde06966 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x48, 0x4c, 0x4f,
	0xcd, 0xcc, 0x4b, 0xcb, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x2e, 0x49, 0x2c, 0x2a, 0x89,
	0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2e, 0xce, 0x4d,
	0x2c, 0x2a, 0x49, 0x49, 0x2d, 0x33, 0x30, 0x36, 0xb2, 0xd0, 0x4b, 0xca, 0x29, 0x4d, 0x4d, 0xce,
	0xc8, 0x2c, 0xd0, 0x43, 0xd1, 0xa6, 0xa4, 0xc4, 0xc5, 0x15, 0x0c, 0xd2, 0xe9, 0x04, 0xd2, 0x28,
	0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe1, 0x38, 0x05, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x59, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xb2, 0x6d, 0xfa, 0x30, 0xdb, 0xf4,
	0x2b, 0xf4, 0x51, 0x9d, 0x59, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x76, 0xa1, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x22, 0xa7, 0xcc, 0xbb, 0xc4, 0x00, 0x00, 0x00,
}

func (m *StartBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StartBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintStartBlock(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStartBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovStartBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StartBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovStartBlock(uint64(l))
	}
	return n
}

func sovStartBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStartBlock(x uint64) (n int) {
	return sovStartBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StartBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStartBlock
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
			return fmt.Errorf("proto: StartBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStartBlock
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
				return ErrInvalidLengthStartBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStartBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStartBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStartBlock
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
func skipStartBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStartBlock
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
					return 0, ErrIntOverflowStartBlock
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
					return 0, ErrIntOverflowStartBlock
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
				return 0, ErrInvalidLengthStartBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStartBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStartBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStartBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStartBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStartBlock = fmt.Errorf("proto: unexpected end of group")
)
