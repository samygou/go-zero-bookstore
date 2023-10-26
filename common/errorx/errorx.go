package errorx

import "fmt"

type ErrCode int32

const (
	CodeUnknown            ErrCode = 10000 //未知错误
	CodeUnauthenticated    ErrCode = 10001 //没有授权,或者授权不被认可
	CodeInvalidArgument    ErrCode = 10002 //参数错误
	CodePermissionDenied   ErrCode = 10003 //没有权限
	CodeNotFound           ErrCode = 10004 //目标未找到
	CodeAlreadyExists      ErrCode = 10005 //目标已存在
	CodeCanceled           ErrCode = 10006 //请求被取消
	CodeDeadlineExceeded   ErrCode = 10007 //过期了
	CodeResourceExhausted  ErrCode = 10008 //资源耗尽
	CodeFailedPrecondition ErrCode = 10009 //前提条件故障
	CodeAborted            ErrCode = 10010 //断言
	CodeOutOfRange         ErrCode = 10011 //超过范围
	CodeUnimplemented      ErrCode = 10012 //未实现
	CodeInternal           ErrCode = 10013 //内部错误
	CodeUnavailable        ErrCode = 10014 //对象不可用
	CodeDataLoss           ErrCode = 10015 //数据丢失
)

type Status struct {
	Code    ErrCode
	Message string
}

func (s *Status) Error() string {
	return fmt.Sprintf("%v:%v", s.Code, s.Message)
}

func NewError(code ErrCode, format string, args ...interface{}) *Status {
	return &Status{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

func (s *Status) NewMessage(message string) *Status {
	ns := *s
	ns.Message = message

	return &ns
}
