// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jinxprotocol/bridge/genesis.proto

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

// GenesisState defines the bridge module's genesis state.
type GenesisState struct {
	// The parameters of the module.
	EventParams   EventParams   `protobuf:"bytes,1,opt,name=event_params,json=eventParams,proto3" json:"event_params"`
	ProposeParams ProposeParams `protobuf:"bytes,2,opt,name=propose_params,json=proposeParams,proto3" json:"propose_params"`
	SafetyParams  SafetyParams  `protobuf:"bytes,3,opt,name=safety_params,json=safetyParams,proto3" json:"safety_params"`
	// Acknowledged event info that stores:
	// - the next event ID to be added to consensus.
	// - Ethereum block height of the most recently acknowledged bridge event.
	AcknowledgedEventInfo BridgeEventInfo `protobuf:"bytes,4,opt,name=acknowledged_event_info,json=acknowledgedEventInfo,proto3" json:"acknowledged_event_info"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8fc7ee5d5949615, []int{0}
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

func (m *GenesisState) GetEventParams() EventParams {
	if m != nil {
		return m.EventParams
	}
	return EventParams{}
}

func (m *GenesisState) GetProposeParams() ProposeParams {
	if m != nil {
		return m.ProposeParams
	}
	return ProposeParams{}
}

func (m *GenesisState) GetSafetyParams() SafetyParams {
	if m != nil {
		return m.SafetyParams
	}
	return SafetyParams{}
}

func (m *GenesisState) GetAcknowledgedEventInfo() BridgeEventInfo {
	if m != nil {
		return m.AcknowledgedEventInfo
	}
	return BridgeEventInfo{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "jinxprotocol.bridge.GenesisState")
}

func init() { proto.RegisterFile("jinxprotocol/bridge/genesis.proto", fileDescriptor_c8fc7ee5d5949615) }

var fileDescriptor_c8fc7ee5d5949615 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4f, 0x02, 0x31,
	0x14, 0xc7, 0xef, 0x90, 0x38, 0x14, 0x70, 0x38, 0x35, 0x12, 0x86, 0x0a, 0xc4, 0xc1, 0xc4, 0x78,
	0x97, 0xa8, 0x83, 0x33, 0x89, 0x31, 0x24, 0x26, 0x12, 0xd8, 0x5c, 0x48, 0xef, 0x78, 0x94, 0x2a,
	0xb4, 0xcd, 0xb5, 0x22, 0x7c, 0x0b, 0x3f, 0x16, 0x23, 0xa3, 0x83, 0x31, 0x06, 0xbe, 0x88, 0xa1,
	0xe5, 0xe4, 0x48, 0x3a, 0xbd, 0xe6, 0xff, 0x7e, 0xf9, 0xf5, 0xb5, 0x0f, 0x35, 0x5e, 0x19, 0x9f,
	0xc9, 0x54, 0x68, 0x91, 0x88, 0x71, 0x14, 0xa7, 0x6c, 0x40, 0x21, 0xa2, 0xc0, 0x41, 0x31, 0x15,
	0x9a, 0x3c, 0x38, 0xce, 0x23, 0xa1, 0x45, 0x6a, 0x27, 0x54, 0x50, 0x61, 0xc2, 0x68, 0x73, 0xb2,
	0x68, 0xed, 0xca, 0x65, 0xb3, 0xa5, 0x0f, 0x53, 0xe0, 0xba, 0xcf, 0xf8, 0x30, 0x83, 0xeb, 0x2e,
	0x58, 0x92, 0x94, 0x4c, 0xb6, 0x37, 0x37, 0xbf, 0x0b, 0xa8, 0xfc, 0x68, 0x67, 0xe9, 0x69, 0xa2,
	0x21, 0x68, 0xa3, 0xb2, 0xd5, 0x58, 0xac, 0xea, 0xd7, 0xfd, 0xcb, 0xd2, 0x4d, 0x3d, 0x74, 0x4c,
	0x18, 0x3e, 0x6c, 0xc0, 0x8e, 0xe1, 0x5a, 0xc5, 0xc5, 0xcf, 0xb9, 0xd7, 0x2d, 0xc1, 0x2e, 0x0a,
	0x9e, 0xd1, 0x91, 0x4c, 0x85, 0x14, 0x0a, 0x32, 0x59, 0xc1, 0xc8, 0x9a, 0x4e, 0x59, 0xc7, 0xa2,
	0x7b, 0xba, 0x8a, 0xcc, 0x87, 0xc1, 0x13, 0xaa, 0x28, 0x32, 0x04, 0x3d, 0xcf, 0x7c, 0x07, 0xc6,
	0xd7, 0x70, 0xfa, 0x7a, 0x86, 0xdc, 0xd3, 0x95, 0x55, 0x2e, 0x0b, 0x62, 0x74, 0x46, 0x92, 0x37,
	0x2e, 0x3e, 0xc6, 0x30, 0xa0, 0x30, 0xc8, 0xfd, 0x5e, 0xb5, 0x68, 0xbc, 0x17, 0x4e, 0x6f, 0xcb,
	0x14, 0xf3, 0xf4, 0x36, 0x1f, 0x8a, 0xad, 0xfa, 0x34, 0xaf, 0xda, 0x35, 0xbb, 0x8b, 0x15, 0xf6,
	0x97, 0x2b, 0xec, 0xff, 0xae, 0xb0, 0xff, 0xb9, 0xc6, 0xde, 0x72, 0x8d, 0xbd, 0xaf, 0x35, 0xf6,
	0x5e, 0xee, 0x29, 0xd3, 0xa3, 0xf7, 0x38, 0x4c, 0xc4, 0x24, 0xda, 0xdb, 0xd2, 0xf4, 0xee, 0x3a,
	0x19, 0x11, 0xc6, 0xa3, 0xff, 0x64, 0x96, 0x6d, 0x4e, 0xcf, 0x25, 0xa8, 0xf8, 0xd0, 0x34, 0x6e,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x88, 0x9e, 0xde, 0x41, 0x58, 0x02, 0x00, 0x00,
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
		size, err := m.AcknowledgedEventInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.SafetyParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ProposeParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.EventParams.MarshalToSizedBuffer(dAtA[:i])
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EventParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ProposeParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SafetyParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.AcknowledgedEventInfo.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field EventParams", wireType)
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
			if err := m.EventParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposeParams", wireType)
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
			if err := m.ProposeParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SafetyParams", wireType)
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
			if err := m.SafetyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcknowledgedEventInfo", wireType)
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
			if err := m.AcknowledgedEventInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
