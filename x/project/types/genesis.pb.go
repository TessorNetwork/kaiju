// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/project/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the project module's genesis state.
type GenesisState struct {
	ProjectDocs      []ProjectDoc         `protobuf:"bytes,1,rep,name=project_docs,json=projectDocs,proto3" json:"project_docs" yaml:"project_docs"`
	AccountMaps      []GenesisAccountMap  `protobuf:"bytes,2,rep,name=account_maps,json=accountMaps,proto3" json:"account_maps" yaml:"account_maps"`
	WithdrawalsInfos []WithdrawalInfoDocs `protobuf:"bytes,3,rep,name=withdrawals_infos,json=withdrawalsInfos,proto3" json:"withdrawal_infos" yaml:"withdrawal_infos"`
	Claims           []Claims             `protobuf:"bytes,4,rep,name=claims,proto3" json:"claims" yaml:"claims"`
	Params           Params               `protobuf:"bytes,5,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b70d6215711b0fe, []int{0}
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

func (m *GenesisState) GetProjectDocs() []ProjectDoc {
	if m != nil {
		return m.ProjectDocs
	}
	return nil
}

func (m *GenesisState) GetAccountMaps() []GenesisAccountMap {
	if m != nil {
		return m.AccountMaps
	}
	return nil
}

func (m *GenesisState) GetWithdrawalsInfos() []WithdrawalInfoDocs {
	if m != nil {
		return m.WithdrawalsInfos
	}
	return nil
}

func (m *GenesisState) GetClaims() []Claims {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kaiju.project.v1.GenesisState")
}

func init() { proto.RegisterFile("kaiju/project/v1/genesis.proto", fileDescriptor_1b70d6215711b0fe) }

var fileDescriptor_1b70d6215711b0fe = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x31, 0x4f, 0xea, 0x50,
	0x18, 0x86, 0xdb, 0xcb, 0xbd, 0x0c, 0x85, 0x9b, 0x70, 0x7b, 0x35, 0x36, 0x68, 0x5a, 0x52, 0x1d,
	0x9c, 0xda, 0x20, 0x9b, 0x9b, 0xd5, 0x04, 0x1d, 0x4c, 0x48, 0x1d, 0x4c, 0x8c, 0x09, 0x39, 0x1c,
	0x4a, 0x29, 0xd0, 0x9e, 0xa6, 0xdf, 0x01, 0xe4, 0x0f, 0x38, 0xfb, 0xb3, 0x18, 0x19, 0x9d, 0x1a,
	0x03, 0x9b, 0x71, 0xe2, 0x17, 0x98, 0x9e, 0x53, 0x2b, 0x69, 0xdd, 0x9a, 0xbc, 0xef, 0xfb, 0x3c,
	0xa7, 0xc9, 0x27, 0xa9, 0x63, 0xe4, 0x8d, 0xa6, 0x66, 0x18, 0x91, 0x91, 0x83, 0xa9, 0x39, 0x6b,
	0x9a, 0xae, 0x13, 0x38, 0xe0, 0x81, 0x11, 0x46, 0x84, 0x12, 0xb9, 0xc6, 0x72, 0x23, 0xcd, 0x8d,
	0x59, 0xb3, 0xbe, 0xe7, 0x12, 0x97, 0xb0, 0xd0, 0x4c, 0xbe, 0x78, 0xaf, 0x5e, 0xe4, 0x7c, 0x4d,
	0x58, 0xae, 0x7f, 0x94, 0xa4, 0x6a, 0x9b, 0x93, 0xef, 0x28, 0xa2, 0x8e, 0xfc, 0x28, 0x55, 0xd3,
	0x46, 0xb7, 0x4f, 0x30, 0x28, 0x62, 0xa3, 0x74, 0x5a, 0x39, 0x3b, 0x32, 0xf2, 0x3e, 0xa3, 0xc3,
	0x3f, 0xaf, 0x08, 0xb6, 0x0e, 0x97, 0xb1, 0x26, 0x6c, 0x63, 0xed, 0xff, 0x02, 0xf9, 0x93, 0x73,
	0x7d, 0x77, 0xaf, 0xdb, 0x95, 0x30, 0x2b, 0x82, 0x8c, 0xa5, 0x2a, 0xc2, 0x98, 0x4c, 0x03, 0xda,
	0xf5, 0x51, 0x08, 0xca, 0x2f, 0x46, 0x3f, 0x2e, 0xd2, 0xd3, 0x37, 0x5d, 0xf0, 0xf2, 0x2d, 0x0a,
	0xf3, 0x92, 0x5d, 0x8c, 0x6e, 0x57, 0x50, 0x56, 0x04, 0xf9, 0x59, 0x94, 0xfe, 0xcd, 0x3d, 0x3a,
	0xec, 0x47, 0x68, 0x8e, 0x26, 0xd0, 0xf5, 0x82, 0x01, 0x01, 0xa5, 0xc4, 0x54, 0x27, 0x45, 0xd5,
	0x7d, 0x56, 0xbd, 0x09, 0x06, 0x24, 0x79, 0xa6, 0xd5, 0x4a, 0x5c, 0xef, 0xb1, 0x56, 0xfb, 0xc6,
	0x70, 0xca, 0x36, 0xd6, 0x0e, 0xb8, 0x3f, 0x9f, 0xe8, 0xf6, 0x4e, 0x19, 0x12, 0x12, 0xc8, 0x6d,
	0xa9, 0x8c, 0x27, 0xc8, 0xf3, 0x41, 0xf9, 0xcd, 0xe4, 0x4a, 0x51, 0x7e, 0xc9, 0x72, 0x6b, 0x3f,
	0xfd, 0xb9, 0xbf, 0x1c, 0xce, 0x57, 0xba, 0x9d, 0xce, 0x13, 0x50, 0x88, 0x22, 0xe4, 0x83, 0xf2,
	0xa7, 0x21, 0xfe, 0x0c, 0xea, 0xb0, 0x3c, 0x0f, 0xe2, 0x2b, 0xdd, 0x4e, 0xe7, 0xd6, 0xf5, 0x72,
	0xad, 0x8a, 0xab, 0xb5, 0x2a, 0xbe, 0xad, 0x55, 0xf1, 0x65, 0xa3, 0x0a, 0xab, 0x8d, 0x2a, 0xbc,
	0x6e, 0x54, 0xe1, 0xc1, 0x70, 0x3d, 0x3a, 0x9c, 0xf6, 0x0c, 0x4c, 0x7c, 0x93, 0x3a, 0x00, 0x24,
	0x0a, 0x1c, 0x3a, 0x27, 0xd1, 0xd8, 0xe4, 0x17, 0xf4, 0x94, 0xdd, 0x10, 0x5d, 0x84, 0x0e, 0xf4,
	0xca, 0xec, 0x7e, 0x5a, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xfc, 0x1c, 0x76, 0xa9, 0x02,
	0x00, 0x00,
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Claims) > 0 {
		for iNdEx := len(m.Claims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.WithdrawalsInfos) > 0 {
		for iNdEx := len(m.WithdrawalsInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawalsInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.AccountMaps) > 0 {
		for iNdEx := len(m.AccountMaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountMaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ProjectDocs) > 0 {
		for iNdEx := len(m.ProjectDocs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectDocs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.ProjectDocs) > 0 {
		for _, e := range m.ProjectDocs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AccountMaps) > 0 {
		for _, e := range m.AccountMaps {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WithdrawalsInfos) > 0 {
		for _, e := range m.WithdrawalsInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectDocs", wireType)
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
			m.ProjectDocs = append(m.ProjectDocs, ProjectDoc{})
			if err := m.ProjectDocs[len(m.ProjectDocs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountMaps", wireType)
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
			m.AccountMaps = append(m.AccountMaps, GenesisAccountMap{})
			if err := m.AccountMaps[len(m.AccountMaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalsInfos", wireType)
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
			m.WithdrawalsInfos = append(m.WithdrawalsInfos, WithdrawalInfoDocs{})
			if err := m.WithdrawalsInfos[len(m.WithdrawalsInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
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
			m.Claims = append(m.Claims, Claims{})
			if err := m.Claims[len(m.Claims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
