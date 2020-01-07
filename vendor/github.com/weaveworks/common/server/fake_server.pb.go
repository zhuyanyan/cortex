// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fake_server.proto

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	fake_server.proto

It has these top-level messages:
	FailWithHTTPErrorRequest
*/
package server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type FailWithHTTPErrorRequest struct {
	Code int32 `protobuf:"varint,1,opt,name=Code" json:"Code,omitempty"`
}

func (m *FailWithHTTPErrorRequest) Reset()                    { *m = FailWithHTTPErrorRequest{} }
func (m *FailWithHTTPErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*FailWithHTTPErrorRequest) ProtoMessage()               {}
func (*FailWithHTTPErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FailWithHTTPErrorRequest) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*FailWithHTTPErrorRequest)(nil), "server.FailWithHTTPErrorRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FakeServer service

type FakeServerClient interface {
	Succeed(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	FailWithError(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	FailWithHTTPError(ctx context.Context, in *FailWithHTTPErrorRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type fakeServerClient struct {
	cc *grpc.ClientConn
}

func NewFakeServerClient(cc *grpc.ClientConn) FakeServerClient {
	return &fakeServerClient{cc}
}

func (c *fakeServerClient) Succeed(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/server.FakeServer/Succeed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakeServerClient) FailWithError(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/server.FakeServer/FailWithError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakeServerClient) FailWithHTTPError(ctx context.Context, in *FailWithHTTPErrorRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/server.FakeServer/FailWithHTTPError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FakeServer service

type FakeServerServer interface {
	Succeed(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
	FailWithError(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
	FailWithHTTPError(context.Context, *FailWithHTTPErrorRequest) (*google_protobuf.Empty, error)
}

func RegisterFakeServerServer(s *grpc.Server, srv FakeServerServer) {
	s.RegisterService(&_FakeServer_serviceDesc, srv)
}

func _FakeServer_Succeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeServerServer).Succeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.FakeServer/Succeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeServerServer).Succeed(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakeServer_FailWithError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeServerServer).FailWithError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.FakeServer/FailWithError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeServerServer).FailWithError(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakeServer_FailWithHTTPError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailWithHTTPErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeServerServer).FailWithHTTPError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.FakeServer/FailWithHTTPError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeServerServer).FailWithHTTPError(ctx, req.(*FailWithHTTPErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FakeServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.FakeServer",
	HandlerType: (*FakeServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Succeed",
			Handler:    _FakeServer_Succeed_Handler,
		},
		{
			MethodName: "FailWithError",
			Handler:    _FakeServer_FailWithError_Handler,
		},
		{
			MethodName: "FailWithHTTPError",
			Handler:    _FakeServer_FailWithHTTPError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fake_server.proto",
}

func init() { proto.RegisterFile("fake_server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4b, 0xcc, 0x4e,
	0x8d, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83,
	0xf0, 0xa4, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xa2, 0x49, 0xa5, 0x69, 0xfa,
	0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10, 0x45, 0x4a, 0x7a, 0x5c, 0x12, 0x6e, 0x89, 0x99, 0x39, 0xe1,
	0x99, 0x25, 0x19, 0x1e, 0x21, 0x21, 0x01, 0xae, 0x45, 0x45, 0xf9, 0x45, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0xce, 0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xac, 0x41, 0x60, 0xb6, 0xd1, 0x5d, 0x46, 0x2e, 0x2e, 0xb7, 0xc4, 0xec, 0xd4, 0x60, 0xb0, 0xd9,
	0x42, 0xd6, 0x5c, 0xec, 0xc1, 0xa5, 0xc9, 0xc9, 0xa9, 0xa9, 0x29, 0x42, 0x62, 0x7a, 0x10, 0x7b,
	0xf4, 0x60, 0xf6, 0xe8, 0xb9, 0x82, 0xec, 0x91, 0xc2, 0x21, 0xae, 0xc4, 0x20, 0xe4, 0xc8, 0xc5,
	0x0b, 0xb3, 0x1b, 0x6c, 0x2f, 0x19, 0x46, 0xf8, 0x73, 0x09, 0x62, 0x38, 0x5f, 0x48, 0x41, 0x0f,
	0x1a, 0x0e, 0xb8, 0x7c, 0x86, 0xdb, 0x40, 0x27, 0xd5, 0x28, 0xe5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2,
	0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf2, 0xd4, 0xc4, 0xb2, 0xd4, 0xf2, 0xfc, 0xa2, 0xec, 0x62,
	0xfd, 0xe4, 0xfc, 0xdc, 0xdc, 0xfc, 0x3c, 0x7d, 0x88, 0xc9, 0x49, 0x6c, 0x60, 0x8d, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x74, 0xb3, 0xd0, 0x77, 0x01, 0x00, 0x00,
}