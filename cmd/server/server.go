package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"webserver/common"
	"webserver/common/config"
	"webserver/database"
	ginpkg "webserver/gin"
	nacospkg "webserver/nacos"
	redisinit "webserver/redis"
)

var (
	configYml string
	BuConfig  config.BuConfig
	StartCmd  = &cobra.Command{
		Use:     "server",
		Short:   "run gin server",
		Example: "mysql server",
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

	serverport := fmt.Sprintf("0.0.0.0:%d", common.Bargconfig.Server.Port)
	ginpkg.InitGin(serverport)

	return nil
}
