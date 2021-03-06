// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.4
// source: api/reporting.proto

package api

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

// ReportingClient is the client API for Reporting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportingClient interface {
	GenerateStatement(ctx context.Context, opts ...grpc.CallOption) (Reporting_GenerateStatementClient, error)
}

type reportingClient struct {
	cc grpc.ClientConnInterface
}

func NewReportingClient(cc grpc.ClientConnInterface) ReportingClient {
	return &reportingClient{cc}
}

func (c *reportingClient) GenerateStatement(ctx context.Context, opts ...grpc.CallOption) (Reporting_GenerateStatementClient, error) {
	stream, err := c.cc.NewStream(ctx, &Reporting_ServiceDesc.Streams[0], "/protobuf.Reporting/GenerateStatement", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportingGenerateStatementClient{stream}
	return x, nil
}

type Reporting_GenerateStatementClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type reportingGenerateStatementClient struct {
	grpc.ClientStream
}

func (x *reportingGenerateStatementClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reportingGenerateStatementClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReportingServer is the server API for Reporting service.
// All implementations must embed UnimplementedReportingServer
// for forward compatibility
type ReportingServer interface {
	GenerateStatement(Reporting_GenerateStatementServer) error
	mustEmbedUnimplementedReportingServer()
}

// UnimplementedReportingServer must be embedded to have forward compatible implementations.
type UnimplementedReportingServer struct {
}

func (UnimplementedReportingServer) GenerateStatement(Reporting_GenerateStatementServer) error {
	return status.Errorf(codes.Unimplemented, "method GenerateStatement not implemented")
}
func (UnimplementedReportingServer) mustEmbedUnimplementedReportingServer() {}

// UnsafeReportingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportingServer will
// result in compilation errors.
type UnsafeReportingServer interface {
	mustEmbedUnimplementedReportingServer()
}

func RegisterReportingServer(s grpc.ServiceRegistrar, srv ReportingServer) {
	s.RegisterService(&Reporting_ServiceDesc, srv)
}

func _Reporting_GenerateStatement_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReportingServer).GenerateStatement(&reportingGenerateStatementServer{stream})
}

type Reporting_GenerateStatementServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type reportingGenerateStatementServer struct {
	grpc.ServerStream
}

func (x *reportingGenerateStatementServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reportingGenerateStatementServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Reporting_ServiceDesc is the grpc.ServiceDesc for Reporting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reporting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Reporting",
	HandlerType: (*ReportingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateStatement",
			Handler:       _Reporting_GenerateStatement_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/reporting.proto",
}
