package result

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	// ApplicationJson stands for application/json.
	ApplicationJson = "application/json"
	// ContentType is the header key for Content-Type.
	ContentType = "Content-Type"
	// JsonContentType is the content type for JSON.
	JsonContentType = "application/json; charset=utf-8"
)

type M = map[string]interface{}

type Param struct {
	key   string
	exist bool
	v     string
}

type Context struct {
	RequestId string
	StartTime int64
	HTTPCtx   context.Context
	Resp      Response
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) GetRequestTime() int64 {
	return c.StartTime
}

func (c *Context) SetResponse(resp Response) {
	if c.Resp != nil {
		return
	}
	c.Resp = resp
	c.Resp.Set(c)
}

func (c *Context) ResponseJson(w http.ResponseWriter, resp ...Response) {
	if len(resp) > 0 {
		resp[0].Set(c)
		c.Resp = resp[0]
	}

	httpCode := http.StatusOK
	if c.Resp.GetHttpCode() != 0 {
		httpCode = c.Resp.GetHttpCode()
	}

	if err := c.doWriteJson(w, httpCode); err != nil {
		logx.Error(err)
	}
}

func (c *Context) doWriteJson(w http.ResponseWriter, code int) error {
	w.Header().Set(ContentType, JsonContentType)
	w.WriteHeader(code)

	if n, err := w.Write(c.Resp.GetContent()); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if !errors.Is(err, http.ErrHandlerTimeout) {
			return fmt.Errorf("write response failed, error: %w", err)
		}
	} else if n < len(c.Resp.GetContent()) {
		return fmt.Errorf("actual bytes: %d, written bytes: %d", len(c.Resp.GetContent()), n)
	}

	return nil
}
