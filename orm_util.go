package business

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

/**
 * 事务开始
 */
func TransactionStart() (orm.Ormer, error) {
	o := orm.NewOrm()
	err := o.Begin()
	if IsError(err) {
		logs.Error("transaction  start failed. ")
		return nil, errors.New("连接数据库失败")
	}
	return o, nil
}

/**
 * 执行中
 */
func TransactionProcess(o orm.Ormer, err error) bool {
	if IsError(err) {
		o.Rollback()
		return true
	}
	return false
}

/**
 * 事务结束
 */
func TransactionEnd(o orm.Ormer, err error) bool {
	if IsError(err) {
		logs.Error("transaction end failed. error: %+v, o: %+v", err, o)
		o.Rollback()
		return true
	} else {
		o.Commit()
		return false
	}
}
