package logic

import (
	"context"
	"errors"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"
	"go-zero-bookstore/app/usercenter/repository"
	"go-zero-bookstore/common/logx"
	tool "go-zero-bookstore/common/tools"
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

func (l *LoginLogic) Login(in *usercenter.LoginReq) (*usercenter.LoginResp, error) {
	logx.Info("rpc login...")
	if len(in.Mobile) == 0 {
		return nil, ErrMobileIncorrect
	}

	if len(in.Password) < 8 {
		return nil, ErrPasswordIncorrect
	}

	account, err := l.svcCtx.AccountModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrAccountNotExists
	}

	if tool.Md5ByString(in.Password) != account.Password {
		return nil, ErrAccountLoginFailed
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{UserId: account.Id})
	if err != nil {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}
	logx.Info(tokenResp)

	return &usercenter.LoginResp{Token: tokenResp.Token}, nil
}
