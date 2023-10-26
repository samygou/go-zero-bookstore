package logic

import (
	"context"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByMobileLogic {
	return &GetUserByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByMobileLogic) GetUserByMobile(in *usercenter.GetUserByMobileReq) (*usercenter.GetUserByMobileResp, error) {
	// todo: add your logic here and delete this line

	return &usercenter.GetUserByMobileResp{}, nil
}
