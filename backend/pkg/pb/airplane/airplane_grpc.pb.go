// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: pb/airplane/airplane.proto

package airplane

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

// AirplaneClient is the client API for Airplane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AirplaneClient interface {
	GetAirplanesPositionStream(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (Airplane_GetAirplanesPositionStreamClient, error)
}

type airplaneClient struct {
	cc grpc.ClientConnInterface
}

func NewAirplaneClient(cc grpc.ClientConnInterface) AirplaneClient {
	return &airplaneClient{cc}
}

func (c *airplaneClient) GetAirplanesPositionStream(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (Airplane_GetAirplanesPositionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Airplane_ServiceDesc.Streams[0], "/airplane/GetAirplanesPositionStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &airplaneGetAirplanesPositionStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Airplane_GetAirplanesPositionStreamClient interface {
	Recv() (*PositionReply, error)
	grpc.ClientStream
}

type airplaneGetAirplanesPositionStreamClient struct {
	grpc.ClientStream
}

func (x *airplaneGetAirplanesPositionStreamClient) Recv() (*PositionReply, error) {
	m := new(PositionReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AirplaneServer is the server API for Airplane service.
// All implementations must embed UnimplementedAirplaneServer
// for forward compatibility
type AirplaneServer interface {
	GetAirplanesPositionStream(*PositionRequest, Airplane_GetAirplanesPositionStreamServer) error
	mustEmbedUnimplementedAirplaneServer()
}

// UnimplementedAirplaneServer must be embedded to have forward compatible implementations.
type UnimplementedAirplaneServer struct {
}

func (UnimplementedAirplaneServer) GetAirplanesPositionStream(*PositionRequest, Airplane_GetAirplanesPositionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAirplanesPositionStream not implemented")
}
func (UnimplementedAirplaneServer) mustEmbedUnimplementedAirplaneServer() {}

// UnsafeAirplaneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AirplaneServer will
// result in compilation errors.
type UnsafeAirplaneServer interface {
	mustEmbedUnimplementedAirplaneServer()
}

func RegisterAirplaneServer(s grpc.ServiceRegistrar, srv AirplaneServer) {
	s.RegisterService(&Airplane_ServiceDesc, srv)
}

func _Airplane_GetAirplanesPositionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AirplaneServer).GetAirplanesPositionStream(m, &airplaneGetAirplanesPositionStreamServer{stream})
}

type Airplane_GetAirplanesPositionStreamServer interface {
	Send(*PositionReply) error
	grpc.ServerStream
}

type airplaneGetAirplanesPositionStreamServer struct {
	grpc.ServerStream
}

func (x *airplaneGetAirplanesPositionStreamServer) Send(m *PositionReply) error {
	return x.ServerStream.SendMsg(m)
}

// Airplane_ServiceDesc is the grpc.ServiceDesc for Airplane service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Airplane_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "airplane",
	HandlerType: (*AirplaneServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAirplanesPositionStream",
			Handler:       _Airplane_GetAirplanesPositionStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/airplane/airplane.proto",
}