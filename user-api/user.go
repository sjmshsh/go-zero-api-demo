package main

import (
	config2 "Go-Zero/user-api/internal/config"
	handler2 "Go-Zero/user-api/internal/handler"
	svc2 "Go-Zero/user-api/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// 这个是获取配置文件的路径
var configFile = flag.String("f", "user-api/etc/user-api.yaml", "the config file")

func main() {
	// 这个是开始解析
	flag.Parse()

	// 对你的yaml文件里面的配置内容就性映射
	var c config2.Config
	conf.MustLoad(*configFile, &c)

	// 启动服务，当结束的时候把服务关掉
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 初始化了一个ServiceContext
	ctx := svc2.NewServiceContext(c)
	// 这个是注册路由
	handler2.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
