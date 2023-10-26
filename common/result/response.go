package result

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/status"

	"go-zero-bookstore/common/errorx"
)

type ResponseType int32

const (
	ResponseTypeJson ResponseType = iota + 1
	ResponseTypeFile
)

type Response interface {
	GetType() ResponseType
	GetErr() *errorx.Status
	GetContent() []byte
	Set(c *Context)
	GetHttpCode() int
}

type JsonResponse struct {
	Err        error
	err        *errorx.Status
	Value      M
	HttpStatus int
}

func (jr *JsonResponse) GetType() ResponseType {
	return ResponseTypeJson
}

func (jr *JsonResponse) GetContent() []byte {
	if jr.Value == nil {
		return nil
	}
	b, err := json.Marshal(jr.Value)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

func (jr *JsonResponse) GetErr() *errorx.Status {
	if jr.err == nil {
		jr.err = toResponseErr(jr.Err)
	}

	return jr.err
}

func (jr *JsonResponse) GetHttpCode() int {
	return jr.HttpStatus
}

func (jr *JsonResponse) Set(c *Context) {
	endTime := time.Now().UnixNano() / 1000000

	if jr.Value == nil {
		jr.Value = make(M)
	}

	jr.Value["process_cost"] = strconv.FormatInt(endTime-c.StartTime, 10) + " ms"
	jr.Value["request_id"] = c.RequestId
	jr.Value["code"] = 0
	jr.Value["message"] = "success"

	jr.err = toResponseErr(jr.Err)
	if jr.err != nil {
		jr.Value["code"] = jr.err.Code
		jr.Value["message"] = jr.err.Message
	}
}

type FileResponse struct {
	Err                error
	err                *errorx.Status
	Content            []byte
	ContentType        string
	ContentDisposition string
	HttpStatus         int
}

func (fr *FileResponse) GetType() ResponseType {
	return ResponseTypeFile
}

func (fr *FileResponse) GetContent() []byte {
	return fr.Content
}

func (fr *FileResponse) GetErr() *errorx.Status {
	if fr.err == nil {
		fr.err = toResponseErr(fr.Err)
	}

	return fr.err
}

func (fr *FileResponse) Set(c *Context) {}

func toResponseErr(err error) *errorx.Status {
	var out *errorx.Status

	if err != nil {
		switch err.(type) {
		case *errorx.Status:
			out = err.(*errorx.Status)
			return out
		default:
			if gstatus, ok := status.FromError(err); ok {
				message := gstatus.Message()
				errStatus := strings.Split(message, ":")
				if len(errStatus) >= 2 {
					errCode, err := strconv.Atoi(errStatus[0])
					// 如果错误解析失败, 返回grpc原始错误
					if err != nil {
						return &errorx.Status{
							Code:    errorx.ErrCode(gstatus.Code()),
							Message: message,
						}
					}
					return &errorx.Status{
						Code:    errorx.ErrCode(errCode),
						Message: errStatus[1],
					}
				}
				return &errorx.Status{
					Code:    errorx.CodeUnknown,
					Message: "unknown err",
				}
			}
			return &errorx.Status{
				Code:    errorx.CodeInternal,
				Message: "server err",
			}
		}
	}

	return nil
}
