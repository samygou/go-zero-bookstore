package bookstoreCenter

import (
	"context"
	"go-zero-bookstore/app/bookstore/cmd/rpc/bookstore"

	"go-zero-bookstore/app/bookstore/cmd/api/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBookLogic {
	return &DeleteBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBookLogic) DeleteBook(req *types.DeleteBookReq) (resp *types.DeleteBookResp, err error) {
	r, err := l.svcCtx.Bookstore.DeleteBook(l.ctx, &bookstore.DeleteBookReq{Id: req.ID})
	if err != nil {
		return nil, err
	}

	return &types.DeleteBookResp{
		ID: r.Id,
	}, nil
}
