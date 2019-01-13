package business

import (
	"errors"
	"github.com/astaxie/beego/logs"

	"golang.org/x/crypto/bcrypt"
)

/**
 * bcrypt 加密
 */
func GenerateCrypto(pass string) (string, error) {
	r, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if IsError(err) {
		logs.Error("generate crypto error. pass: %v, err: %v", pass, err)
		err = errors.New("加密失败")
	}
	return string(r), err
}

/**
 * 对提交的密码与加密后的进行对比
 */
func ValidateCrypto(pass, encrypto string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encrypto), []byte(pass))
	if IsError(err) {
		return false, errors.New("密码错误")
	} else {
		return true, nil
	}
}
