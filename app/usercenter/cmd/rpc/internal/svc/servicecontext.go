package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/config"
	"go-zero-bookstore/app/usercenter/repository"
)

type ServiceContext struct {
	Config       config.Config
	AccountModel repository.AccountsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AccountModel: repository.NewAccountsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
