// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol_x/query.proto

package types

import (
	context "context"
	fmt "fmt"
	accesscontrol "github.com/cosmos/cosmos-sdk/types/accesscontrol"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type ResourceDepedencyMappingFromMessageKeyRequest struct {
	MessageKey string `protobuf:"bytes,1,opt,name=message_key,json=messageKey,proto3" json:"message_key,omitempty" yaml:"message_key"`
}

func (m *ResourceDepedencyMappingFromMessageKeyRequest) Reset() {
	*m = ResourceDepedencyMappingFromMessageKeyRequest{}
}
func (m *ResourceDepedencyMappingFromMessageKeyRequest) String() string {
	return proto.CompactTextString(m)
}
func (*ResourceDepedencyMappingFromMessageKeyRequest) ProtoMessage() {}
func (*ResourceDepedencyMappingFromMessageKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{2}
}
func (m *ResourceDepedencyMappingFromMessageKeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceDepedencyMappingFromMessageKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceDepedencyMappingFromMessageKeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceDepedencyMappingFromMessageKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDepedencyMappingFromMessageKeyRequest.Merge(m, src)
}
func (m *ResourceDepedencyMappingFromMessageKeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *ResourceDepedencyMappingFromMessageKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDepedencyMappingFromMessageKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDepedencyMappingFromMessageKeyRequest proto.InternalMessageInfo

func (m *ResourceDepedencyMappingFromMessageKeyRequest) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

type ResourceDepedencyMappingFromMessageKeyResponse struct {
	MessageDependencyMapping accesscontrol.MessageDependencyMapping `protobuf:"bytes,1,opt,name=message_dependency_mapping,json=messageDependencyMapping,proto3" json:"message_dependency_mapping" yaml:"message_dependency_mapping"`
}

func (m *ResourceDepedencyMappingFromMessageKeyResponse) Reset() {
	*m = ResourceDepedencyMappingFromMessageKeyResponse{}
}
func (m *ResourceDepedencyMappingFromMessageKeyResponse) String() string {
	return proto.CompactTextString(m)
}
func (*ResourceDepedencyMappingFromMessageKeyResponse) ProtoMessage() {}
func (*ResourceDepedencyMappingFromMessageKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{3}
}
func (m *ResourceDepedencyMappingFromMessageKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceDepedencyMappingFromMessageKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceDepedencyMappingFromMessageKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceDepedencyMappingFromMessageKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceDepedencyMappingFromMessageKeyResponse.Merge(m, src)
}
func (m *ResourceDepedencyMappingFromMessageKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *ResourceDepedencyMappingFromMessageKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceDepedencyMappingFromMessageKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceDepedencyMappingFromMessageKeyResponse proto.InternalMessageInfo

func (m *ResourceDepedencyMappingFromMessageKeyResponse) GetMessageDependencyMapping() accesscontrol.MessageDependencyMapping {
	if m != nil {
		return m.MessageDependencyMapping
	}
	return accesscontrol.MessageDependencyMapping{}
}

type ListResourceDepedencyMappingRequest struct {
}

func (m *ListResourceDepedencyMappingRequest) Reset()         { *m = ListResourceDepedencyMappingRequest{} }
func (m *ListResourceDepedencyMappingRequest) String() string { return proto.CompactTextString(m) }
func (*ListResourceDepedencyMappingRequest) ProtoMessage()    {}
func (*ListResourceDepedencyMappingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{4}
}
func (m *ListResourceDepedencyMappingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListResourceDepedencyMappingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListResourceDepedencyMappingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListResourceDepedencyMappingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourceDepedencyMappingRequest.Merge(m, src)
}
func (m *ListResourceDepedencyMappingRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListResourceDepedencyMappingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourceDepedencyMappingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourceDepedencyMappingRequest proto.InternalMessageInfo

type ListResourceDepedencyMappingResponse struct {
	MessageDependencyMappingList []accesscontrol.MessageDependencyMapping `protobuf:"bytes,1,rep,name=message_dependency_mapping_list,json=messageDependencyMappingList,proto3" json:"message_dependency_mapping_list" yaml:"message_dependency_mapping_list"`
}

func (m *ListResourceDepedencyMappingResponse) Reset()         { *m = ListResourceDepedencyMappingResponse{} }
func (m *ListResourceDepedencyMappingResponse) String() string { return proto.CompactTextString(m) }
func (*ListResourceDepedencyMappingResponse) ProtoMessage()    {}
func (*ListResourceDepedencyMappingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d83f2274e13e6a16, []int{5}
}
func (m *ListResourceDepedencyMappingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListResourceDepedencyMappingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListResourceDepedencyMappingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListResourceDepedencyMappingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResourceDepedencyMappingResponse.Merge(m, src)
}
func (m *ListResourceDepedencyMappingResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListResourceDepedencyMappingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResourceDepedencyMappingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResourceDepedencyMappingResponse proto.InternalMessageInfo

func (m *ListResourceDepedencyMappingResponse) GetMessageDependencyMappingList() []accesscontrol.MessageDependencyMapping {
	if m != nil {
		return m.MessageDependencyMappingList
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.accesscontrol_x.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.accesscontrol_x.v1beta1.QueryParamsResponse")
	proto.RegisterType((*ResourceDepedencyMappingFromMessageKeyRequest)(nil), "cosmos.accesscontrol_x.v1beta1.ResourceDepedencyMappingFromMessageKeyRequest")
	proto.RegisterType((*ResourceDepedencyMappingFromMessageKeyResponse)(nil), "cosmos.accesscontrol_x.v1beta1.ResourceDepedencyMappingFromMessageKeyResponse")
	proto.RegisterType((*ListResourceDepedencyMappingRequest)(nil), "cosmos.accesscontrol_x.v1beta1.ListResourceDepedencyMappingRequest")
	proto.RegisterType((*ListResourceDepedencyMappingResponse)(nil), "cosmos.accesscontrol_x.v1beta1.ListResourceDepedencyMappingResponse")
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol_x/query.proto", fileDescriptor_d83f2274e13e6a16)
}

var fileDescriptor_d83f2274e13e6a16 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xd1, 0x6a, 0x13, 0x4d,
	0x14, 0xc7, 0x33, 0xfd, 0xbe, 0x06, 0x9c, 0xde, 0x8d, 0x45, 0xc2, 0x12, 0x36, 0x75, 0xac, 0xb1,
	0x15, 0xba, 0x4b, 0x53, 0x50, 0xf0, 0x4a, 0x62, 0x10, 0xc1, 0x46, 0x74, 0x2f, 0xab, 0xb0, 0x4c,
	0x36, 0xc7, 0xed, 0xd2, 0xec, 0xce, 0x76, 0x67, 0x23, 0x5d, 0xc4, 0x1b, 0x9f, 0x40, 0xf0, 0x01,
	0x04, 0x6f, 0xc5, 0xf7, 0xe8, 0x65, 0x50, 0x10, 0xaf, 0x82, 0x24, 0x3e, 0x41, 0x7d, 0x01, 0xc9,
	0xec, 0x34, 0x24, 0x69, 0x36, 0x9b, 0x52, 0xaf, 0x92, 0x9d, 0x3d, 0xe7, 0xfc, 0xcf, 0xef, 0xcc,
	0xff, 0x2c, 0xa6, 0x0e, 0x17, 0x3e, 0x17, 0x26, 0x73, 0x1c, 0x10, 0xc2, 0xe1, 0x41, 0x1c, 0xf1,
	0x8e, 0x7d, 0x62, 0x1e, 0x77, 0x21, 0x4a, 0x8c, 0x30, 0xe2, 0x31, 0x27, 0x7a, 0x1a, 0x63, 0xcc,
	0xc4, 0x18, 0x6f, 0x76, 0x5b, 0x10, 0xb3, 0x5d, 0x6d, 0xdd, 0xe5, 0x2e, 0x97, 0xa1, 0xe6, 0xe8,
	0x5f, 0x9a, 0xa5, 0x95, 0x5d, 0xce, 0xdd, 0x0e, 0x98, 0x2c, 0xf4, 0x4c, 0x16, 0x04, 0x3c, 0x66,
	0xb1, 0xc7, 0x03, 0xa1, 0xde, 0xde, 0x55, 0xba, 0x2d, 0x26, 0x20, 0x15, 0x33, 0x55, 0x39, 0x33,
	0x64, 0xae, 0x17, 0xc8, 0x60, 0x15, 0xbb, 0x35, 0xaf, 0xc7, 0xe9, 0x27, 0x15, 0xb9, 0x99, 0x41,
	0xe3, 0x42, 0x00, 0xc2, 0x53, 0xda, 0x74, 0x1d, 0x93, 0x17, 0x23, 0xc5, 0xe7, 0x2c, 0x62, 0xbe,
	0xb0, 0xe0, 0xb8, 0x0b, 0x22, 0xa6, 0x2f, 0xf1, 0xf5, 0xa9, 0x53, 0x11, 0xf2, 0x40, 0x00, 0x69,
	0xe0, 0x62, 0x28, 0x4f, 0x4a, 0x68, 0x03, 0x6d, 0xad, 0xd5, 0xaa, 0xc6, 0xe2, 0x69, 0x18, 0x69,
	0x7e, 0xfd, 0xff, 0xd3, 0x7e, 0xa5, 0x60, 0xa9, 0x5c, 0x7a, 0x88, 0x77, 0x2c, 0x10, 0xbc, 0x1b,
	0x39, 0xd0, 0x80, 0x10, 0xda, 0x10, 0x38, 0x49, 0x93, 0x85, 0xa1, 0x17, 0xb8, 0x8f, 0x23, 0xee,
	0x37, 0x41, 0x08, 0xe6, 0xc2, 0x53, 0x48, 0x54, 0x37, 0xe4, 0x3e, 0x5e, 0xf3, 0xd3, 0x43, 0xfb,
	0x08, 0x12, 0xa9, 0x7d, 0xad, 0x7e, 0xe3, 0xac, 0x5f, 0x21, 0x09, 0xf3, 0x3b, 0x0f, 0xe8, 0xc4,
	0x4b, 0x6a, 0x61, 0x7f, 0x9c, 0x4f, 0xbf, 0x21, 0x6c, 0x2c, 0x2b, 0xa5, 0x10, 0x3f, 0x21, 0xac,
	0x9d, 0xd7, 0x6b, 0x43, 0x08, 0x81, 0xcc, 0xb1, 0xfd, 0x34, 0x49, 0x71, 0xdf, 0x9b, 0xcb, 0x3d,
	0xa6, 0x56, 0x65, 0x1b, 0xe3, 0x74, 0x25, 0x59, 0xdf, 0x1e, 0xcd, 0xe1, 0xac, 0x5f, 0xb9, 0x39,
	0xdd, 0xf7, 0x45, 0x1d, 0x6a, 0x95, 0xfc, 0x8c, 0x22, 0xf4, 0x36, 0xbe, 0xb5, 0xef, 0x89, 0x38,
	0x8b, 0xeb, 0xfc, 0x0a, 0x7f, 0x20, 0xbc, 0xb9, 0x38, 0x4e, 0x11, 0x7f, 0x45, 0xb8, 0x92, 0xdd,
	0x89, 0xdd, 0xf1, 0x44, 0x5c, 0x42, 0x1b, 0xff, 0x5d, 0x01, 0xdb, 0x50, 0xd8, 0xd5, 0x3c, 0x6c,
	0x29, 0x46, 0xad, 0x72, 0x16, 0xfb, 0x88, 0xa7, 0xd6, 0x5b, 0xc5, 0xab, 0xd2, 0x9c, 0xe4, 0x33,
	0xc2, 0xc5, 0xd4, 0x61, 0xa4, 0x96, 0xe7, 0xc4, 0x8b, 0x26, 0xd7, 0xf6, 0x2e, 0x95, 0x93, 0x4e,
	0x8b, 0x9a, 0xef, 0xbf, 0xff, 0xfe, 0xb8, 0xb2, 0x4d, 0xee, 0x98, 0x6a, 0xbd, 0xd2, 0x9f, 0x1d,
	0xd1, 0x3e, 0x9a, 0xd9, 0xc9, 0xd4, 0xed, 0xe4, 0xcb, 0x0a, 0xae, 0x2e, 0xe7, 0x41, 0xd2, 0xcc,
	0x6b, 0xe8, 0x52, 0x6b, 0xa3, 0x3d, 0xfb, 0x57, 0xe5, 0x14, 0x7a, 0x4b, 0xa2, 0xbf, 0x22, 0x07,
	0xb9, 0xe8, 0x91, 0x2a, 0x2c, 0xaf, 0x78, 0xfa, 0x86, 0x5f, 0x47, 0xdc, 0xb7, 0x27, 0x16, 0xd6,
	0x7c, 0x3b, 0xf1, 0xf0, 0x8e, 0xfc, 0x41, 0xb8, 0xbc, 0xc8, 0xb5, 0xe4, 0x51, 0x1e, 0xd4, 0x12,
	0xbb, 0xa1, 0x35, 0xae, 0x56, 0x44, 0xcd, 0xe3, 0x89, 0x9c, 0x47, 0x9d, 0x3c, 0xcc, 0x9d, 0xc7,
	0xc8, 0xd6, 0x76, 0xf6, 0x50, 0xea, 0xfb, 0xa7, 0x03, 0x1d, 0xf5, 0x06, 0x3a, 0xfa, 0x35, 0xd0,
	0xd1, 0x87, 0xa1, 0x5e, 0xe8, 0x0d, 0xf5, 0xc2, 0xcf, 0xa1, 0x5e, 0x38, 0xa8, 0xb9, 0x5e, 0x7c,
	0xd8, 0x6d, 0x19, 0x0e, 0xf7, 0xe7, 0xa8, 0x9c, 0xcc, 0xe8, 0xc4, 0x49, 0x08, 0xa2, 0x55, 0x94,
	0x5f, 0xf6, 0xbd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0xbc, 0x36, 0xac, 0xcf, 0x06, 0x00,
	0x00,
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
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	ResourceDepedencyMappingFromMessageKey(ctx context.Context, in *ResourceDepedencyMappingFromMessageKeyRequest, opts ...grpc.CallOption) (*ResourceDepedencyMappingFromMessageKeyResponse, error)
	ListResourceDepedencyMapping(ctx context.Context, in *ListResourceDepedencyMappingRequest, opts ...grpc.CallOption) (*ListResourceDepedencyMappingResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.accesscontrol_x.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ResourceDepedencyMappingFromMessageKey(ctx context.Context, in *ResourceDepedencyMappingFromMessageKeyRequest, opts ...grpc.CallOption) (*ResourceDepedencyMappingFromMessageKeyResponse, error) {
	out := new(ResourceDepedencyMappingFromMessageKeyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.accesscontrol_x.v1beta1.Query/ResourceDepedencyMappingFromMessageKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListResourceDepedencyMapping(ctx context.Context, in *ListResourceDepedencyMappingRequest, opts ...grpc.CallOption) (*ListResourceDepedencyMappingResponse, error) {
	out := new(ListResourceDepedencyMappingResponse)
	err := c.cc.Invoke(ctx, "/cosmos.accesscontrol_x.v1beta1.Query/ListResourceDepedencyMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	ResourceDepedencyMappingFromMessageKey(context.Context, *ResourceDepedencyMappingFromMessageKeyRequest) (*ResourceDepedencyMappingFromMessageKeyResponse, error)
	ListResourceDepedencyMapping(context.Context, *ListResourceDepedencyMappingRequest) (*ListResourceDepedencyMappingResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ResourceDepedencyMappingFromMessageKey(ctx context.Context, req *ResourceDepedencyMappingFromMessageKeyRequest) (*ResourceDepedencyMappingFromMessageKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceDepedencyMappingFromMessageKey not implemented")
}
func (*UnimplementedQueryServer) ListResourceDepedencyMapping(ctx context.Context, req *ListResourceDepedencyMappingRequest) (*ListResourceDepedencyMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResourceDepedencyMapping not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.accesscontrol_x.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ResourceDepedencyMappingFromMessageKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceDepedencyMappingFromMessageKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ResourceDepedencyMappingFromMessageKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.accesscontrol_x.v1beta1.Query/ResourceDepedencyMappingFromMessageKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ResourceDepedencyMappingFromMessageKey(ctx, req.(*ResourceDepedencyMappingFromMessageKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListResourceDepedencyMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourceDepedencyMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListResourceDepedencyMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.accesscontrol_x.v1beta1.Query/ListResourceDepedencyMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListResourceDepedencyMapping(ctx, req.(*ListResourceDepedencyMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.accesscontrol_x.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ResourceDepedencyMappingFromMessageKey",
			Handler:    _Query_ResourceDepedencyMappingFromMessageKey_Handler,
		},
		{
			MethodName: "ListResourceDepedencyMapping",
			Handler:    _Query_ListResourceDepedencyMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/accesscontrol_x/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ResourceDepedencyMappingFromMessageKeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceDepedencyMappingFromMessageKeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceDepedencyMappingFromMessageKeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageKey) > 0 {
		i -= len(m.MessageKey)
		copy(dAtA[i:], m.MessageKey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MessageKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResourceDepedencyMappingFromMessageKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceDepedencyMappingFromMessageKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceDepedencyMappingFromMessageKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MessageDependencyMapping.MarshalToSizedBuffer(dAtA[:i])
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

func (m *ListResourceDepedencyMappingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResourceDepedencyMappingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListResourceDepedencyMappingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListResourceDepedencyMappingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListResourceDepedencyMappingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListResourceDepedencyMappingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageDependencyMappingList) > 0 {
		for iNdEx := len(m.MessageDependencyMappingList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MessageDependencyMappingList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *ResourceDepedencyMappingFromMessageKeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageKey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ResourceDepedencyMappingFromMessageKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MessageDependencyMapping.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *ListResourceDepedencyMappingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListResourceDepedencyMappingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MessageDependencyMappingList) > 0 {
		for _, e := range m.MessageDependencyMappingList {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ResourceDepedencyMappingFromMessageKeyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ResourceDepedencyMappingFromMessageKeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceDepedencyMappingFromMessageKeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageKey", wireType)
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
			m.MessageKey = string(dAtA[iNdEx:postIndex])
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
func (m *ResourceDepedencyMappingFromMessageKeyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ResourceDepedencyMappingFromMessageKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceDepedencyMappingFromMessageKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageDependencyMapping", wireType)
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
			if err := m.MessageDependencyMapping.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ListResourceDepedencyMappingRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListResourceDepedencyMappingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResourceDepedencyMappingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *ListResourceDepedencyMappingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListResourceDepedencyMappingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListResourceDepedencyMappingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageDependencyMappingList", wireType)
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
			m.MessageDependencyMappingList = append(m.MessageDependencyMappingList, accesscontrol.MessageDependencyMapping{})
			if err := m.MessageDependencyMappingList[len(m.MessageDependencyMappingList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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