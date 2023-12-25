// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jinxprotocol/bridge/bridge_event.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

// BridgeEvent is a recognized event from the Ethereum blockchain.
type BridgeEvent struct {
	// The unique id of the Ethereum event log.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The tokens bridged.
	Coin types.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
	// The account address or module address to bridge to.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// The Ethereum block height of the event.
	EthBlockHeight uint64 `protobuf:"varint,4,opt,name=eth_block_height,json=ethBlockHeight,proto3" json:"eth_block_height,omitempty"`
}

func (m *BridgeEvent) Reset()         { *m = BridgeEvent{} }
func (m *BridgeEvent) String() string { return proto.CompactTextString(m) }
func (*BridgeEvent) ProtoMessage()    {}
func (*BridgeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ff43cb4f3a7b88, []int{0}
}
func (m *BridgeEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BridgeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BridgeEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BridgeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeEvent.Merge(m, src)
}
func (m *BridgeEvent) XXX_Size() int {
	return m.Size()
}
func (m *BridgeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeEvent proto.InternalMessageInfo

func (m *BridgeEvent) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BridgeEvent) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *BridgeEvent) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BridgeEvent) GetEthBlockHeight() uint64 {
	if m != nil {
		return m.EthBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*BridgeEvent)(nil), "jinxprotocol.bridge.BridgeEvent")
}

func init() {
	proto.RegisterFile("jinxprotocol/bridge/bridge_event.proto", fileDescriptor_06ff43cb4f3a7b88)
}

var fileDescriptor_06ff43cb4f3a7b88 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4e, 0xf3, 0x30,
	0x18, 0x86, 0xe3, 0xfe, 0xd1, 0x8f, 0x48, 0x45, 0x85, 0x42, 0x87, 0xb4, 0x83, 0x89, 0x18, 0x50,
	0x96, 0xc6, 0x6a, 0xcb, 0xc0, 0x4a, 0x10, 0x12, 0x73, 0xd8, 0x58, 0xa2, 0xd8, 0xb1, 0x6c, 0x43,
	0x6b, 0x57, 0xb1, 0xa9, 0xca, 0x2d, 0x38, 0x0a, 0x03, 0x87, 0xe8, 0x58, 0x31, 0x31, 0x21, 0xd4,
	0x5e, 0x04, 0xc5, 0x4e, 0x11, 0x4c, 0xf6, 0xf7, 0x3e, 0x8f, 0x5f, 0x4b, 0x5f, 0x70, 0xfe, 0x20,
	0xe4, 0x6a, 0x51, 0x2b, 0xa3, 0x88, 0x9a, 0x21, 0x5c, 0x8b, 0x8a, 0xd1, 0xf6, 0x28, 0xe8, 0x92,
	0x4a, 0x93, 0x5a, 0x18, 0x9e, 0xfc, 0xf6, 0x52, 0x27, 0x0c, 0xfb, 0x4c, 0x31, 0x65, 0x43, 0xd4,
	0xdc, 0x9c, 0x3a, 0x1c, 0x10, 0xa5, 0xe7, 0x4a, 0x17, 0x0e, 0xb8, 0xa1, 0x45, 0xd0, 0x4d, 0x08,
	0x97, 0x9a, 0xa2, 0xe5, 0x18, 0x53, 0x53, 0x8e, 0x11, 0x51, 0x42, 0x3a, 0x7e, 0xf6, 0x0a, 0x82,
	0x6e, 0x66, 0xbb, 0x6f, 0x9a, 0xbf, 0xc3, 0x5e, 0xd0, 0x11, 0x55, 0x04, 0x62, 0x90, 0x1c, 0xe5,
	0x1d, 0x51, 0x85, 0xd3, 0xc0, 0x6f, 0xec, 0xa8, 0x13, 0x83, 0xa4, 0x3b, 0x19, 0xa4, 0x6d, 0x79,
	0x53, 0x97, 0xb6, 0x75, 0xe9, 0xb5, 0x12, 0x32, 0xf3, 0xd7, 0x9f, 0xa7, 0x5e, 0x6e, 0xe5, 0x70,
	0x12, 0x1c, 0x94, 0x55, 0x55, 0x53, 0xad, 0xa3, 0x7f, 0x31, 0x48, 0x0e, 0xb3, 0xe8, 0xfd, 0x6d,
	0xd4, 0x6f, 0x9f, 0x5e, 0x39, 0x72, 0x67, 0x6a, 0x21, 0x59, 0xbe, 0x17, 0xc3, 0x24, 0x38, 0xa6,
	0x86, 0x17, 0x78, 0xa6, 0xc8, 0x63, 0xc1, 0xa9, 0x60, 0xdc, 0x44, 0x7e, 0x0c, 0x12, 0x3f, 0xef,
	0x51, 0xc3, 0xb3, 0x26, 0xbe, 0xb5, 0x69, 0x96, 0xaf, 0xb7, 0x10, 0x6c, 0xb6, 0x10, 0x7c, 0x6d,
	0x21, 0x78, 0xd9, 0x41, 0x6f, 0xb3, 0x83, 0xde, 0xc7, 0x0e, 0x7a, 0xf7, 0x97, 0x4c, 0x18, 0xfe,
	0x84, 0x53, 0xa2, 0xe6, 0xe8, 0xcf, 0x96, 0x97, 0x17, 0x23, 0xc2, 0x4b, 0x21, 0xd1, 0x4f, 0xb2,
	0xda, 0x6f, 0xde, 0x3c, 0x2f, 0xa8, 0xc6, 0xff, 0x2d, 0x98, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x25, 0x29, 0x28, 0x68, 0x9d, 0x01, 0x00, 0x00,
}

func (m *BridgeEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BridgeEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BridgeEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EthBlockHeight != 0 {
		i = encodeVarintBridgeEvent(dAtA, i, uint64(m.EthBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBridgeEvent(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBridgeEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintBridgeEvent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBridgeEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovBridgeEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BridgeEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovBridgeEvent(uint64(m.Id))
	}
	l = m.Coin.Size()
	n += 1 + l + sovBridgeEvent(uint64(l))
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBridgeEvent(uint64(l))
	}
	if m.EthBlockHeight != 0 {
		n += 1 + sovBridgeEvent(uint64(m.EthBlockHeight))
	}
	return n
}

func sovBridgeEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBridgeEvent(x uint64) (n int) {
	return sovBridgeEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BridgeEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBridgeEvent
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
			return fmt.Errorf("proto: BridgeEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BridgeEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeEvent
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
				return ErrInvalidLengthBridgeEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBridgeEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeEvent
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
				return ErrInvalidLengthBridgeEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBridgeEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthBlockHeight", wireType)
			}
			m.EthBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBridgeEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBridgeEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBridgeEvent
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
func skipBridgeEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBridgeEvent
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
					return 0, ErrIntOverflowBridgeEvent
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
					return 0, ErrIntOverflowBridgeEvent
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
				return 0, ErrInvalidLengthBridgeEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBridgeEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBridgeEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBridgeEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBridgeEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBridgeEvent = fmt.Errorf("proto: unexpected end of group")
)
