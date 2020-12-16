package server

import (
	"Go-000/Week04/config"
	"fmt"
)

func StartServer() {
	fmt.Println("开启server")

	config.NewUserConfig(new(config.Config))
}
