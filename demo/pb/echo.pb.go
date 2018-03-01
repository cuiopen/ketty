// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

/*
Package test_pb is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	Req
	Rsp
*/
package test_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import (
	ketty "github.com/yyzybb537/ketty"
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

type Req struct {
	Val int64 `protobuf:"varint,1,opt,name=val" json:"val,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetVal() int64 {
	if m != nil {
		return m.Val
	}
	return 0
}

type Rsp struct {
	Val int64 `protobuf:"varint,1,opt,name=val" json:"val,omitempty"`
}

func (m *Rsp) Reset()                    { *m = Rsp{} }
func (m *Rsp) String() string            { return proto.CompactTextString(m) }
func (*Rsp) ProtoMessage()               {}
func (*Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rsp) GetVal() int64 {
	if m != nil {
		return m.Val
	}
	return 0
}

func init() {
	proto.RegisterType((*Req)(nil), "test_pb.Req")
	proto.RegisterType((*Rsp)(nil), "test_pb.Rsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	Echo(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Rsp, error) {
	out := new(Rsp)
	err := grpc.Invoke(ctx, "/test_pb.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Echo(context.Context, *Req) (*Rsp, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test_pb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test_pb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}

// Reference imports to suppress errors if they are not otherwise used.
var _ ketty.Dummy

// This is a compile-time assertion to ensure that this generated file
// is compatible with the ketty package it is being compiled against.

type EchoServiceHandleT struct {
	desc *grpc.ServiceDesc
}

func (h *EchoServiceHandleT) Implement() interface{} {
	return h.desc
}

func (h *EchoServiceHandleT) ServiceName() string {
	return h.desc.ServiceName
}

var EchoServiceHandle = &EchoServiceHandleT{desc: &_EchoService_serviceDesc}

type KettyEchoServiceClient struct {
	client ketty.Client
}

func NewKettyEchoServiceClient(client ketty.Client) *KettyEchoServiceClient {
	return &KettyEchoServiceClient{client}
}

func (this *KettyEchoServiceClient) Echo(ctx context.Context, in *Req) (*Rsp, error) {
	out := new(Rsp)
	err := this.client.Invoke(ctx, EchoServiceHandle, "Echo", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x2f, 0x48, 0x52,
	0x12, 0xe7, 0x62, 0x0e, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0x2e, 0x4b, 0xcc, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x0e, 0x02, 0x31, 0xc1, 0x12, 0xc5, 0x05, 0x98, 0x12, 0x46, 0xc6, 0x5c, 0xdc,
	0xae, 0xc9, 0x19, 0xf9, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x2a, 0x5c, 0x2c, 0x20,
	0xae, 0x10, 0x8f, 0x1e, 0xd4, 0x48, 0xbd, 0xa0, 0xd4, 0x42, 0x29, 0x24, 0x5e, 0x71, 0x81, 0x12,
	0x43, 0x12, 0x1b, 0xd8, 0x5a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xa9, 0x26, 0x0c,
	0x84, 0x00, 0x00, 0x00,
}
