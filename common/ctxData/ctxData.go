package ctxData

import (
	"context"
	"encoding/json"

	"go-zero-bookstore/common/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "account_id"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}
