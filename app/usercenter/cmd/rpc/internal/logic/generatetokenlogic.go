package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"
	"go-zero-bookstore/common/errorx"
	"go-zero-bookstore/common/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *usercenter.GenerateTokenReq) (*usercenter.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpired := now + l.svcCtx.Config.JWTAuth.AccessExpired
	fmt.Println(l.svcCtx.Config.JWTAuth.AccessSecret)
	accessToken, err := l.generateToken(l.svcCtx.Config.JWTAuth.AccessSecret, in.UserId, accessExpired)
	if err != nil {
		logx.Error(err)
		return nil, errorx.NewError(errorx.CodeInternal, "Generate access token failed!")
	}

	return &usercenter.GenerateTokenResp{Token: accessToken}, nil
}

func (l *GenerateTokenLogic) generateToken(secretKey string, accountId int64, expired int64) (string, error) {
	claim := make(jwt.MapClaims)
	claim["expired"] = expired
	claim["account_id"] = accountId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claim

	return token.SignedString([]byte(secretKey))
}
