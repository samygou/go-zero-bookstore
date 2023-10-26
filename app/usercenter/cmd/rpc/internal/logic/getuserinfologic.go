package logic

import (
	"context"
	"errors"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"
	"go-zero-bookstore/app/usercenter/repository"
	"go-zero-bookstore/common/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *usercenter.GetUserInfoReq) (*usercenter.GetUserInfoResp, error) {
	account, err := l.svcCtx.AccountModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrAccountNotExists
	}

	userInfo := &usercenter.User{
		Id:       account.Id,
		Mobile:   account.Mobile,
		Username: account.Username,
		Sex:      account.Sex,
		Avatar:   account.Avatar,
		Remark:   account.Remark,
	}

	return &usercenter.GetUserInfoResp{
		User: userInfo,
	}, nil
}
