package logic

import (
	"context"
	"errors"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"
	"go-zero-bookstore/app/usercenter/repository"
	"go-zero-bookstore/common/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(in *usercenter.UpdateUserReq) (*usercenter.UpdateUserResp, error) {
	account, err := l.svcCtx.AccountModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrAccountNotExists
	}

	logx.Infof("%+v", account)

	// 更新的时候更全全部字段, 创建时间和更新时间系统默认更新, 不能自主更新
	err = l.svcCtx.AccountModel.Update(l.ctx, &repository.Accounts{
		Id:       in.Id,
		Mobile:   in.Mobile,
		Username: in.Username,
		Password: account.Password,
		Sex:      in.Sex,
		Avatar:   in.Avatar,
		Remark:   in.Remark,
	})
	if err != nil {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}

	return &usercenter.UpdateUserResp{}, nil
}
