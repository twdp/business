package model

import "github.com/astaxie/beego/orm"

type Base struct {
	Id        int64             `json:"id"`
	CreatedAt orm.DateTimeField `json:"created_at" orm:"auto_now_add"`
	UpdatedAt orm.DateTimeField `json:"updated_at" orm:"auto_now"`
}
