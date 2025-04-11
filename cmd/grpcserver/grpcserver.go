package grpcserver

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
	"webserver/common"
	"webserver/common/config"
	"webserver/database"
	nacospkg "webserver/nacos"
	pb "webserver/protobuf/protoObj"
	redisinit "webserver/redis"

	"context"
)

var (
	configYml string
	BuConfig  config.BuConfig
	StartCmd  = &cobra.Command{
		Use:     "grpcserver",
		Short:   "run grpc server",
		Example: "webserver grpcserver",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func initTools() {
	common.RedisCli = redisinit.InitializeRedis()
	common.OrmCli = database.InitializeDatabases()
}

type mainServer struct {
	pb.UnimplementedMainServiceServer
}
type gameServer struct {
	pb.UnimplementedGameServiceServer
}

//
//type Server struct {
//	greeterpb.UnimplementedGreeterServer
//	mathpb.UnimplementedCalculatorServer
//}

func (s *mainServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {

	//...

	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func (s *gameServer) SayHello(ctx context.Context, req *pb.GameRequest) (*pb.GameReply, error) {

	//...

	return &pb.GameReply{Message: "Hello " + req.Name}, nil
}

func run() error {

	cfg, err := config.LoadConfig(configYml)
	if err != nil {
		log.Fatalf("❌ 加載配置失敗: %v", err)
	}
	fmt.Println("✅ 配置加載成功！")
	fmt.Printf("🌍 Nacos %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("🛢  Nacos: %s (用戶: %s %s)\n", cfg.Server.Namespace, cfg.Server.Dataid, cfg.Server.Group)

	nacospkg.InitNacos(cfg)
	initTools()
	serverport := fmt.Sprintf(":%d", common.Bargconfig.Server.Port)

	lis, err := net.Listen("tcp", serverport)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMainServiceServer(grpcServer, &mainServer{})
	pb.RegisterGameServiceServer(grpcServer, &gameServer{})

	log.Println("gRPC server listening on port {}", serverport)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}
