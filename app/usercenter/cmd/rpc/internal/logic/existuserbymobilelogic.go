package logic

import (
	"context"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExistUserByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistUserByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistUserByMobileLogic {
	return &ExistUserByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExistUserByMobileLogic) ExistUserByMobile(in *usercenter.ExistUserByMobileReq) (*usercenter.ExistUserByMobileResp, error) {
	// todo: add your logic here and delete this line

	return &usercenter.ExistUserByMobileResp{}, nil
}
