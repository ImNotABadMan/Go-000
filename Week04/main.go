package main

import (
	"Go-000/Week04/server"
	"fmt"
)

func main() {
	fmt.Println("1. 按照自己的构想，写一个项目满足基本的目录结构和工程，",
		"代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，",
		"使用 Wire 构建依赖。可以使用自己熟悉的框架")

}

func init() {
	fmt.Println("使用wire初始化server等资源")
	fmt.Println(" wire.Build(Server, Signal)")
	server.StartServer()

}
