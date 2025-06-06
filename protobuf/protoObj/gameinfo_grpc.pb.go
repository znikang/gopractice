// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: gameinfo.proto

package protoObj

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
	GameService_SayHello_FullMethodName = "/RpcStruct.GameService/SayHello"
)

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	SayHello(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameReply, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) SayHello(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameReply, error) {
	out := new(GameReply)
	err := c.cc.Invoke(ctx, GameService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	SayHello(context.Context, *GameRequest) (*GameReply, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) SayHello(context.Context, *GameRequest) (*GameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).SayHello(ctx, req.(*GameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RpcStruct.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GameService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gameinfo.proto",
}
