package user

import (
	"context"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenterclient"
	"go-zero-bookstore/common/ctxData"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail() (resp *types.UserInfoResp, err error) {
	accountId := ctxData.GetUidFromCtx(l.ctx)

	accountDetailResp, err := l.svcCtx.Usercenter.GetUserInfo(l.ctx, &usercenterclient.GetUserInfoReq{
		Id: accountId,
	})
	if err != nil {
		return nil, err
	}

	userInfo := types.User{
		Id:       accountId,
		Mobile:   accountDetailResp.User.Mobile,
		Username: accountDetailResp.User.Username,
		Sex:      accountDetailResp.User.Sex,
		Avator:   accountDetailResp.User.Avatar,
		Remark:   accountDetailResp.User.Remark,
	}

	return &types.UserInfoResp{UserInfo: userInfo}, nil
}
