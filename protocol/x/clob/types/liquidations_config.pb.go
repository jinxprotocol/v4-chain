// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jinxprotocol/clob/liquidations_config.proto

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

// LiquidationsConfig stores all configurable fields related to liquidations.
type LiquidationsConfig struct {
	// The maximum liquidation fee (in parts-per-million). This fee goes
	// 100% to the insurance fund.
	MaxLiquidationFeePpm uint32 `protobuf:"varint,1,opt,name=max_liquidation_fee_ppm,json=maxLiquidationFeePpm,proto3" json:"max_liquidation_fee_ppm,omitempty"`
	// Limits around how much of a single position can be liquidated
	// within a single block.
	PositionBlockLimits PositionBlockLimits `protobuf:"bytes,2,opt,name=position_block_limits,json=positionBlockLimits,proto3" json:"position_block_limits"`
	// Limits around how many quote quantums from a single subaccount can
	// be liquidated within a single block.
	SubaccountBlockLimits SubaccountBlockLimits `protobuf:"bytes,3,opt,name=subaccount_block_limits,json=subaccountBlockLimits,proto3" json:"subaccount_block_limits"`
	// Config about how the fillable-price spread from the oracle price
	// increases based on the adjusted bankruptcy rating of the subaccount.
	FillablePriceConfig FillablePriceConfig `protobuf:"bytes,4,opt,name=fillable_price_config,json=fillablePriceConfig,proto3" json:"fillable_price_config"`
}

