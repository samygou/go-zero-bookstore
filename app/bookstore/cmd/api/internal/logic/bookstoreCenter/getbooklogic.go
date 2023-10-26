package bookstoreCenter

import (
	"context"
	"go-zero-bookstore/app/bookstore/cmd/rpc/bookstore"

	"go-zero-bookstore/app/bookstore/cmd/api/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookLogic {
	return &GetBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookLogic) GetBook(req *types.GetBookReq) (resp *types.GetBookResp, err error) {
	record, err := l.svcCtx.Bookstore.GetBook(l.ctx, &bookstore.GetBookReq{
		Id: req.ID,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetBookResp{
		ID:    record.Book.Id,
		Name:  record.Book.Name,
		Price: record.Book.Price,
	}, nil
}
