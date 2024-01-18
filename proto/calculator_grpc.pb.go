// calculator.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/calculator.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Calculator_Add_FullMethodName       = "/calculator.Calculator/Add"
	Calculator_StreamAdd_FullMethodName = "/calculator.Calculator/StreamAdd"
	Calculator_AddStream_FullMethodName = "/calculator.Calculator/AddStream"
)

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	StreamAdd(ctx context.Context, opts ...grpc.CallOption) (Calculator_StreamAddClient, error)
	AddStream(ctx context.Context, in *NumList, opts ...grpc.CallOption) (Calculator_AddStreamClient, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, Calculator_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) StreamAdd(ctx context.Context, opts ...grpc.CallOption) (Calculator_StreamAddClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[0], Calculator_StreamAdd_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorStreamAddClient{stream}
	return x, nil
}

type Calculator_StreamAddClient interface {
	Send(*StreamNumList) error
	CloseAndRecv() (*AddResponse, error)
	grpc.ClientStream
}

type calculatorStreamAddClient struct {
	grpc.ClientStream
}

func (x *calculatorStreamAddClient) Send(m *StreamNumList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorStreamAddClient) CloseAndRecv() (*AddResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) AddStream(ctx context.Context, in *NumList, opts ...grpc.CallOption) (Calculator_AddStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[1], Calculator_AddStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorAddStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calculator_AddStreamClient interface {
	Recv() (*StreamAddResponse, error)
	grpc.ClientStream
}

type calculatorAddStreamClient struct {
	grpc.ClientStream
}

func (x *calculatorAddStreamClient) Recv() (*StreamAddResponse, error) {
	m := new(StreamAddResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility
type CalculatorServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	StreamAdd(Calculator_StreamAddServer) error
	AddStream(*NumList, Calculator_AddStreamServer) error
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (UnimplementedCalculatorServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServer) StreamAdd(Calculator_StreamAddServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAdd not implemented")
}
func (UnimplementedCalculatorServer) AddStream(*NumList, Calculator_AddStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AddStream not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_StreamAdd_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).StreamAdd(&calculatorStreamAddServer{stream})
}

type Calculator_StreamAddServer interface {
	SendAndClose(*AddResponse) error
	Recv() (*StreamNumList, error)
	grpc.ServerStream
}

type calculatorStreamAddServer struct {
	grpc.ServerStream
}

func (x *calculatorStreamAddServer) SendAndClose(m *AddResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorStreamAddServer) Recv() (*StreamNumList, error) {
	m := new(StreamNumList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_AddStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NumList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).AddStream(m, &calculatorAddStreamServer{stream})
}

type Calculator_AddStreamServer interface {
	Send(*StreamAddResponse) error
	grpc.ServerStream
}

type calculatorAddStreamServer struct {
	grpc.ServerStream
}

func (x *calculatorAddStreamServer) Send(m *StreamAddResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAdd",
			Handler:       _Calculator_StreamAdd_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AddStream",
			Handler:       _Calculator_AddStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/calculator.proto",
}