func (m *LiquidationsConfig) Reset()         { *m = LiquidationsConfig{} }
func (m *LiquidationsConfig) String() string { return proto.CompactTextString(m) }
func (*LiquidationsConfig) ProtoMessage()    {}
func (*LiquidationsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc32f56c88e1e215, []int{0}
}
func (m *LiquidationsConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidationsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidationsConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidationsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidationsConfig.Merge(m, src)
}
func (m *LiquidationsConfig) XXX_Size() int {
	return m.Size()
}
func (m *LiquidationsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidationsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidationsConfig proto.InternalMessageInfo

func (m *LiquidationsConfig) GetMaxLiquidationFeePpm() uint32 {
	if m != nil {
		return m.MaxLiquidationFeePpm
	}
	return 0
}

func (m *LiquidationsConfig) GetPositionBlockLimits() PositionBlockLimits {
	if m != nil {
		return m.PositionBlockLimits
	}
	return PositionBlockLimits{}
}

func (m *LiquidationsConfig) GetSubaccountBlockLimits() SubaccountBlockLimits {
	if m != nil {
		return m.SubaccountBlockLimits
	}
	return SubaccountBlockLimits{}
}

func (m *LiquidationsConfig) GetFillablePriceConfig() FillablePriceConfig {
	if m != nil {
		return m.FillablePriceConfig
	}
	return FillablePriceConfig{}
}

// PositionBlockLimits stores all configurable fields related to limits
// around how much of a single position can be liquidated within a single block.
type PositionBlockLimits struct {
	// The minimum amount of quantums to liquidate for each message (in
	// quote quantums).
	// Overridden by the maximum size of the position.
	MinPositionNotionalLiquidated uint64 `protobuf:"varint,1,opt,name=min_position_notional_liquidated,json=minPositionNotionalLiquidated,proto3" json:"min_position_notional_liquidated,omitempty"`
	// The maximum portion of the position liquidated (in parts-per-
	// million). Overridden by min_position_notional_liquidated.
	MaxPositionPortionLiquidatedPpm uint32 `protobuf:"varint,2,opt,name=max_position_portion_liquidated_ppm,json=maxPositionPortionLiquidatedPpm,proto3" json:"max_position_portion_liquidated_ppm,omitempty"`
}

func (m *PositionBlockLimits) Reset()         { *m = PositionBlockLimits{} }
func (m *PositionBlockLimits) String() string { return proto.CompactTextString(m) }
func (*PositionBlockLimits) ProtoMessage()    {}
func (*PositionBlockLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc32f56c88e1e215, []int{1}
}
func (m *PositionBlockLimits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PositionBlockLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PositionBlockLimits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PositionBlockLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionBlockLimits.Merge(m, src)
}
func (m *PositionBlockLimits) XXX_Size() int {
	return m.Size()
}
func (m *PositionBlockLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionBlockLimits.DiscardUnknown(m)
}

var xxx_messageInfo_PositionBlockLimits proto.InternalMessageInfo

func (m *PositionBlockLimits) GetMinPositionNotionalLiquidated() uint64 {
	if m != nil {
		return m.MinPositionNotionalLiquidated
	}
	return 0
}

func (m *PositionBlockLimits) GetMaxPositionPortionLiquidatedPpm() uint32 {
	if m != nil {
		return m.MaxPositionPortionLiquidatedPpm
	}
	return 0
}

// SubaccountBlockLimits stores all configurable fields related to limits
// around how many quote quantums from a single subaccount can
// be liquidated within a single block.
type SubaccountBlockLimits struct {
	// The maximum notional amount that a single subaccount can have
	// liquidated (in quote quantums) per block.
	MaxNotionalLiquidated uint64 `protobuf:"varint,1,opt,name=max_notional_liquidated,json=maxNotionalLiquidated,proto3" json:"max_notional_liquidated,omitempty"`
	// The maximum insurance-fund payout amount for a given subaccount
	// per block. I.e. how much it can cover for that subaccount.
	MaxQuantumsInsuranceLost uint64 `protobuf:"varint,2,opt,name=max_quantums_insurance_lost,json=maxQuantumsInsuranceLost,proto3" json:"max_quantums_insurance_lost,omitempty"`
}

func (m *SubaccountBlockLimits) Reset()         { *m = SubaccountBlockLimits{} }
func (m *SubaccountBlockLimits) String() string { return proto.CompactTextString(m) }
func (*SubaccountBlockLimits) ProtoMessage()    {}
func (*SubaccountBlockLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc32f56c88e1e215, []int{2}
}
func (m *SubaccountBlockLimits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubaccountBlockLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubaccountBlockLimits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubaccountBlockLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubaccountBlockLimits.Merge(m, src)
}
func (m *SubaccountBlockLimits) XXX_Size() int {
	return m.Size()
}
func (m *SubaccountBlockLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_SubaccountBlockLimits.DiscardUnknown(m)
}

var xxx_messageInfo_SubaccountBlockLimits proto.InternalMessageInfo

func (m *SubaccountBlockLimits) GetMaxNotionalLiquidated() uint64 {
	if m != nil {
		return m.MaxNotionalLiquidated
	}
	return 0
}

func (m *SubaccountBlockLimits) GetMaxQuantumsInsuranceLost() uint64 {
	if m != nil {
		return m.MaxQuantumsInsuranceLost
	}
	return 0
}

// FillablePriceConfig stores all configurable fields related to calculating
// the fillable price for liquidating a position.
type FillablePriceConfig struct {
	// The rate at which the Adjusted Bankruptcy Rating increases.
	BankruptcyAdjustmentPpm uint32 `protobuf:"varint,1,opt,name=bankruptcy_adjustment_ppm,json=bankruptcyAdjustmentPpm,proto3" json:"bankruptcy_adjustment_ppm,omitempty"`
	// The maximum value that the liquidation spread can take, as
	// a ratio against the position's maintenance margin.
	SpreadToMaintenanceMarginRatioPpm uint32 `protobuf:"varint,2,opt,name=spread_to_maintenance_margin_ratio_ppm,json=spreadToMaintenanceMarginRatioPpm,proto3" json:"spread_to_maintenance_margin_ratio_ppm,omitempty"`
}

func (m *FillablePriceConfig) Reset()         { *m = FillablePriceConfig{} }
func (m *FillablePriceConfig) String() string { return proto.CompactTextString(m) }
func (*FillablePriceConfig) ProtoMessage()    {}
func (*FillablePriceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc32f56c88e1e215, []int{3}
}
func (m *FillablePriceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FillablePriceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FillablePriceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FillablePriceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FillablePriceConfig.Merge(m, src)
}
func (m *FillablePriceConfig) XXX_Size() int {
	return m.Size()
}
func (m *FillablePriceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FillablePriceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FillablePriceConfig proto.InternalMessageInfo

func (m *FillablePriceConfig) GetBankruptcyAdjustmentPpm() uint32 {
	if m != nil {
		return m.BankruptcyAdjustmentPpm
	}
	return 0
}

func (m *FillablePriceConfig) GetSpreadToMaintenanceMarginRatioPpm() uint32 {
	if m != nil {
		return m.SpreadToMaintenanceMarginRatioPpm
	}
	return 0
}

func init() {
	proto.RegisterType((*LiquidationsConfig)(nil), "jinxprotocol.clob.LiquidationsConfig")
	proto.RegisterType((*PositionBlockLimits)(nil), "jinxprotocol.clob.PositionBlockLimits")
	proto.RegisterType((*SubaccountBlockLimits)(nil), "jinxprotocol.clob.SubaccountBlockLimits")
	proto.RegisterType((*FillablePriceConfig)(nil), "jinxprotocol.clob.FillablePriceConfig")
}

func init() {
	proto.RegisterFile("jinxprotocol/clob/liquidations_config.proto", fileDescriptor_cc32f56c88e1e215)
}

var fileDescriptor_cc32f56c88e1e215 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe3, 0x12, 0x71, 0x58, 0xc4, 0x01, 0xa7, 0x51, 0x02, 0x08, 0x37, 0x04, 0xa9, 0x8a,
	0x84, 0x48, 0x24, 0x3e, 0x7a, 0x40, 0xe2, 0x40, 0x90, 0x8a, 0x90, 0x52, 0xe4, 0x06, 0x4e, 0x5c,
	0x96, 0xf5, 0x66, 0xed, 0x6e, 0xbb, 0x5f, 0xf5, 0xae, 0x91, 0xfb, 0x12, 0x88, 0x87, 0xe0, 0xc8,
	0x83, 0xf4, 0xd8, 0x23, 0x27, 0x84, 0x92, 0x87, 0xe0, 0x8a, 0x76, 0xed, 0x38, 0xae, 0xe2, 0x9e,
	0x6c, 0xed, 0xfc, 0xf6, 0x3f, 0x33, 0xff, 0x99, 0x05, 0x4f, 0x4f, 0xa9, 0xc8, 0x55, 0x2a, 0x8d,
	0xc4, 0x92, 0x4d, 0x30, 0x93, 0xd1, 0x84, 0xd1, 0xf3, 0x8c, 0x2e, 0x90, 0xa1, 0x52, 0x68, 0x88,
	0xa5, 0x88, 0x69, 0x32, 0x76, 0x84, 0x7f, 0xaf, 0x0e, 0x8f, 0x2d, 0xfc, 0x60, 0x37, 0x91, 0x89,
	0x74, 0x47, 0x13, 0xfb, 0x57, 0x80, 0xc3, 0x7f, 0x3b, 0xc0, 0x9f, 0xd5, 0x64, 0xde, 0x39, 0x15,
	0xff, 0x15, 0xe8, 0x71, 0x94, 0xc3, 0x5a, 0x02, 0x18, 0x13, 0x02, 0x95, 0xe2, 0x7d, 0x6f, 0xe0,
	0x8d, 0xee, 0xce, 0x77, 0x39, 0xca, 0x6b, 0xf7, 0x0e, 0x09, 0x09, 0x15, 0xf7, 0xbf, 0x82, 0xae,
	0x92, 0x9a, 0x3a, 0x3e, 0x62, 0x12, 0x9f, 0x41, 0x46, 0x39, 0x35, 0xba, 0xbf, 0x33, 0xf0, 0x46,
	0x77, 0x9e, 0xef, 0x8f, 0xb7, 0xca, 0x1a, 0x87, 0x25, 0x3f, 0xb5, 0xf8, 0xcc, 0xd1, 0xd3, 0xf6,
	0xe5, 0x9f, 0xbd, 0xd6, 0xbc, 0xa3, 0xb6, 0x43, 0x7e, 0x0c, 0x7a, 0x3a, 0x8b, 0x10, 0xc6, 0x32,
	0x13, 0xe6, 0x7a, 0x8e, 0x5b, 0x2e, 0xc7, 0xa8, 0x21, 0xc7, 0xa7, 0xea, 0xc6, 0x76, 0x96, 0xae,
	0x6e, 0x0a, 0xda, 0x4e, 0x62, 0xca, 0x18, 0x8a, 0x18, 0x81, 0x2a, 0xa5, 0x98, 0x94, 0xfe, 0xf6,
	0xdb, 0x37, 0x76, 0x72, 0x58, 0xf2, 0xa1, 0xc5, 0x0b, 0x1f, 0xd7, 0x9d, 0xc4, 0xdb, 0xa1, 0xe1,
	0x2f, 0x0f, 0x74, 0x1a, 0x9a, 0xf7, 0xdf, 0x83, 0x01, 0xa7, 0x02, 0x56, 0x3e, 0x0a, 0x69, 0x3f,
	0x88, 0x55, 0xc3, 0x20, 0x0b, 0x37, 0x83, 0xf6, 0xfc, 0x11, 0xa7, 0x62, 0xad, 0xf0, 0xb1, 0xa4,
	0x66, 0x15, 0xe4, 0xcf, 0xc0, 0x13, 0x3b, 0xc3, 0x4a, 0x48, 0xc9, 0xd4, 0x7d, 0x37, 0x3a, 0x6e,
	0x9e, 0x3b, 0x6e, 0x9e, 0x7b, 0x1c, 0xe5, 0x6b, 0xad, 0xb0, 0x00, 0x37, 0x52, 0xa1, 0xe2, 0xc3,
	0xef, 0x1e, 0xe8, 0x36, 0xfa, 0xe8, 0x1f, 0x14, 0xbb, 0x72, 0x73, 0x9d, 0x5d, 0x8e, 0xf2, 0x86,
	0xfa, 0xde, 0x80, 0x87, 0xf6, 0xde, 0x79, 0x86, 0x84, 0xc9, 0xb8, 0x86, 0x54, 0xe8, 0x2c, 0x45,
	0x02, 0x13, 0xc8, 0xa4, 0x36, 0xae, 0xae, 0xf6, 0xbc, 0xcf, 0x51, 0x7e, 0x5c, 0x12, 0x1f, 0xd6,
	0xc0, 0x4c, 0x6a, 0x33, 0xfc, 0xe9, 0x81, 0x4e, 0x83, 0xe5, 0xfe, 0x6b, 0x70, 0x3f, 0x42, 0xe2,
	0x2c, 0xcd, 0x94, 0xc1, 0x17, 0x10, 0x2d, 0x4e, 0x33, 0x6d, 0x38, 0x11, 0xa6, 0xb6, 0xbc, 0xbd,
	0x0d, 0xf0, 0xb6, 0x8a, 0xdb, 0xfd, 0x3d, 0x06, 0xfb, 0x5a, 0xa5, 0x04, 0x2d, 0xa0, 0x91, 0x90,
	0x23, 0x2a, 0x0c, 0x11, 0xae, 0x22, 0x8e, 0xd2, 0x84, 0x0a, 0x98, 0xda, 0x65, 0xaf, 0xb9, 0xf6,
	0xb8, 0xa0, 0x3f, 0xcb, 0xa3, 0x0d, 0x7b, 0xe4, 0xd0, 0xb9, 0x25, 0x43, 0xc5, 0xa7, 0xe1, 0xe5,
	0x32, 0xf0, 0xae, 0x96, 0x81, 0xf7, 0x77, 0x19, 0x78, 0x3f, 0x56, 0x41, 0xeb, 0x6a, 0x15, 0xb4,
	0x7e, 0xaf, 0x82, 0xd6, 0x97, 0x83, 0x84, 0x9a, 0x93, 0x2c, 0x1a, 0x63, 0xc9, 0x27, 0xd7, 0xde,
	0xf6, 0xb7, 0x97, 0xcf, 0xf0, 0x09, 0xa2, 0x62, 0x52, 0x9d, 0xe4, 0xc5, 0x7b, 0x37, 0x17, 0x8a,
	0xe8, 0xe8, 0xb6, 0x3b, 0x7e, 0xf1, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x43, 0xc6, 0x17, 0x11,
	0x04, 0x00, 0x00,
}

func (m *LiquidationsConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidationsConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidationsConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FillablePriceConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.SubaccountBlockLimits.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PositionBlockLimits.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.MaxLiquidationFeePpm != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MaxLiquidationFeePpm))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PositionBlockLimits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PositionBlockLimits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PositionBlockLimits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxPositionPortionLiquidatedPpm != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MaxPositionPortionLiquidatedPpm))
		i--
		dAtA[i] = 0x10
	}
	if m.MinPositionNotionalLiquidated != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MinPositionNotionalLiquidated))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SubaccountBlockLimits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubaccountBlockLimits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubaccountBlockLimits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxQuantumsInsuranceLost != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MaxQuantumsInsuranceLost))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxNotionalLiquidated != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MaxNotionalLiquidated))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FillablePriceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FillablePriceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FillablePriceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SpreadToMaintenanceMarginRatioPpm != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.SpreadToMaintenanceMarginRatioPpm))
		i--
		dAtA[i] = 0x10
	}
	if m.BankruptcyAdjustmentPpm != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.BankruptcyAdjustmentPpm))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidationsConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidationsConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiquidationsConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxLiquidationFeePpm != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MaxLiquidationFeePpm))
	}
	l = m.PositionBlockLimits.Size()
	n += 1 + l + sovLiquidationsConfig(uint64(l))
	l = m.SubaccountBlockLimits.Size()
	n += 1 + l + sovLiquidationsConfig(uint64(l))
	l = m.FillablePriceConfig.Size()
	n += 1 + l + sovLiquidationsConfig(uint64(l))
	return n
}

