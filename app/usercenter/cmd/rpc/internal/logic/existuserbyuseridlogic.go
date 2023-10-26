package logic

import (
	"context"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExistUserByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistUserByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistUserByUserIdLogic {
	return &ExistUserByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExistUserByUserIdLogic) ExistUserByUserId(in *usercenter.ExistUserByUserIdReq) (*usercenter.ExistUserByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &usercenter.ExistUserByUserIdResp{}, nil
}
