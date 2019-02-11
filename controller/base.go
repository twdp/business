package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"tianwei.pro/business"
)

type RestfulController struct {
	beego.Controller
}

func (r *RestfulController) ReturnJson(data interface{}) {
	r.Data["json"] = data
	r.ServeJSON()
	r.StopRun()
}

var (
	ReadBodyFailed = errors.New("读取请求信息失败")
)

func (r *RestfulController) ReadBody(result interface{}) error {
	b := r.Ctx.Input.RequestBody
	err := json.Unmarshal(b, result)
	if business.IsError(err) {
		logs.Error("read body failed. %v", err)
		return ReadBodyFailed
	}
	return nil
}

func (r *RestfulController) E500(body string) {
	r.Code(500, body)
}

func (r *RestfulController) Code(code int, body string) {
	out := r.Ctx.Output
	out.SetStatus(code)
	out.Body([]byte(body))
	r.StopRun()
}

func (r *RestfulController) Return(body interface{}) {
	b, err := json.Marshal(body)
	if business.IsError(err) {
		logs.Error("read body failed. %v", err)
	}
	r.Code(200, string(b))
}