package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/logic/user"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/types"
	"go-zero-bookstore/common/errorx"
	"go-zero-bookstore/common/result"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Err: errorx.NewError(errorx.CodeInvalidArgument, err.Error())})
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Err: err})
			return
		} else {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Value: result.M{
				"data": resp,
			}})
			return
		}
	}
}
