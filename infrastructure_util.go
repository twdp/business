package util

/**
 * 判断是否发生了错误
 */
func IsError(err error) bool {
	if nil != err {
		return true
	}
	return false
}

