package grpcclient

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
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
		Use:     "grpcclient",
		Short:   "run grpc client test",
		Example: "webserver grpcclient",
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

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
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
	serverport := fmt.Sprintf("%s:%d", common.Bargconfig.Server.Host, common.Bargconfig.Server.Port)

	conn, err := grpc.Dial(serverport, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{Name: "ChatGPT"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Printf("Response: %s", res.Message)

	return nil
}
