package config

type BuConfig struct {
	Server struct {
		Port       uint64 `yaml:"port"`
		Host       string `yaml:"host"`
		Secrectkey string `yaml:"secretkey"`
	} `yaml:"server"`

	Redis struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
}
