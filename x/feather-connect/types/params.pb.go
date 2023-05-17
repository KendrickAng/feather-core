// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feather_connect/params.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/terra-money/alliance/x/alliance/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	HaltIfNoChannel    bool                            `protobuf:"varint,1,opt,name=halt_if_no_channel,json=haltIfNoChannel,proto3" json:"halt_if_no_channel,omitempty"`
	BaseDenom          string                          `protobuf:"bytes,2,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	BaseChainId        string                          `protobuf:"bytes,3,opt,name=base_chain_id,json=baseChainId,proto3" json:"base_chain_id,omitempty"`
	AllianceBondHeight int64                           `protobuf:"varint,4,opt,name=alliance_bond_height,json=allianceBondHeight,proto3" json:"alliance_bond_height,omitempty"`
	Alliance           types.MsgCreateAllianceProposal `protobuf:"bytes,5,opt,name=alliance,proto3" json:"alliance"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db56e51105fa5f2, []int{0}
}

func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}

func (m *Params) XXX_Size() int {
	return m.Size()
}

func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetHaltIfNoChannel() bool {
	if m != nil {
		return m.HaltIfNoChannel
	}
	return false
}

func (m *Params) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *Params) GetBaseChainId() string {
	if m != nil {
		return m.BaseChainId
	}
	return ""
}

func (m *Params) GetAllianceBondHeight() int64 {
	if m != nil {
		return m.AllianceBondHeight
	}
	return 0
}

func (m *Params) GetAlliance() types.MsgCreateAllianceProposal {
	if m != nil {
		return m.Alliance
	}
	return types.MsgCreateAllianceProposal{}
}

func init() {
	proto.RegisterType((*Params)(nil), "feather_connect.Params")
}

func init() { proto.RegisterFile("feather_connect/params.proto", fileDescriptor_6db56e51105fa5f2) }

var fileDescriptor_6db56e51105fa5f2 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd1, 0x41, 0x4b, 0xe3, 0x40,
	0x14, 0x07, 0xf0, 0xcc, 0xb6, 0x5b, 0xda, 0x29, 0x4b, 0x61, 0xe8, 0x21, 0x94, 0xdd, 0x6c, 0xa8,
	0x97, 0x80, 0x34, 0x11, 0x3d, 0x7a, 0xb2, 0x55, 0xb0, 0x07, 0xa5, 0x04, 0xbc, 0x78, 0x19, 0x26,
	0xc9, 0x6b, 0x12, 0x48, 0xe6, 0x85, 0xc9, 0x28, 0xf6, 0x5b, 0xf8, 0xb1, 0x7a, 0xec, 0xd1, 0x93,
	0x48, 0xfb, 0x09, 0xfc, 0x06, 0x92, 0xa4, 0x2d, 0xe2, 0xed, 0xcd, 0xef, 0xff, 0x66, 0x98, 0x79,
	0x43, 0xff, 0x2e, 0x41, 0xe8, 0x04, 0x14, 0x0f, 0x51, 0x4a, 0x08, 0xb5, 0x57, 0x08, 0x25, 0xf2,
	0xd2, 0x2d, 0x14, 0x6a, 0x64, 0x83, 0x1f, 0xe9, 0x88, 0x89, 0x2c, 0x4b, 0x85, 0x0c, 0xc1, 0x8b,
	0xf1, 0xb9, 0x69, 0x1a, 0x0d, 0x63, 0x8c, 0xb1, 0x2e, 0xbd, 0xaa, 0x6a, 0x74, 0xfc, 0x49, 0x68,
	0x67, 0x51, 0x9f, 0xc5, 0x4e, 0x29, 0x4b, 0x44, 0xa6, 0x79, 0xba, 0xe4, 0x12, 0x79, 0x98, 0x08,
	0x29, 0x21, 0x33, 0x89, 0x4d, 0x9c, 0xae, 0x3f, 0xa8, 0x92, 0xf9, 0xf2, 0x1e, 0x67, 0x0d, 0xb3,
	0x7f, 0x94, 0x06, 0xa2, 0x04, 0x1e, 0x81, 0xc4, 0xdc, 0xfc, 0x65, 0x13, 0xa7, 0xe7, 0xf7, 0x2a,
	0xb9, 0xae, 0x80, 0x8d, 0xe9, 0x9f, 0x3a, 0x0e, 0x13, 0x91, 0x4a, 0x9e, 0x46, 0x66, 0xab, 0xee,
	0xe8, 0x57, 0x38, 0xab, 0x6c, 0x1e, 0xb1, 0x33, 0x3a, 0x3c, 0x5c, 0x93, 0x07, 0x28, 0x23, 0x9e,
	0x40, 0x1a, 0x27, 0xda, 0x6c, 0xdb, 0xc4, 0x69, 0xf9, 0xc7, 0x27, 0x4c, 0x51, 0x46, 0xb7, 0x75,
	0xc2, 0x6e, 0x68, 0xf7, 0xa0, 0xe6, 0x6f, 0x9b, 0x38, 0xfd, 0xf3, 0x13, 0xf7, 0x00, 0xee, 0x5d,
	0x19, 0xcf, 0x14, 0x08, 0x0d, 0x57, 0x7b, 0x59, 0x28, 0x2c, 0xb0, 0x14, 0xd9, 0xb4, 0xbd, 0x7e,
	0xff, 0x6f, 0xf8, 0xc7, 0xad, 0xd3, 0x87, 0xf5, 0xd6, 0x22, 0x9b, 0xad, 0x45, 0x3e, 0xb6, 0x16,
	0x79, 0xdd, 0x59, 0xc6, 0x66, 0x67, 0x19, 0x6f, 0x3b, 0xcb, 0x78, 0xbc, 0x8c, 0x53, 0x9d, 0x3c,
	0x05, 0x6e, 0x88, 0xb9, 0xa7, 0x41, 0x29, 0x31, 0xc9, 0x51, 0xc2, 0xca, 0xdb, 0xcf, 0x77, 0x12,
	0xa2, 0x02, 0xef, 0xe5, 0xdb, 0xb2, 0xf9, 0x0c, 0xbd, 0x2a, 0xa0, 0x0c, 0x3a, 0xf5, 0x44, 0x2f,
	0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x20, 0xf8, 0x6b, 0xac, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Alliance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.AllianceBondHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AllianceBondHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BaseChainId) > 0 {
		i -= len(m.BaseChainId)
		copy(dAtA[i:], m.BaseChainId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BaseChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.HaltIfNoChannel {
		i--
		if m.HaltIfNoChannel {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HaltIfNoChannel {
		n += 2
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.BaseChainId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.AllianceBondHeight != 0 {
		n += 1 + sovParams(uint64(m.AllianceBondHeight))
	}
	l = m.Alliance.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HaltIfNoChannel", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.HaltIfNoChannel = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllianceBondHeight", wireType)
			}
			m.AllianceBondHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllianceBondHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alliance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Alliance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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

func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
