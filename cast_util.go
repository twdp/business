package util

import "strconv"

/**
 * 将int转换为字符串
 */
func CastIntToString(v int) string {
	return strconv.Itoa(v)
}

/**
 * 将int64转换为string
 */
func CastInt64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func CastIntToInt64(v int) int64 {
	return CastIntToInt64DefaultValue(v, 0)
}

func CastIntToInt64DefaultValue(v int, defaultValue int64) int64 {
	return CastStringToInt64DefaultValue(CastIntToString(v), defaultValue)
}

func CastStringToInt64(v string) int64 {
	return CastStringToInt64DefaultValue(v, 0)
}

func CastStringToInt64DefaultValue(v string, defaultValue int64) int64 {
	int64, err := strconv.ParseInt(v, 10, 64)
	if IsError(err) {
		return defaultValue
	}
	return int64
}

func CastStringToInt64ReturnError(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func CastStringToInt(v string) int {
	return CastStringToIntDefaultValue(v, 0)
}

func CastStringToIntDefaultValue(v string, defaultValue int) int {
	int, err := strconv.Atoi(v)
	if IsError(err) {
		return defaultValue
	}
	return int
}

func CastStringToIntReturnError(v string) (int, error) {
	return strconv.Atoi(v)
}
