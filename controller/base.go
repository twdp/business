package controller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"golang.org/x/xerrors"
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
	ReadBodyFailed = xerrors.New("读取请求信息失败")
)

func (r *RestfulController) ReadBody(result interface{}) error {
	b := r.Ctx.Input.RequestBody
	err := json.Unmarshal(b, result)
	if business.IsError(err) {
		logs.Error("read body failed. %+v", err)
		return ReadBodyFailed
	}
	return nil
}

func (r *RestfulController) E500(body interface{}) {
	r.Code(500, body)
}

func (r *RestfulController) Code(code int, body interface{}) {
	var b []byte
	switch body.(type) {
	case string:
		result := make(map[string]interface{})
		result["msg"] = body
		if bb, err := json.Marshal(result); err != nil {
			logs.Warn("json marshal failed.  body: %v, err: %v", body, err)
		} else {
			b = bb
		}
	default:
		if bb, err := json.Marshal(body); err != nil {
			logs.Warn("json marshal failed.  body: %v, err: %v", body, err)
		} else {
			b = bb
		}

	}
	out := r.Ctx.Output
	out.SetStatus(code)
	out.Body(b)
	r.StopRun()
}

func (r *RestfulController) Return(body interface{}) {
	r.Code(200, body)
}
