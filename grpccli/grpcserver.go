package grpccli

import (
	"context"
	"google.golang.org/grpc"
	pb "webserver/protobuf/protoObj"
)

type mainServer struct {
	pb.UnimplementedMainServiceServer
}
type gameServer struct {
	pb.UnimplementedGameServiceServer
}

//type Server struct {
//	pb.UnimplementedMainServiceServer
//	pb.UnimplementedGameServiceServer
//}

func (s *mainServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {

	//...
	return &pb.HelloReply{Message: "Main Hello " + req.Name}, nil
}

func (s *gameServer) SayHello(ctx context.Context, req *pb.GameRequest) (*pb.GameReply, error) {

	//...
	return &pb.GameReply{Message: "Game Hello " + req.Name}, nil
}

func InitGrpc(grpcServer grpc.ServiceRegistrar) {
	pb.RegisterMainServiceServer(grpcServer, &mainServer{})
	pb.RegisterGameServiceServer(grpcServer, &gameServer{})
}
