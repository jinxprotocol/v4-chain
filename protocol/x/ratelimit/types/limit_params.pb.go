// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jinxprotocol/ratelimit/limit_params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_jinxprotocol_v4_chain_protocol_dtypes "github.com/jinxprotocol/v4-chain/protocol/dtypes"
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

// LimitParams defines rate limit params on a denom.
type LimitParams struct {
	// denom is the denomination of the token being rate limited.
	// e.g. ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// limiters is a list of rate-limiters on this denom. All limiters
	// must be satified for a withdrawal to proceed.
	Limiters []Limiter `protobuf:"bytes,2,rep,name=limiters,proto3" json:"limiters"`
}

func (m *LimitParams) Reset()         { *m = LimitParams{} }
func (m *LimitParams) String() string { return proto.CompactTextString(m) }
func (*LimitParams) ProtoMessage()    {}
func (*LimitParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ac4e080b94759db, []int{0}
}
func (m *LimitParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitParams.Merge(m, src)
}
func (m *LimitParams) XXX_Size() int {
	return m.Size()
}
func (m *LimitParams) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitParams.DiscardUnknown(m)
}

var xxx_messageInfo_LimitParams proto.InternalMessageInfo

func (m *LimitParams) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *LimitParams) GetLimiters() []Limiter {
	if m != nil {
		return m.Limiters
	}
	return nil
}

// Limiter defines one rate-limiter on a specfic denom.
type Limiter struct {
	// period_sec is the rolling time period for which the limit applies
	// e.g. 3600 (an hour)
	PeriodSec uint32 `protobuf:"varint,2,opt,name=period_sec,json=periodSec,proto3" json:"period_sec,omitempty"`
	// baseline_minimum is the minimum maximum withdrawal coin amount within the
	// time period.
	// e.g. 100_000_000_000 uusdc for 100k USDC; 5e22 jinx for 50k DV4TNT
	BaselineMinimum github_com_jinxprotocol_v4_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,3,opt,name=baseline_minimum,json=baselineMinimum,proto3,customtype=github.com/jinxprotocol/v4-chain/protocol/dtypes.SerializableInt" json:"baseline_minimum"`
	// baseline_tvl_ppm is the maximum ratio of TVL withdrawable in
	// the time period, in part-per-million.
	// e.g. 100_000 (10%)
	BaselineTvlPpm uint32 `protobuf:"varint,4,opt,name=baseline_tvl_ppm,json=baselineTvlPpm,proto3" json:"baseline_tvl_ppm,omitempty"`
}

func (m *Limiter) Reset()         { *m = Limiter{} }
func (m *Limiter) String() string { return proto.CompactTextString(m) }
func (*Limiter) ProtoMessage()    {}
func (*Limiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ac4e080b94759db, []int{1}
}
func (m *Limiter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Limiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Limiter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Limiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Limiter.Merge(m, src)
}
func (m *Limiter) XXX_Size() int {
	return m.Size()
}
func (m *Limiter) XXX_DiscardUnknown() {
	xxx_messageInfo_Limiter.DiscardUnknown(m)
}

var xxx_messageInfo_Limiter proto.InternalMessageInfo

func (m *Limiter) GetPeriodSec() uint32 {
	if m != nil {
		return m.PeriodSec
	}
	return 0
}

func (m *Limiter) GetBaselineTvlPpm() uint32 {
	if m != nil {
		return m.BaselineTvlPpm
	}
	return 0
}

func init() {
	proto.RegisterType((*LimitParams)(nil), "jinxprotocol.ratelimit.LimitParams")
	proto.RegisterType((*Limiter)(nil), "jinxprotocol.ratelimit.Limiter")
}

func init() {
	proto.RegisterFile("jinxprotocol/ratelimit/limit_params.proto", fileDescriptor_9ac4e080b94759db)
}