func (m *PositionBlockLimits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinPositionNotionalLiquidated != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MinPositionNotionalLiquidated))
	}
	if m.MaxPositionPortionLiquidatedPpm != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MaxPositionPortionLiquidatedPpm))
	}
	return n
}

func (m *SubaccountBlockLimits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxNotionalLiquidated != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MaxNotionalLiquidated))
	}
	if m.MaxQuantumsInsuranceLost != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MaxQuantumsInsuranceLost))
	}
	return n
}

func (m *FillablePriceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BankruptcyAdjustmentPpm != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.BankruptcyAdjustmentPpm))
	}
	if m.SpreadToMaintenanceMarginRatioPpm != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.SpreadToMaintenanceMarginRatioPpm))
	}
	return n
}

func sovLiquidationsConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidationsConfig(x uint64) (n int) {
	return sovLiquidationsConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiquidationsConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidationsConfig
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
			return fmt.Errorf("proto: LiquidationsConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidationsConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLiquidationFeePpm", wireType)
			}
			m.MaxLiquidationFeePpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxLiquidationFeePpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionBlockLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
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
				return ErrInvalidLengthLiquidationsConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidationsConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PositionBlockLimits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountBlockLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
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
				return ErrInvalidLengthLiquidationsConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidationsConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubaccountBlockLimits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FillablePriceConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
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
				return ErrInvalidLengthLiquidationsConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidationsConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FillablePriceConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidationsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidationsConfig
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
func (m *PositionBlockLimits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidationsConfig
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
			return fmt.Errorf("proto: PositionBlockLimits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PositionBlockLimits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPositionNotionalLiquidated", wireType)
			}
			m.MinPositionNotionalLiquidated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinPositionNotionalLiquidated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPositionPortionLiquidatedPpm", wireType)
			}
			m.MaxPositionPortionLiquidatedPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPositionPortionLiquidatedPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidationsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidationsConfig
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
func (m *SubaccountBlockLimits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidationsConfig
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
			return fmt.Errorf("proto: SubaccountBlockLimits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubaccountBlockLimits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNotionalLiquidated", wireType)
			}
			m.MaxNotionalLiquidated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNotionalLiquidated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxQuantumsInsuranceLost", wireType)
			}
			m.MaxQuantumsInsuranceLost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxQuantumsInsuranceLost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidationsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidationsConfig
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
func (m *FillablePriceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidationsConfig
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
			return fmt.Errorf("proto: FillablePriceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FillablePriceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BankruptcyAdjustmentPpm", wireType)
			}
			m.BankruptcyAdjustmentPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BankruptcyAdjustmentPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadToMaintenanceMarginRatioPpm", wireType)
			}
			m.SpreadToMaintenanceMarginRatioPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpreadToMaintenanceMarginRatioPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidationsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidationsConfig
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
func skipLiquidationsConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidationsConfig
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
					return 0, ErrIntOverflowLiquidationsConfig
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
					return 0, ErrIntOverflowLiquidationsConfig
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
				return 0, ErrInvalidLengthLiquidationsConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidationsConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidationsConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidationsConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidationsConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidationsConfig = fmt.Errorf("proto: unexpected end of group")
)
