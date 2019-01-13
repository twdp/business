package util

import "github.com/astaxie/beego/orm"

func DateTimeFormat(in orm.DateTimeField) (out string) {
	return in.Value().Format(DateTimeFormatString)
}