var fileDescriptor_9ac4e080b94759db = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0x6d, 0xb6, 0xf9, 0x67, 0x99, 0xff, 0x28, 0x43, 0x8a, 0x60, 0x57, 0x76, 0xaa, 0x07, 0x5b,
	0x50, 0x6f, 0x5e, 0x74, 0x27, 0x85, 0x09, 0xa3, 0xd3, 0x8b, 0x97, 0x92, 0x76, 0x71, 0x8b, 0x24,
	0x4d, 0x48, 0xb2, 0x31, 0xfd, 0x14, 0x7e, 0x2b, 0x77, 0xdc, 0x51, 0x3c, 0x0c, 0xd9, 0xbe, 0x88,
	0x2c, 0x75, 0xa3, 0x82, 0x07, 0x2f, 0x21, 0x79, 0xef, 0xe5, 0xbd, 0xc7, 0xef, 0x07, 0x4f, 0x9e,
	0x49, 0x36, 0x16, 0x92, 0x6b, 0x9e, 0x72, 0x1a, 0x4a, 0xa4, 0x31, 0x25, 0x8c, 0xe8, 0xd0, 0x9c,
	0xb1, 0x40, 0x12, 0x31, 0x15, 0x18, 0xde, 0x3e, 0x2c, 0x4a, 0x83, 0xb5, 0xf4, 0xa8, 0xde, 0xe7,
	0x7d, 0x6e, 0xf0, 0x70, 0x79, 0xcb, 0xd5, 0xcd, 0x27, 0x58, 0x6b, 0x2f, 0xe9, 0x8e, 0xb1, 0xb0,
	0xeb, 0x70, 0xa3, 0x87, 0x33, 0xce, 0x1c, 0xe0, 0x01, 0xbf, 0x1a, 0xe5, 0x0f, 0xfb, 0x1a, 0x6e,
	0x1b, 0x0f, 0x2c, 0x95, 0x53, 0xf2, 0xca, 0x7e, 0xed, 0xac, 0x11, 0xfc, 0x9d, 0x12, 0xb4, 0x73,
	0x5d, 0xab, 0x32, 0x99, 0x35, 0xac, 0x68, 0xfd, 0xad, 0xf9, 0x0e, 0xe0, 0xd6, 0x0f, 0x67, 0x1f,
	0x43, 0x28, 0xb0, 0x24, 0xbc, 0x17, 0x2b, 0x9c, 0x3a, 0x25, 0x0f, 0xf8, 0xbb, 0x51, 0x35, 0x47,
	0xba, 0x38, 0xb5, 0x15, 0x3c, 0x48, 0x90, 0xc2, 0x94, 0x64, 0x38, 0x66, 0x24, 0x23, 0x6c, 0xc8,
	0x9c, 0xb2, 0x07, 0xfc, 0x9d, 0xd6, 0xcd, 0xd2, 0xf4, 0x73, 0xd6, 0xb8, 0xea, 0x13, 0x3d, 0x18,
	0x26, 0x41, 0xca, 0x59, 0xf8, 0x6b, 0x30, 0xa3, 0x8b, 0xd3, 0x74, 0x80, 0x48, 0x16, 0xae, 0x91,
	0x9e, 0x7e, 0x11, 0x58, 0x05, 0x5d, 0x2c, 0x09, 0xa2, 0xe4, 0x15, 0x25, 0x14, 0xdf, 0x66, 0x3a,
	0xda, 0x5f, 0x25, 0xdc, 0xe5, 0x01, 0xb6, 0x5f, 0x08, 0xd5, 0x23, 0x1a, 0x0b, 0xc1, 0x9c, 0x8a,
	0x69, 0xb6, 0xb7, 0xc2, 0xef, 0x47, 0xb4, 0x23, 0x58, 0xeb, 0x61, 0x32, 0x77, 0xc1, 0x74, 0xee,
	0x82, 0xaf, 0xb9, 0x0b, 0xde, 0x16, 0xae, 0x35, 0x5d, 0xb8, 0xd6, 0xc7, 0xc2, 0xb5, 0x1e, 0x2f,
	0xff, 0x5f, 0x6b, 0x5c, 0xd8, 0xa1, 0x69, 0x98, 0x6c, 0x1a, 0xee, 0xfc, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x40, 0x12, 0x97, 0xeb, 0xea, 0x01, 0x00, 0x00,
}

func (m *LimitParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Limiters) > 0 {
		for iNdEx := len(m.Limiters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Limiters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLimitParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintLimitParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Limiter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Limiter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Limiter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaselineTvlPpm != 0 {
		i = encodeVarintLimitParams(dAtA, i, uint64(m.BaselineTvlPpm))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.BaselineMinimum.Size()
		i -= size
		if _, err := m.BaselineMinimum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.PeriodSec != 0 {
		i = encodeVarintLimitParams(dAtA, i, uint64(m.PeriodSec))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovLimitParams(uint64(l))
	}
	if len(m.Limiters) > 0 {
		for _, e := range m.Limiters {
			l = e.Size()
			n += 1 + l + sovLimitParams(uint64(l))
		}
	}
	return n
}

func (m *Limiter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PeriodSec != 0 {
		n += 1 + sovLimitParams(uint64(m.PeriodSec))
	}
	l = m.BaselineMinimum.Size()
	n += 1 + l + sovLimitParams(uint64(l))
	if m.BaselineTvlPpm != 0 {
		n += 1 + sovLimitParams(uint64(m.BaselineTvlPpm))
	}
	return n
}

func sovLimitParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitParams(x uint64) (n int) {
	return sovLimitParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitParams
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
			return fmt.Errorf("proto: LimitParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitParams
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
				return ErrInvalidLengthLimitParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limiters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitParams
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
				return ErrInvalidLengthLimitParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Limiters = append(m.Limiters, Limiter{})
			if err := m.Limiters[len(m.Limiters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitParams
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
func (m *Limiter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitParams
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
			return fmt.Errorf("proto: Limiter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Limiter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodSec", wireType)
			}
			m.PeriodSec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeriodSec |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaselineMinimum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLimitParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaselineMinimum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaselineTvlPpm", wireType)
			}
			m.BaselineTvlPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaselineTvlPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLimitParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitParams
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
func skipLimitParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitParams
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
					return 0, ErrIntOverflowLimitParams
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
					return 0, ErrIntOverflowLimitParams
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
				return 0, ErrInvalidLengthLimitParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitParams = fmt.Errorf("proto: unexpected end of group")
)
