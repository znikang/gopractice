package authserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"time"
	"yaml/api/models/login"
	"yaml/common"
	"yaml/common/config"
	yamlredis "yaml/redis"
)

var (
	configYml string
	BuConfig  config.BuConfig
	StartCmd  = &cobra.Command{
		Use:     "auth",
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
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/authsettings.yml", "Start server with provided configuration file")
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

	common.RedisCli = yamlredis.InitializeRedis()

	router := gin.Default()
	serverport := fmt.Sprintf("0.0.0.0:%d", newCfg.Server.Port)

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	router.POST("login", loginHandler)
	// 沒作保護

	router.Run(serverport)

	return nil
}

func generateToken(username string) (string, int64, error) {

	expirationTime := time.Now().Add(1 * time.Hour).Unix() // Token 過期時間 (1小時)
	claims := &login.Claims{
		Username: username,
		RegisteredClaims: jwt5.RegisteredClaims{
			ExpiresAt: jwt5.NewNumericDate(time.Unix(expirationTime, 0)),
			IssuedAt:  jwt5.NewNumericDate(time.Now()),
		},
	}

	// 簽發 Token
	token := jwt5.NewWithClaims(jwt5.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(login.JwtSecret)
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

func loginHandler(c *gin.Context) {

	var loginRequest = login.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if loginRequest.Username != "admin" || loginRequest.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, expiresAt, err := generateToken(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx := context.Background()
	err = common.RedisCli.Set(ctx, "token:"+loginRequest.Username, token, time.Hour).Err()
	if err != nil {
		fmt.Println("Redis 記錄 Token 失敗:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"expires_at": expiresAt,
	})
}
