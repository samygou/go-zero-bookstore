package user

import (
	"net/http"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/logic/user"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/svc"
	"go-zero-bookstore/common/result"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()
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
