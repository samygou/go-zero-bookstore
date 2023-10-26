package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"go-zero-bookstore/app/bookstore/cmd/api/internal/config"
	"go-zero-bookstore/app/bookstore/cmd/rpc/bookstore"
	"go-zero-bookstore/common/middleware"
)

type ServiceContext struct {
	Config               config.Config
	Bookstore            bookstore.Bookstore
	RecoverMiddlewareCtx *middleware.RecoverMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		Bookstore:            bookstore.NewBookstore(zrpc.MustNewClient(c.Bookstore)),
		RecoverMiddlewareCtx: middleware.RecoverMiddlewareHandler,
	}
}
