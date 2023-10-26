package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/config"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenterclient"
	"go-zero-bookstore/common/middleware"
)

type ServiceContext struct {
	Config               config.Config
	Usercenter           usercenterclient.Usercenter
	RecoverMiddlewareCtx *middleware.RecoverMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		Usercenter:           usercenterclient.NewUsercenter(zrpc.MustNewClient(c.Usercenter)),
		RecoverMiddlewareCtx: middleware.RecoverMiddlewareHandler,
	}
}
