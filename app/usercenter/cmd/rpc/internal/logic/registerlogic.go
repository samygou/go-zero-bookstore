package logic

import (
	"context"
	"errors"
	"time"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"
	"go-zero-bookstore/app/usercenter/repository"
	"go-zero-bookstore/common/logx"
	tool "go-zero-bookstore/common/tools"
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

func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.RegisterResp, error) {
	if len(in.Username) == 0 {
		return nil, ErrUsernameIsEmpty
	}

	if len(in.Password) == 0 {
		return nil, ErrPasswordIncorrect
	}

	if in.Password != in.Password2 {
		return nil, ErrPasswordUnEqual
	}

	account, err := l.svcCtx.AccountModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}
	logx.Info(account)
	if account != nil {
		logx.Info("账户已经存在")
		return nil, ErrAccountAlreadyExists
	}

	now := time.Now().Unix()

	result, err := l.svcCtx.AccountModel.Insert(l.ctx, &repository.Accounts{
		CreateTime: now,
		UpdateTime: now,
		DelStatus:  0,
		Mobile:     in.Mobile,
		Username:   in.Username,
		Password:   tool.Md5ByString(in.Password),
		Sex:        0,
		Avatar:     "",
		Remark:     "",
	})
	if err != nil {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}
	accountId, err := result.LastInsertId()
	if err != nil {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}

	//Generate token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{UserId: accountId})
	if err != nil {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}

	return &usercenter.RegisterResp{Token: tokenResp.Token}, nil
}
