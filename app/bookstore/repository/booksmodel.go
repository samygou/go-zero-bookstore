package repository

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BooksModel = (*customBooksModel)(nil)

type (
	// BooksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBooksModel.
	BooksModel interface {
		booksModel
	}

	customBooksModel struct {
		*defaultBooksModel
	}
)

// NewBooksModel returns a model for the database table.
func NewBooksModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BooksModel {
	return &customBooksModel{
		defaultBooksModel: newBooksModel(conn, c, opts...),
	}
}
