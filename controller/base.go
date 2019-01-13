package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"tianwei.pro/business"
)

type TemplateController struct {
	beego.Controller
}

func (t *TemplateController) E500(body string) {
	t.Code(500, body)
}

func (t *TemplateController) Code(code int, body string) {
	out := t.Ctx.Output
	out.SetStatus(code)
	out.Body([]byte(body))
}

func (t *TemplateController) Return(body string) {
	t.Code(200, body)
}

func (r *TemplateController) ReadBody(result interface{}) {
	b := r.Ctx.Input.RequestBody
	err := json.Unmarshal(b, result)
	if util.IsError(err) {
		logs.Error("read body failed. %v", err)
	}
}

type RestfulController struct {
	beego.Controller
}

func (r *RestfulController) ReturnJson(data interface{}) {
	r.Data["json"] = data
	r.ServeJSON()
	//r.StopRun()
}

func (r *RestfulController) ReadBody(result interface{}) {
	b := r.Ctx.Input.RequestBody
	err := json.Unmarshal(b, result)
	if util.IsError(err) {
		logs.Error("read body failed. %v", err)
	}
}
