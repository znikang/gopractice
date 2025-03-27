package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v2"
	"log"
	"webserver/api/models/login"
	"webserver/common"
	"webserver/common/config"
)

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

func InitNacos(cfg *config.ServerConfig) {

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
	common.Bargconfig = *newCfg
	login.JwtSecret = newCfg.Server.Secretkey
	login.RefshToeknSecret = newCfg.Server.Refeshkey

	err = client.ListenConfig(vo.ConfigParam{
		DataId: cfg.Server.Dataid,
		Group:  cfg.Server.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("🔄 檢測到配置變更，重新加載...")
			var newCfg config.BuConfig
			if err := yaml.Unmarshal([]byte(data), &newCfg); err == nil {
				common.Bargconfig = newCfg
				login.JwtSecret = newCfg.Server.Secretkey
				login.RefshToeknSecret = newCfg.Server.Refeshkey
				fmt.Println("✅ 配置已更新！")
			} else {
				fmt.Println("❌ 配置解析失敗！")
			}
		},
	})

}
