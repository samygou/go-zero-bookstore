package bookstoreCenter

import (
	"context"
	"fmt"
	"go-zero-bookstore/app/bookstore/cmd/rpc/bookstore"

	"go-zero-bookstore/app/bookstore/cmd/api/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookListLogic {
	return &GetBookListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookListLogic) GetBookList(req *types.GetBookListReq) (resp *types.GetBookListResp, err error) {
	fmt.Printf("%#v\n", req)
	records, err := l.svcCtx.Bookstore.GetBookList(l.ctx, &bookstore.GetBookListReq{
		Name:     req.Name,
		Price:    req.Price,
		OrderBy:  req.OrderBy,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var books []types.Book

	for _, book := range records.Books {
		books = append(books, types.Book{
			ID:    book.Id,
			Name:  book.Name,
			Price: book.Price,
		})
	}

	return &types.GetBookListResp{Books: books}, nil
}
