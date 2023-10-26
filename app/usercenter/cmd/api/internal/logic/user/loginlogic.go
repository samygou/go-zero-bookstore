package user

import (
	"context"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenterclient"
	"go-zero-bookstore/common/logx"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/types"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	r, err := l.svcCtx.Usercenter.Login(l.ctx, &usercenterclient.LoginReq{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &types.LoginResp{
		Token: r.Token,
	}, nil
}
