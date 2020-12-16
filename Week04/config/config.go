package config

import "fmt"

type Config struct {
}

func NewUserConfig(conf *Config) {
	fmt.Println("初始化资源")
}
