package user

import (
	"context"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/types"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenterclient"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	r, err := l.svcCtx.Usercenter.Register(l.ctx, &usercenterclient.RegisterReq{
		Mobile:    req.Mobile,
		Username:  req.Username,
		Password:  req.Password,
		Password2: req.Password2,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{Token: r.Token}, nil
}
