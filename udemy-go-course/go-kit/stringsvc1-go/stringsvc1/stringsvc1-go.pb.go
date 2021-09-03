// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stringsvc1-go.proto

package stringsvc1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CountRequest struct {
	S                    string   `protobuf:"bytes,1,opt,name=S,proto3" json:"S,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountRequest) Reset()         { *m = CountRequest{} }
func (m *CountRequest) String() string { return proto.CompactTextString(m) }
func (*CountRequest) ProtoMessage()    {}
func (*CountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stringsvc1_go_7f753205d0171be9, []int{0}
}
func (m *CountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountRequest.Unmarshal(m, b)
}
func (m *CountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountRequest.Marshal(b, m, deterministic)
}
func (dst *CountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountRequest.Merge(dst, src)
}
func (m *CountRequest) XXX_Size() int {
	return xxx_messageInfo_CountRequest.Size(m)
}
func (m *CountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountRequest proto.InternalMessageInfo

func (m *CountRequest) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type CountResponse struct {
	V                    int32    `protobuf:"varint,1,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()         { *m = CountResponse{} }
func (m *CountResponse) String() string { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()    {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stringsvc1_go_7f753205d0171be9, []int{1}
}
func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResponse.Unmarshal(m, b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
}
func (dst *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(dst, src)
}
func (m *CountResponse) XXX_Size() int {
	return xxx_messageInfo_CountResponse.Size(m)
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetV() int32 {
	if m != nil {
		return m.V
	}
	return 0
}

type UppercaseRequest struct {
	S                    string   `protobuf:"bytes,1,opt,name=S,proto3" json:"S,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UppercaseRequest) Reset()         { *m = UppercaseRequest{} }
func (m *UppercaseRequest) String() string { return proto.CompactTextString(m) }
func (*UppercaseRequest) ProtoMessage()    {}
func (*UppercaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stringsvc1_go_7f753205d0171be9, []int{2}
}
func (m *UppercaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UppercaseRequest.Unmarshal(m, b)
}
func (m *UppercaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UppercaseRequest.Marshal(b, m, deterministic)
}
func (dst *UppercaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UppercaseRequest.Merge(dst, src)
}
func (m *UppercaseRequest) XXX_Size() int {
	return xxx_messageInfo_UppercaseRequest.Size(m)
}
func (m *UppercaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UppercaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UppercaseRequest proto.InternalMessageInfo

func (m *UppercaseRequest) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type UppercaseResponse struct {
	V                    string   `protobuf:"bytes,1,opt,name=V,proto3" json:"V,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UppercaseResponse) Reset()         { *m = UppercaseResponse{} }
func (m *UppercaseResponse) String() string { return proto.CompactTextString(m) }
func (*UppercaseResponse) ProtoMessage()    {}
func (*UppercaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stringsvc1_go_7f753205d0171be9, []int{3}
}
func (m *UppercaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UppercaseResponse.Unmarshal(m, b)
}
func (m *UppercaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UppercaseResponse.Marshal(b, m, deterministic)
}
func (dst *UppercaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UppercaseResponse.Merge(dst, src)
}
func (m *UppercaseResponse) XXX_Size() int {
	return xxx_messageInfo_UppercaseResponse.Size(m)
}
func (m *UppercaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UppercaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UppercaseResponse proto.InternalMessageInfo

func (m *UppercaseResponse) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func (m *UppercaseResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*CountRequest)(nil), "stringsvc1.CountRequest")
	proto.RegisterType((*CountResponse)(nil), "stringsvc1.CountResponse")
	proto.RegisterType((*UppercaseRequest)(nil), "stringsvc1.UppercaseRequest")
	proto.RegisterType((*UppercaseResponse)(nil), "stringsvc1.UppercaseResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StringServiceClient is the client API for StringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StringServiceClient interface {
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
	Uppercase(ctx context.Context, in *UppercaseRequest, opts ...grpc.CallOption) (*UppercaseResponse, error)
}

type stringServiceClient struct {
	cc *grpc.ClientConn
}

func NewStringServiceClient(cc *grpc.ClientConn) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/stringsvc1.StringService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) Uppercase(ctx context.Context, in *UppercaseRequest, opts ...grpc.CallOption) (*UppercaseResponse, error) {
	out := new(UppercaseResponse)
	err := c.cc.Invoke(ctx, "/stringsvc1.StringService/Uppercase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringServiceServer is the server API for StringService service.
type StringServiceServer interface {
	Count(context.Context, *CountRequest) (*CountResponse, error)
	Uppercase(context.Context, *UppercaseRequest) (*UppercaseResponse, error)
}

func RegisterStringServiceServer(s *grpc.Server, srv StringServiceServer) {
	s.RegisterService(&_StringService_serviceDesc, srv)
}

func _StringService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stringsvc1.StringService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_Uppercase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UppercaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Uppercase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stringsvc1.StringService/Uppercase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Uppercase(ctx, req.(*UppercaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stringsvc1.StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Count",
			Handler:    _StringService_Count_Handler,
		},
		{
			MethodName: "Uppercase",
			Handler:    _StringService_Uppercase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stringsvc1-go.proto",
}

func init() { proto.RegisterFile("stringsvc1-go.proto", fileDescriptor_stringsvc1_go_7f753205d0171be9) }

var fileDescriptor_stringsvc1_go_7f753205d0171be9 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0x2f, 0x2e, 0x4b, 0x36, 0xd4, 0x4d, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x42, 0x08, 0x2a, 0xc9, 0x70, 0xf1, 0x38, 0xe7, 0x97, 0xe6, 0x95, 0x04, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0xf1, 0x70, 0x31, 0x06, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31,
	0x06, 0x2b, 0xc9, 0x72, 0xf1, 0x42, 0x65, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41, 0xd2, 0x61,
	0x60, 0x69, 0xd6, 0x20, 0xc6, 0x30, 0x25, 0x05, 0x2e, 0x81, 0xd0, 0x82, 0x82, 0xd4, 0xa2, 0xe4,
	0xc4, 0xe2, 0x54, 0xec, 0x06, 0x18, 0x73, 0x09, 0x22, 0xa9, 0x40, 0x37, 0x84, 0x33, 0x88, 0x31,
	0x4c, 0x48, 0x80, 0x8b, 0xd9, 0xb5, 0xa8, 0x48, 0x82, 0x09, 0xcc, 0x07, 0x31, 0x8d, 0x66, 0x33,
	0x72, 0xf1, 0x06, 0x83, 0x9d, 0x18, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xc7, 0xc5,
	0x0a, 0x76, 0x87, 0x90, 0x84, 0x1e, 0xc2, 0xed, 0x7a, 0xc8, 0x0e, 0x97, 0x92, 0xc4, 0x22, 0x03,
	0xb1, 0x4f, 0x89, 0x41, 0xc8, 0x8b, 0x8b, 0x13, 0xee, 0x0c, 0x21, 0x19, 0x64, 0x95, 0xe8, 0xee,
	0x97, 0x92, 0xc5, 0x21, 0x0b, 0x33, 0x2b, 0x89, 0x0d, 0x1c, 0x88, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0x65, 0xe5, 0x8b, 0x5b, 0x01, 0x00, 0x00,
}