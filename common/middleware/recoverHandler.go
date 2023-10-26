package middleware

import (
	"net/http"
	"time"

	"github.com/rs/xid"

	"go-zero-bookstore/common/errorx"
	"go-zero-bookstore/common/logx"
	"go-zero-bookstore/common/result"
)

type RecoverMiddleware struct {
	RequestCtx *result.Context
}

var RecoverMiddlewareHandler *RecoverMiddleware

func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{
		RequestCtx: result.NewContext(),
	}
}

// Handle returns a middleware that recovers if panic happens.
func (m *RecoverMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.RequestCtx.RequestId = xid.New().String()
		m.RequestCtx.StartTime = time.Now().UnixNano() / 1000000
		m.RequestCtx.HTTPCtx = r.Context()

		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case errorx.Status:
					logx.Error(err)
				default:
					logx.Error(err)
					m.RequestCtx.ResponseJson(w, &result.JsonResponse{Err: errorx.NewError(errorx.CodeUnknown, "unknown err")})
				}
			}
		}()

		next.ServeHTTP(w, r)
	}
}
