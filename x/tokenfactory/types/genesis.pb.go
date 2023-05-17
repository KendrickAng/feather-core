// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenfactory/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the tokenfactory module's genesis state.
type GenesisState struct {
	// params defines the parameters of the module.
	Params        Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FactoryDenoms []GenesisDenom `protobuf:"bytes,2,rep,name=factory_denoms,json=factoryDenoms,proto3" json:"factory_denoms" yaml:"factory_denoms"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_415d5acd9b7bd461, []int{0}
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

func (m *GenesisState) GetFactoryDenoms() []GenesisDenom {
	if m != nil {
		return m.FactoryDenoms
	}
	return nil
}

// GenesisDenom defines a tokenfactory denom that is defined within genesis
// state. The structure contains DenomAuthorityMetadata which defines the
// denom's admin.
type GenesisDenom struct {
	Denom             string                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	AuthorityMetadata DenomAuthorityMetadata `protobuf:"bytes,2,opt,name=authority_metadata,json=authorityMetadata,proto3" json:"authority_metadata" yaml:"authority_metadata"`
}

func (m *GenesisDenom) Reset()         { *m = GenesisDenom{} }
func (m *GenesisDenom) String() string { return proto.CompactTextString(m) }
func (*GenesisDenom) ProtoMessage()    {}
func (*GenesisDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_415d5acd9b7bd461, []int{1}
}

func (m *GenesisDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GenesisDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GenesisDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDenom.Merge(m, src)
}

func (m *GenesisDenom) XXX_Size() int {
	return m.Size()
}

func (m *GenesisDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDenom.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDenom proto.InternalMessageInfo

func (m *GenesisDenom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *GenesisDenom) GetAuthorityMetadata() DenomAuthorityMetadata {
	if m != nil {
		return m.AuthorityMetadata
	}
	return DenomAuthorityMetadata{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tokenfactory.GenesisState")
	proto.RegisterType((*GenesisDenom)(nil), "tokenfactory.GenesisDenom")
}

func init() { proto.RegisterFile("tokenfactory/genesis.proto", fileDescriptor_415d5acd9b7bd461) }

var fileDescriptor_415d5acd9b7bd461 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0xed, 0xf0, 0xf1, 0x91, 0x58, 0xd0, 0x68, 0x83, 0x09, 0x34, 0xb1, 0xc5, 0x86, 0x18, 0x36,
	0xb4, 0x09, 0x6e, 0x0c, 0x3b, 0x1b, 0x12, 0x57, 0x26, 0xa6, 0xee, 0xdc, 0xe0, 0x00, 0x97, 0x42,
	0xb4, 0x1d, 0x32, 0xbd, 0x18, 0xfb, 0x16, 0x3e, 0x82, 0x0b, 0x1f, 0xc1, 0x87, 0x60, 0xc9, 0xd2,
	0x15, 0x31, 0xb0, 0x71, 0xcd, 0x13, 0x18, 0x66, 0x26, 0x4a, 0x65, 0xd7, 0xde, 0xf3, 0x73, 0xcf,
	0x9c, 0xab, 0x9b, 0xc8, 0x1e, 0x20, 0x1e, 0xd2, 0x3e, 0x32, 0x9e, 0x7a, 0x21, 0xc4, 0x90, 0x8c,
	0x13, 0x77, 0xc2, 0x19, 0x32, 0xa3, 0xb4, 0x8d, 0x99, 0xe5, 0x90, 0x85, 0x4c, 0x00, 0xde, 0xe6,
	0x4b, 0x72, 0xcc, 0x7a, 0x46, 0x4f, 0xa7, 0x38, 0x62, 0x7c, 0x8c, 0xe9, 0x35, 0x20, 0x1d, 0x50,
	0xa4, 0x8a, 0x55, 0xcd, 0xb0, 0x26, 0x94, 0xd3, 0x48, 0x2d, 0x71, 0xde, 0x88, 0x5e, 0xba, 0x92,
	0x6b, 0x6f, 0x91, 0x22, 0x18, 0x2d, 0xbd, 0x20, 0x09, 0x15, 0x52, 0x23, 0x8d, 0x62, 0xab, 0xec,
	0x6e, 0x8b, 0xdd, 0x1b, 0x81, 0xf9, 0xf9, 0xd9, 0xc2, 0xd6, 0x02, 0xc5, 0x34, 0xee, 0xf5, 0x03,
	0x85, 0x77, 0x07, 0x10, 0xb3, 0x28, 0xa9, 0xe4, 0x6a, 0xff, 0x1a, 0xc5, 0x96, 0x99, 0xd5, 0xaa,
	0x3d, 0x9d, 0x0d, 0xc5, 0x3f, 0xd9, 0x38, 0xac, 0x17, 0xf6, 0x71, 0x4a, 0xa3, 0xc7, 0xb6, 0x93,
	0xd5, 0x3b, 0xc1, 0xbe, 0x1a, 0x74, 0xe4, 0xff, 0xfb, 0x6f, 0x4c, 0x31, 0x31, 0xce, 0xf4, 0xff,
	0x82, 0x2a, 0x52, 0xee, 0xf9, 0x87, 0xeb, 0x85, 0x5d, 0x92, 0x4e, 0x62, 0xec, 0x04, 0x12, 0x36,
	0x9e, 0x74, 0xe3, 0xa7, 0x95, 0x6e, 0xa4, 0x6a, 0xa9, 0xe4, 0xc4, 0xd3, 0xea, 0xd9, 0x78, 0xc2,
	0xf8, 0xf2, 0x6f, 0x85, 0xfe, 0xa9, 0x0a, 0x5a, 0x95, 0xf6, 0xbb, 0x6e, 0x4e, 0x70, 0xb4, 0x53,
	0x7c, 0x3b, 0xff, 0xf5, 0x6a, 0x13, 0x3f, 0x98, 0x2d, 0x2d, 0x32, 0x5f, 0x5a, 0xe4, 0x73, 0x69,
	0x91, 0x97, 0x95, 0xa5, 0xcd, 0x57, 0x96, 0xf6, 0xb1, 0xb2, 0xb4, 0xbb, 0x8b, 0x70, 0x8c, 0xa3,
	0x69, 0xcf, 0xed, 0xb3, 0xc8, 0x43, 0xe0, 0x9c, 0x36, 0x23, 0x16, 0x43, 0xea, 0x0d, 0x81, 0xe2,
	0x08, 0x78, 0xb3, 0xcf, 0x38, 0x78, 0xcf, 0x5e, 0xe6, 0x70, 0x98, 0x4e, 0x20, 0xe9, 0x15, 0xc4,
	0xe1, 0xce, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x33, 0x76, 0x9a, 0xbb, 0x3b, 0x02, 0x00, 0x00,
}

func (this *GenesisDenom) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisDenom)
	if !ok {
		that2, ok := that.(GenesisDenom)
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
	if this.Denom != that1.Denom {
		return false
	}
	if !this.AuthorityMetadata.Equal(&that1.AuthorityMetadata) {
		return false
	}
	return true
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
	if len(m.FactoryDenoms) > 0 {
		for iNdEx := len(m.FactoryDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FactoryDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AuthorityMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
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

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FactoryDenoms) > 0 {
		for _, e := range m.FactoryDenoms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.AuthorityMetadata.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FactoryDenoms", wireType)
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
			m.FactoryDenoms = append(m.FactoryDenoms, GenesisDenom{})
			if err := m.FactoryDenoms[len(m.FactoryDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func (m *GenesisDenom) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorityMetadata", wireType)
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
			if err := m.AuthorityMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
