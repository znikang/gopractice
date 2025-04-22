package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type NacosServerConfig struct {
	Server struct {
		Port      uint64 `yaml:"port"`
		Host      string `yaml:"host"`
		Namespace string `yaml:"namespace"`
		Dataid    string `yaml:"dataid"`
		Group     string `yaml:"group"`
	} `yaml:"server"`
}

func LoadConfig(filename string) (*NacosServerConfig, error) {
	data, err := os.ReadFile(filename) // 读取文件
	if err != nil {
		return nil, err
	}

	var config NacosServerConfig

	err = yaml.Unmarshal(data, &config) // 解析 YAML
	if err != nil {
		return nil, err
	}

	return &config, nil
}
