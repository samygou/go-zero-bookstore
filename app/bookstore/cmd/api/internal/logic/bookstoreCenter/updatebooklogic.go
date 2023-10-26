package bookstoreCenter

import (
	"context"
	"go-zero-bookstore/app/bookstore/cmd/rpc/bookstore"

	"go-zero-bookstore/app/bookstore/cmd/api/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBookLogic {
	return &UpdateBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBookLogic) UpdateBook(req *types.UpdateBookReq) (resp *types.UpdateBookResp, err error) {
	r, err := l.svcCtx.Bookstore.UpdateBook(l.ctx, &bookstore.UpdateBookReq{
		Id:    req.ID,
		Name:  req.Name,
		Price: req.Price,
		Desc:  req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateBookResp{
		ID: r.Id,
	}, nil
}
