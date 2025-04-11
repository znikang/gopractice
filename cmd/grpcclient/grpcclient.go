package grpcclient

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
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
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/rpcsettings.yml", "Start server with provided configuration file")
}

func initTools() {
	common.RedisCli = redisinit.InitializeRedis()
	common.OrmCli = database.InitializeDatabases()
}

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

type mainServer struct {
	pb.UnimplementedMainServiceServer
}

func (s *mainServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
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
	serverport := fmt.Sprintf("%s:%d", common.Bargconfig.RpcConnect.Host, common.Bargconfig.RpcConnect.Port)

	clients, err := NewClients(serverport)
	defer clients.CloseRpc()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.HelloRequest{Name: "ChatGPT"}
	res, err := clients.MainClient.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Printf("Response: %s\n", res.Message)
	fmt.Println("end")
	return nil
}
