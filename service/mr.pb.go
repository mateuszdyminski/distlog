// Code generated by protoc-gen-go.
// source: mr.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	mr.proto

It has these top-level messages:
	Input
	Result
*/
package service

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

type Operation int32

const (
	Operation_ADD    Operation = 0
	Operation_SUB    Operation = 1
	Operation_DIVIDE Operation = 2
	Operation_MULTI  Operation = 3
)

var Operation_name = map[int32]string{
	0: "ADD",
	1: "SUB",
	2: "DIVIDE",
	3: "MULTI",
}
var Operation_value = map[string]int32{
	"ADD":    0,
	"SUB":    1,
	"DIVIDE": 2,
	"MULTI":  3,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Input struct {
	Arg1      int32     `protobuf:"varint,1,opt,name=arg1" json:"arg1,omitempty"`
	Arg2      int32     `protobuf:"varint,2,opt,name=arg2" json:"arg2,omitempty"`
	Operation Operation `protobuf:"varint,3,opt,name=operation,enum=service.Operation" json:"operation,omitempty"`
}

func (m *Input) Reset()                    { *m = Input{} }
func (m *Input) String() string            { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()               {}
func (*Input) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Input) GetArg1() int32 {
	if m != nil {
		return m.Arg1
	}
	return 0
}

func (m *Input) GetArg2() int32 {
	if m != nil {
		return m.Arg2
	}
	return 0
}

func (m *Input) GetOperation() Operation {
	if m != nil {
		return m.Operation
	}
	return Operation_ADD
}

type Result struct {
	Result int32  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Input)(nil), "service.Input")
	proto.RegisterType((*Result)(nil), "service.Result")
	proto.RegisterEnum("service.Operation", Operation_name, Operation_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calculator service

type CalculatorClient interface {
	Calculate(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Result, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Calculate(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/service.Calculator/Calculate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calculator service

type CalculatorServer interface {
	Calculate(context.Context, *Input) (*Result, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Calculator/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Calculate(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Calculator_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mr.proto",
}

func init() { proto.RegisterFile("mr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x14, 0x84, 0xbb, 0x5d, 0x37, 0x35, 0xef, 0x50, 0xc3, 0x43, 0x64, 0xf1, 0x54, 0xf6, 0x54, 0x3c,
	0x04, 0x8d, 0xe2, 0xc9, 0x8b, 0x1a, 0x0f, 0x01, 0x45, 0x88, 0xd6, 0x7b, 0x2c, 0x41, 0x0a, 0x6b,
	0xb3, 0xbc, 0x66, 0xfd, 0xfd, 0x62, 0x9a, 0xc6, 0xdb, 0x7c, 0x93, 0x30, 0x33, 0x09, 0x1c, 0x7f,
	0x93, 0x1c, 0x28, 0xc4, 0x80, 0xb3, 0x9d, 0xa7, 0x9f, 0xcd, 0xda, 0x77, 0x0e, 0x1a, 0xb3, 0x1d,
	0xc6, 0x88, 0x08, 0x47, 0x8e, 0xbe, 0xae, 0xda, 0x6a, 0x51, 0x2d, 0x1b, 0x9b, 0x74, 0xf6, 0x54,
	0x3b, 0x2d, 0x9e, 0xc2, 0x4b, 0xe0, 0x61, 0xf0, 0xe4, 0xe2, 0x26, 0x6c, 0xdb, 0x7a, 0x51, 0x2d,
	0xe7, 0x0a, 0x65, 0x4e, 0x93, 0xaf, 0x87, 0x13, 0xfb, 0x7f, 0xa9, 0xbb, 0x05, 0x66, 0xfd, 0x6e,
	0xec, 0x23, 0x9e, 0x01, 0xa3, 0xa4, 0x72, 0x4b, 0x26, 0x3c, 0x85, 0xc6, 0x13, 0x05, 0x4a, 0x45,
	0xdc, 0xee, 0xe1, 0xe2, 0x06, 0x78, 0xc9, 0xc3, 0x19, 0xd4, 0xf7, 0x5a, 0x8b, 0xc9, 0x9f, 0x78,
	0x5b, 0x3d, 0x88, 0x0a, 0x01, 0x98, 0x36, 0x1f, 0x46, 0x3f, 0x89, 0x29, 0x72, 0x68, 0x5e, 0x56,
	0xcf, 0xef, 0x46, 0xd4, 0xea, 0x0e, 0xe0, 0xd1, 0xf5, 0xeb, 0xb1, 0x77, 0x31, 0x10, 0x4a, 0xe0,
	0x07, 0xf2, 0x38, 0x2f, 0x3b, 0xd3, 0x93, 0xcf, 0x4f, 0x0a, 0xef, 0xf7, 0x75, 0x93, 0x4f, 0x96,
	0xbe, 0xe7, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x39, 0xb4, 0xc5, 0x2a, 0x01, 0x00, 0x00,
}
