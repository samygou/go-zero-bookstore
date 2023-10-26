package user

import (
	"context"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateUserReq) error {
	_, err := l.svcCtx.Usercenter.UpdateUser(l.ctx, &usercenter.UpdateUserReq{
		Id:       req.Id,
		Mobile:   req.Mobile,
		Username: req.Username,
		Sex:      req.Sex,
		Avatar:   req.Avatar,
		Remark:   req.Remark,
	})
	if err != nil {
		return err
	}

	return nil
}
