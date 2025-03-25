package jwtginserver

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	v22 "webserver/api/v2"
	"webserver/api/yamljwt"
	"webserver/common"
	"webserver/common/config"
	redisSource "webserver/redis"
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

func getNacosConfig(client config_client.IConfigClient, dataID, group string) (*config.BuConfig, error) {
	// 從 Nacos 獲取配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
	})
	if err != nil {
		return nil, fmt.Errorf("❌ 無法獲取 Nacos 配置: %v", err)
	}

	// 解析 YAML 配置
	var cfg config.BuConfig
	err = yaml.Unmarshal([]byte(content), &cfg)
	if err != nil {
		return nil, fmt.Errorf("❌ 解析 YAML 失敗: %v", err)
	}

	return &cfg, nil
}

func run() error {

	cfg, err := config.LoadConfig(configYml)
	if err != nil {
		log.Fatalf("❌ 加載配置失敗: %v", err)
	}
	fmt.Println("✅ 配置加載成功！")
	fmt.Printf("🌍 Nacos %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("🛢  Nacos: %s (用戶: %s %s)\n", cfg.Server.Namespace, cfg.Server.Dataid, cfg.Server.Group)

	serverConfig := []constant.ServerConfig{
		{
			IpAddr: cfg.Server.Host, // Nacos 服務器地址
			Port:   cfg.Server.Port, // Nacos 端口
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId: cfg.Server.Namespace, // 默認命名空間
		TimeoutMs:   5000,                 // 請求超時
		LogLevel:    "debug",
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		log.Fatalf("❌ 無法連接 Nacos: %v", err)
	}
	newCfg, err := getNacosConfig(client, cfg.Server.Dataid, cfg.Server.Group)
	if err != nil {
		log.Fatalf("❌ 加載配置失敗: %v", err)
	}

	fmt.Printf("🌍 伺服器運行於 %s:%d\n", newCfg.Server.Host, newCfg.Server.Port)
	fmt.Println("✅ 成功加載 Nacos 配置！")
	common.Bargconfig = *newCfg
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "mysqltest",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("🔄 檢測到配置變更，重新加載...")
			var newCfg config.BuConfig
			if err := yaml.Unmarshal([]byte(data), &newCfg); err == nil {
				common.Bargconfig = newCfg
				fmt.Println("✅ 配置已更新！")
			} else {
				fmt.Println("❌ 配置解析失敗！")
			}
		},
	})

	common.RedisCli = redisSource.InitializeRedis()

	router := gin.Default()
	serverport := fmt.Sprintf("0.0.0.0:%d", newCfg.Server.Port)
	//router.GET("test", myapi.PrintMessage).Use(myapi.PrintMessage)

	//router.GET("login", v12.Login)

	authMiddleware, err := jwt.New(yamljwt.InitParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	router.Use(yamljwt.HandlerMiddleWare(authMiddleware))

	router.POST("/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), yamljwt.HandleNoRoute())
	// 這樣就沒有保護
	router.POST("/login3", v22.Test2)

	// 有protected group 後加入問題
	protected := router.Group("/", authMiddleware.MiddlewareFunc())
	protected.GET("/refresh_token", authMiddleware.RefreshHandler)
	protected.POST("login4", v22.Test2)
	//做了保護
	auth := router.Group("/auth", authMiddleware.MiddlewareFunc())
	auth.GET("hello", yamljwt.HelloHandler)
	auth.GET("hello2", v22.Test2)

	// 沒作保護
	v2 := router.Group("apis")
	v2.GET("test", v22.Test2)

	router.Run(serverport)

	return nil
}
