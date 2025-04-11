package grpcclient

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"time"
	"webserver/common"
	"webserver/common/config"
	"webserver/database"
	grpccli "webserver/grpccli"
	nacospkg "webserver/nacos"
	pb "webserver/protobuf/protoObj"
	redisinit "webserver/redis"
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

	clients, err := grpccli.NewClients(serverport)
	common.GrpcCli = clients
	//defer clients.CloseRpc()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.HelloRequest{Name: "ChatGPT"}
	res, err := clients.MainClient.SayHello(ctx, req)
	if err != nil {
		log.Println("could not greet: %v", err)
	} else {
		fmt.Printf("Response: %s\n", res.Message)
	}
	req2 := &pb.GameRequest{Name: "ChatGPT"}
	res2, err := clients.GameClient.SayHello(ctx, req2)
	if err != nil {
		log.Println("could not greet: %v", err)
	} else {
		fmt.Printf("Response: %s\n", res2.Message)
	}
	fmt.Println("end")
	return nil
}
