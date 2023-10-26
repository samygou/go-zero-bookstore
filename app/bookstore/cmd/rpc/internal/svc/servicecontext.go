package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/config"
	"go-zero-bookstore/app/bookstore/repository"
)

type ServiceContext struct {
	Config    config.Config
	BookModel repository.BooksModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		BookModel: repository.NewBooksModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
