package util

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"

)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if IsError(err) {
		logs.Error("get url failed. url: %s, cause: %s", url, err)
		return nil, errors.New("获取信息失败")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if IsError(err) {
		logs.Error("read body failed. body: ", body)
		return nil, errors.New("获取信息失败")
	}

	return body, nil
}
