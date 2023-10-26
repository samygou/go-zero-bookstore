package bookstoreCenter

import (
	"context"
	"go-zero-bookstore/app/bookstore/cmd/rpc/bookstore"

	"go-zero-bookstore/app/bookstore/cmd/api/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBookLogic {
	return &CreateBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateBookLogic) CreateBook(req *types.CreateBookReq) (resp *types.CreateBookResp, err error) {
	r, err := l.svcCtx.Bookstore.CreateBook(l.ctx, &bookstore.CreateBookReq{
		Name:  req.Name,
		Price: req.Price,
		Desc:  req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateBookResp{
		ID: r.Id,
	}, nil
}
