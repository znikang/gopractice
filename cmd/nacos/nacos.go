package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"yaml/common/config"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "nacos",
		Short:   "run gin server",
		Example: "mysql nacos",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func getConfig(client config_client.IConfigClient, dataID, group string) (*config.ServerConfig, error) {
	// 從 Nacos 獲取配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
	})
	if err != nil {
		return nil, fmt.Errorf("❌ 無法獲取 Nacos 配置: %v", err)
	}

	// 解析 YAML 配置
	var cfg config.ServerConfig
	err = yaml.Unmarshal([]byte(content), &cfg)
	if err != nil {
		return nil, fmt.Errorf("❌ 解析 YAML 失敗: %v", err)
	}

	return &cfg, nil
}

func run() error {
	fmt.Println("nacos ")
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: "192.168.1.15", // Nacos 服務器地址
			Port:   8848,           // Nacos 端口
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId: "bbbb", // 默認命名空間
		TimeoutMs:   5000,   // 請求超時
		LogLevel:    "debug",
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		log.Fatalf("❌ 無法連接 Nacos: %v", err)
	}

	// **3. 讀取 Nacos 配置**
	cfg, err := getConfig(client, "mysqltest", "DEFAULT_GROUP")
	if err != nil {
		log.Fatalf("❌ 加載配置失敗: %v", err)
	}
	fmt.Println("✅ 成功加載 Nacos 配置！")
	fmt.Printf("🌍 伺服器運行於 %s:%d\n", cfg.Server.Host, cfg.Server.Port)

	err = client.ListenConfig(vo.ConfigParam{
		DataId: "mysqltest",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("🔄 檢測到配置變更，重新加載...")

			var newCfg config.ServerConfig
			if err := yaml.Unmarshal([]byte(data), &newCfg); err == nil {
				cfg = &newCfg
				fmt.Println("✅ 配置已更新！")
			} else {
				fmt.Println("❌ 配置解析失敗！")
			}
		},
	})
	if err != nil {
		log.Fatalf("❌ 監聽 Nacos 配置變更失敗: %v", err)
	}

	// **保持運行**

	return nil
}
