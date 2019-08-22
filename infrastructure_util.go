package business

/**
 * 判断是否发生了错误
 */
func IsError(err error) bool {
	if nil != err {
		return true
	}
	return false
}

func Check(err error) {
	if nil != err {
		panic(err)
	}
}
