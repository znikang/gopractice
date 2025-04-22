package config

type Database struct {
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type BuConfig struct {
	Server struct {
		Port      uint64 `yaml:"port"`
		Host      string `yaml:"host"`
		Secretkey string `yaml:"secretkey"`
		Refeshkey string `yaml:"refeshkey"`
	} `yaml:"server"`

	Redis struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	Database Database `yaml:"mysql"`

	Databases []Database `yaml:"mysqls"`

	RpcConnect struct {
		Port uint64 `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"grpc"`
}
