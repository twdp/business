package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"runtime"
)

type RestfulErrorHandler struct {
	RestfulController
}


func (c *RestfulErrorHandler) Error404() {
	c.Code(404, "not found")
}

func (c *RestfulErrorHandler) Error501() {
	c.Code(501, "server error")
}


func (c *RestfulErrorHandler) ErrorDb() {
	c.Code(500, "db is now down")
}

func (c *RestfulErrorHandler) Error401() {
	c.Code(401, "unauthorized")
}

func (c *RestfulErrorHandler) Error403() {
	c.Code(403, "forbidden")
}

func (c *RestfulErrorHandler) Error503() {
	c.Code(503, "ServiceUnavailable ")
}

var (
	DiyHandle func(err error, ctx *context.Context)
)

func RestfulErrorHandle(ctx *context.Context) {
	if err := recover(); err != nil {
		if err == beego.ErrAbort {
			return
		}
		logs.Critical("the request url is ", ctx.Input.URL())
		logs.Critical("tw business module handler crashed with error", err)
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			logs.Critical(fmt.Sprintf("%s:%d", file, line))
			//stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
		}

		if DiyHandle != nil {
			DiyHandle(err.(error), ctx)
		} else {
			ctx.ResponseWriter.Write([]byte(fmt.Sprintf("{\"msg\":\"%s\"}", err.(error).Error())))
			if ctx.Output.Status != 0 {
				ctx.ResponseWriter.WriteHeader(ctx.Output.Status)
			} else {
				ctx.ResponseWriter.WriteHeader(500)
			}
		}
	}
}