package model

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Base struct {

	Extra map[string]interface{} `orm:"-" json:"extra"`

	ExtraJson string `orm:"type(text)" json:"-"`

	CreatedAt orm.DateTimeField `json:"created_at" orm:"auto_now_add"`
	UpdatedAt orm.DateTimeField `json:"updated_at" orm:"auto_now"`
}


func (b *Base) SetExtra(m map[string]interface{}) {
	mm, err := json.Marshal(m)
	if err != nil {
		logs.Error("marshal extra failed. m: %v, err: %v", m, err)
	}
	b.ExtraJson = string(mm)
}

func (b *Base) GetExtra() map[string]interface{} {
	if b.Extra == nil && b.ExtraJson != ""{
		b.Extra = make(map[string]interface{})
		err := json.Unmarshal([]byte(b.ExtraJson), b.Extra)
		if err != nil {
			logs.Error("unmarshal extra failed. extra-json: %v, err: %v", b.ExtraJson, err)
		}
		return b.Extra
	} else if b.Extra == nil {
		b.Extra = make(map[string]interface{})
		return b.Extra
	} else {
		return b.Extra
	}
}

func (b *Base) AddVariable(n string, p interface{}) {
	b.GetExtra()[n] = p
}