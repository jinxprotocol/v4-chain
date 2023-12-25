// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jinxprotocol/subaccounts/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryGetSubaccountRequest is request type for the Query RPC method.
type QueryGetSubaccountRequest struct {
	Owner  string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Number uint32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (m *QueryGetSubaccountRequest) Reset()         { *m = QueryGetSubaccountRequest{} }
func (m *QueryGetSubaccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetSubaccountRequest) ProtoMessage()    {}
func (*QueryGetSubaccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b70f176a60f66338, []int{0}
}
func (m *QueryGetSubaccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSubaccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSubaccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSubaccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSubaccountRequest.Merge(m, src)
}
func (m *QueryGetSubaccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSubaccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSubaccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSubaccountRequest proto.InternalMessageInfo

func (m *QueryGetSubaccountRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *QueryGetSubaccountRequest) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// QuerySubaccountResponse is response type for the Query RPC method.
type QuerySubaccountResponse struct {
	Subaccount Subaccount `protobuf:"bytes,1,opt,name=subaccount,proto3" json:"subaccount"`
}

func (m *QuerySubaccountResponse) Reset()         { *m = QuerySubaccountResponse{} }
func (m *QuerySubaccountResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySubaccountResponse) ProtoMessage()    {}
func (*QuerySubaccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b70f176a60f66338, []int{1}
}
func (m *QuerySubaccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubaccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubaccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubaccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubaccountResponse.Merge(m, src)
}
func (m *QuerySubaccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubaccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubaccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubaccountResponse proto.InternalMessageInfo

func (m *QuerySubaccountResponse) GetSubaccount() Subaccount {
	if m != nil {
		return m.Subaccount
	}
	return Subaccount{}
}

// QueryAllSubaccountRequest is request type for the Query RPC method.
type QueryAllSubaccountRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSubaccountRequest) Reset()         { *m = QueryAllSubaccountRequest{} }
func (m *QueryAllSubaccountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllSubaccountRequest) ProtoMessage()    {}
func (*QueryAllSubaccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b70f176a60f66338, []int{2}
}
func (m *QueryAllSubaccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSubaccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSubaccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSubaccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSubaccountRequest.Merge(m, src)
}
func (m *QueryAllSubaccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSubaccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSubaccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSubaccountRequest proto.InternalMessageInfo

func (m *QueryAllSubaccountRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QuerySubaccountAllResponse is response type for the Query RPC method.
type QuerySubaccountAllResponse struct {
	Subaccount []Subaccount        `protobuf:"bytes,1,rep,name=subaccount,proto3" json:"subaccount"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySubaccountAllResponse) Reset()         { *m = QuerySubaccountAllResponse{} }
func (m *QuerySubaccountAllResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySubaccountAllResponse) ProtoMessage()    {}
func (*QuerySubaccountAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b70f176a60f66338, []int{3}
}
func (m *QuerySubaccountAllResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubaccountAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubaccountAllResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubaccountAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubaccountAllResponse.Merge(m, src)
}
func (m *QuerySubaccountAllResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubaccountAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubaccountAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubaccountAllResponse proto.InternalMessageInfo

func (m *QuerySubaccountAllResponse) GetSubaccount() []Subaccount {
	if m != nil {
		return m.Subaccount
	}
	return nil
}

func (m *QuerySubaccountAllResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetSubaccountRequest)(nil), "jinxprotocol.subaccounts.QueryGetSubaccountRequest")
	proto.RegisterType((*QuerySubaccountResponse)(nil), "jinxprotocol.subaccounts.QuerySubaccountResponse")
	proto.RegisterType((*QueryAllSubaccountRequest)(nil), "jinxprotocol.subaccounts.QueryAllSubaccountRequest")
	proto.RegisterType((*QuerySubaccountAllResponse)(nil), "jinxprotocol.subaccounts.QuerySubaccountAllResponse")
}

func init() {
	proto.RegisterFile("jinxprotocol/subaccounts/query.proto", fileDescriptor_b70f176a60f66338)
}

var fileDescriptor_b70f176a60f66338 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x37, 0xab, 0x2d, 0x18, 0xe9, 0x25, 0x14, 0xdd, 0x0e, 0x32, 0x96, 0x61, 0xa9, 0x55,
	0x6c, 0xc2, 0xb6, 0x15, 0x4f, 0x3d, 0xec, 0x1e, 0x2c, 0x78, 0xd2, 0xe9, 0x41, 0xf0, 0x22, 0x99,
	0x69, 0x48, 0x47, 0x66, 0x93, 0xe9, 0x24, 0x53, 0x5b, 0x4a, 0x2f, 0x7e, 0x02, 0xc1, 0x2f, 0xe1,
	0x55, 0xf4, 0x43, 0xf4, 0x58, 0xf4, 0xe2, 0x49, 0x64, 0xd7, 0x83, 0x1f, 0x43, 0x36, 0x89, 0x3b,
	0x33, 0xae, 0xc3, 0x2e, 0xde, 0x92, 0xbc, 0xf7, 0x7f, 0xef, 0xf7, 0x7f, 0x79, 0xb0, 0xfb, 0x3a,
	0x11, 0xa7, 0x59, 0x2e, 0xb5, 0x8c, 0x65, 0x4a, 0x54, 0x11, 0xd1, 0x38, 0x96, 0x85, 0xd0, 0x8a,
	0x1c, 0x17, 0x2c, 0x3f, 0xc3, 0x26, 0x84, 0x3a, 0xd5, 0x2c, 0x5c, 0xc9, 0xf2, 0xd6, 0x62, 0xa9,
	0x86, 0x52, 0xbd, 0x32, 0x41, 0x62, 0x2f, 0x56, 0xe4, 0xad, 0x72, 0xc9, 0xa5, 0x7d, 0x9f, 0x9c,
	0xdc, 0xeb, 0x1d, 0x2e, 0x25, 0x4f, 0x19, 0xa1, 0x59, 0x42, 0xa8, 0x10, 0x52, 0x53, 0x9d, 0x48,
	0xf1, 0x47, 0xf3, 0xc0, 0x56, 0x20, 0x11, 0x55, 0xcc, 0x12, 0x90, 0x93, 0x5e, 0xc4, 0x34, 0xed,
	0x91, 0x8c, 0xf2, 0x44, 0x98, 0x64, 0x97, 0x7b, 0xbf, 0x11, 0xbd, 0x3c, 0xdb, 0xd4, 0x20, 0x86,
	0x6b, 0xcf, 0x27, 0xc5, 0xf6, 0x99, 0x3e, 0x98, 0xc6, 0x42, 0x76, 0x5c, 0x30, 0xa5, 0x11, 0x86,
	0x4b, 0xf2, 0x8d, 0x60, 0x79, 0x07, 0xac, 0x83, 0xcd, 0x1b, 0x83, 0xce, 0x97, 0xcf, 0x5b, 0xab,
	0xce, 0x48, 0xff, 0xf0, 0x30, 0x67, 0x4a, 0x1d, 0xe8, 0x3c, 0x11, 0x3c, 0xb4, 0x69, 0xe8, 0x16,
	0x5c, 0x16, 0xc5, 0x30, 0x62, 0x79, 0xa7, 0xbd, 0x0e, 0x36, 0x57, 0x42, 0x77, 0x0b, 0x18, 0xbc,
	0x6d, 0x9a, 0x54, 0x3b, 0xa8, 0x4c, 0x0a, 0xc5, 0xd0, 0x53, 0x08, 0x4b, 0x26, 0xd3, 0xe7, 0xe6,
	0x76, 0x17, 0x37, 0x0d, 0x15, 0x97, 0x15, 0x06, 0xd7, 0x2f, 0xbf, 0xdf, 0x6d, 0x85, 0x15, 0xf5,
	0xd4, 0x4b, 0x3f, 0x4d, 0x67, 0xbd, 0x3c, 0x81, 0xb0, 0x9c, 0x93, 0x6b, 0xb4, 0x81, 0x9d, 0x9b,
	0xc9, 0x50, 0xb1, 0xfd, 0x56, 0x37, 0x54, 0xfc, 0x8c, 0x72, 0xe6, 0xb4, 0x61, 0x45, 0x19, 0x7c,
	0x04, 0xd0, 0xfb, 0xcb, 0x4c, 0x3f, 0x4d, 0x1b, 0xfd, 0x5c, 0xfb, 0x7f, 0x3f, 0x68, 0xbf, 0x86,
	0xdc, 0x36, 0xc8, 0xf7, 0xe6, 0x22, 0x5b, 0x90, 0x2a, 0xf3, 0xf6, 0xaf, 0x36, 0x5c, 0x32, 0xcc,
	0xe8, 0x13, 0x80, 0xb0, 0xec, 0x89, 0x76, 0x9a, 0xc9, 0x1a, 0xb7, 0xc2, 0xeb, 0xcd, 0x11, 0xcd,
	0xfe, 0x72, 0xb0, 0xf7, 0xf6, 0xeb, 0xcf, 0xf7, 0xed, 0xc7, 0xe8, 0x11, 0x59, 0x60, 0x33, 0xc9,
	0xb9, 0xd9, 0xa6, 0x0b, 0x72, 0x6e, 0xd7, 0xe7, 0x02, 0x7d, 0x00, 0x70, 0xa5, 0x36, 0xee, 0xb9,
	0xe0, 0xff, 0x5a, 0x01, 0x6f, 0x77, 0x61, 0xf0, 0xca, 0x8f, 0x06, 0x0f, 0x0d, 0xfb, 0x06, 0xea,
	0x2e, 0xc2, 0x3e, 0x78, 0x71, 0x39, 0xf2, 0xc1, 0xd5, 0xc8, 0x07, 0x3f, 0x46, 0x3e, 0x78, 0x37,
	0xf6, 0x5b, 0x57, 0x63, 0xbf, 0xf5, 0x6d, 0xec, 0xb7, 0x5e, 0xee, 0xf1, 0x44, 0x1f, 0x15, 0x11,
	0x8e, 0xe5, 0xb0, 0x5e, 0xe9, 0x64, 0x77, 0x2b, 0x3e, 0xa2, 0x89, 0x20, 0xd3, 0x97, 0xd3, 0x5a,
	0x75, 0x7d, 0x96, 0x31, 0x15, 0x2d, 0x9b, 0xe8, 0xce, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33,
	0xd3, 0x0d, 0x9b, 0x97, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a Subaccount by id
	Subaccount(ctx context.Context, in *QueryGetSubaccountRequest, opts ...grpc.CallOption) (*QuerySubaccountResponse, error)
	// Queries a list of Subaccount items.
	SubaccountAll(ctx context.Context, in *QueryAllSubaccountRequest, opts ...grpc.CallOption) (*QuerySubaccountAllResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Subaccount(ctx context.Context, in *QueryGetSubaccountRequest, opts ...grpc.CallOption) (*QuerySubaccountResponse, error) {
	out := new(QuerySubaccountResponse)
	err := c.cc.Invoke(ctx, "/jinxprotocol.subaccounts.Query/Subaccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SubaccountAll(ctx context.Context, in *QueryAllSubaccountRequest, opts ...grpc.CallOption) (*QuerySubaccountAllResponse, error) {
	out := new(QuerySubaccountAllResponse)
	err := c.cc.Invoke(ctx, "/jinxprotocol.subaccounts.Query/SubaccountAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a Subaccount by id
	Subaccount(context.Context, *QueryGetSubaccountRequest) (*QuerySubaccountResponse, error)
	// Queries a list of Subaccount items.
	SubaccountAll(context.Context, *QueryAllSubaccountRequest) (*QuerySubaccountAllResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Subaccount(ctx context.Context, req *QueryGetSubaccountRequest) (*QuerySubaccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subaccount not implemented")
}
func (*UnimplementedQueryServer) SubaccountAll(ctx context.Context, req *QueryAllSubaccountRequest) (*QuerySubaccountAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubaccountAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Subaccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSubaccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Subaccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jinxprotocol.subaccounts.Query/Subaccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Subaccount(ctx, req.(*QueryGetSubaccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SubaccountAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSubaccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SubaccountAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jinxprotocol.subaccounts.Query/SubaccountAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SubaccountAll(ctx, req.(*QueryAllSubaccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jinxprotocol.subaccounts.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subaccount",
			Handler:    _Query_Subaccount_Handler,
		},
		{
			MethodName: "SubaccountAll",
			Handler:    _Query_SubaccountAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jinxprotocol/subaccounts/query.proto",
}

func (m *QueryGetSubaccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSubaccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSubaccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySubaccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubaccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubaccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Subaccount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllSubaccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSubaccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSubaccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySubaccountAllResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubaccountAllResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubaccountAllResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subaccount) > 0 {
		for iNdEx := len(m.Subaccount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subaccount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetSubaccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Number != 0 {
		n += 1 + sovQuery(uint64(m.Number))
	}
	return n
}

func (m *QuerySubaccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Subaccount.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllSubaccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySubaccountAllResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Subaccount) > 0 {
		for _, e := range m.Subaccount {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetSubaccountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetSubaccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSubaccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QuerySubaccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QuerySubaccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubaccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subaccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Subaccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllSubaccountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllSubaccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSubaccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QuerySubaccountAllResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QuerySubaccountAllResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubaccountAllResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subaccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subaccount = append(m.Subaccount, Subaccount{})
			if err := m.Subaccount[len(m.Subaccount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
