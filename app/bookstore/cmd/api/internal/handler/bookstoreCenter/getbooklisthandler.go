package bookstoreCenter

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-bookstore/app/bookstore/cmd/api/internal/logic/bookstoreCenter"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/types"
	"go-zero-bookstore/common/errorx"
	"go-zero-bookstore/common/result"
)

func GetBookListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBookListReq
		if err := httpx.Parse(r, &req); err != nil {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Err: errorx.NewError(errorx.CodeInvalidArgument, err.Error())})
			return
		}

		l := bookstoreCenter.NewGetBookListLogic(r.Context(), svcCtx)
		resp, err := l.GetBookList(&req)
		if err != nil {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Err: err})
		} else {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Value: result.M{
				"data": resp,
			}})
		}
	}
}
