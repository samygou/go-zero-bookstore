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

func DeleteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteBookReq
		if err := httpx.Parse(r, &req); err != nil {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Err: errorx.NewError(errorx.CodeInvalidArgument, err.Error())})
			return
		}

		l := bookstoreCenter.NewDeleteBookLogic(r.Context(), svcCtx)
		resp, err := l.DeleteBook(&req)
		if err != nil {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Err: err})
		} else {
			svcCtx.RecoverMiddlewareCtx.RequestCtx.ResponseJson(w, &result.JsonResponse{Value: result.M{
				"data": resp,
			}})
		}
	}
}
