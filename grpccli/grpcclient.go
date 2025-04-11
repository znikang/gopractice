package grpccli

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "webserver/protobuf/protoObj"
)

type Clients struct {
	GameClient pb.GameServiceClient
	MainClient pb.MainServiceClient
	conn       *grpc.ClientConn
}

func NewClients(grpcAddr string) (*Clients, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, grpcAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	return &Clients{
		GameClient: pb.NewGameServiceClient(conn),
		MainClient: pb.NewMainServiceClient(conn),
		conn:       conn,
	}, nil
}

func (c *Clients) CloseRpc() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			log.Printf("Error closing gRPC conn: %v", err)
		}
	}
}
