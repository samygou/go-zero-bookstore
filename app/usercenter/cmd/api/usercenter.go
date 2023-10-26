package main

import (
	"flag"
	"fmt"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/config"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/handler"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/svc"
	"go-zero-bookstore/common/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 注册一个全局错误拦截器, 用于处理未知错误
	middleware.RecoverMiddlewareHandler = middleware.NewRecoverMiddleware()
	server.Use(middleware.RecoverMiddlewareHandler.Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
