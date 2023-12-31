// Code generated by goctl. DO NOT EDIT.

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"

	"go-zero-bookstore/common/logx"
)

var (
	booksFieldNames          = builder.RawFieldNames(&Books{})
	booksRows                = strings.Join(booksFieldNames, ",")
	booksRowsExpectAutoSet   = strings.Join(stringx.Remove(booksFieldNames, "`id`"), ",")
	booksRowsWithPlaceHolder = strings.Join(stringx.Remove(booksFieldNames, "`id`", "`create_time`", "`create_at`"), "=?,") + "=?"

	cacheBooksIdPrefix = "cache:books:id:"
)

type GetBookListReq struct {
	Name     *string `json:"name"`
	Price    *int64  `json:"price"`
	OrderBy  *string `json:"order_by"`
	Page     int64   `json:"page"`
	PageSize int64   `json:"page_size"`
}

type (
	booksModel interface {
		Insert(ctx context.Context, data *Books) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Books, error)
		Update(ctx context.Context, data *Books) error
		Delete(ctx context.Context, id int64) error
		GetBookList(ctx context.Context, req *GetBookListReq) ([]*Books, error)
	}

	defaultBooksModel struct {
		sqlc.CachedConn
		table string
	}

	Books struct {
		Id       int64  `db:"id"`        // 主键id
		CreateAt int64  `db:"create_at"` // 创建时间
		UpdateAt int64  `db:"update_at"` // 更新时间
		Name     string `db:"name"`      // 书名
		Price    int64  `db:"price"`     // 价格
		Desc     string `db:"desc"`      // 描述
	}
)

func newBooksModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultBooksModel {
	return &defaultBooksModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`books`",
	}
}

func (m *defaultBooksModel) withSession(session sqlx.Session) *defaultBooksModel {
	return &defaultBooksModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`books`",
	}
}

func (m *defaultBooksModel) Delete(ctx context.Context, id int64) error {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, booksIdKey)
	return err
}

func (m *defaultBooksModel) FindOne(ctx context.Context, id int64) (*Books, error) {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, id)
	var resp Books
	err := m.QueryRowCtx(ctx, &resp, booksIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", booksRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBooksModel) Insert(ctx context.Context, data *Books) (sql.Result, error) {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, booksRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreateAt, data.UpdateAt, data.Name, data.Price, data.Desc)
	}, booksIdKey)
	return ret, err
}

func (m *defaultBooksModel) Update(ctx context.Context, data *Books) error {
	booksIdKey := fmt.Sprintf("%s%v", cacheBooksIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, booksRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UpdateAt, data.Name, data.Price, data.Desc, data.Id)
	}, booksIdKey)
	return err
}

func (m *defaultBooksModel) GetBookList(ctx context.Context, req *GetBookListReq) ([]*Books, error) {
	whereBuilder := m.selectBuilder()

	whereBuilder = whereBuilder.Columns(booksRows)

	if req.Name != nil {
		whereBuilder = whereBuilder.Where(squirrel.Eq{"name": *req.Name})
	}

	if req.Price != nil {
		whereBuilder = whereBuilder.Where(squirrel.Eq{"price": *req.Price})
	}

	if req.OrderBy != nil {
		whereBuilder = whereBuilder.OrderBy(*req.OrderBy)
	}

	offset := (req.Page - 1) * req.PageSize

	query, values, err := whereBuilder.Offset(uint64(offset)).Limit(uint64(req.PageSize)).ToSql()
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	var books []*Books
	err = m.QueryRowsNoCacheCtx(ctx, &books, query, values...)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return books, nil
}

func (m *defaultBooksModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheBooksIdPrefix, primary)
}

func (m *defaultBooksModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", booksRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBooksModel) tableName() string {
	return m.table
}

func (m *defaultBooksModel) selectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
