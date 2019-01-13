package util

import (
	"time"
	"fmt"
)

/**
 * @param uid 用户id
 * @param bizCode 业务编码
 * 订单来源1位+年月日6位+自增序号8位+用户ID取余100（16位）
 * - 订单来源：1: 微信小程序
 * - 时间信息:两位年，两位月，两位日（6位）
 * - 自增序号7位：采用随机步长进行自增，保证不重复同时保证不被猜出订单增长规律
 * - 用户标示2位
 * @return uuid
 */
func Uuid(uid int64, bizCode int) int64 {
	now := time.Now()
	serialize := CastInt64ToString(now.UnixNano() / 1000)[6:]
	if 0 == uid {
		return CastStringToInt64(fmt.Sprintf("%d%d%d%d%s", bizCode, now.Year()%100, now.Month(), now.Day(), serialize))
	} else {
		return CastStringToInt64(fmt.Sprintf("%d%d%d%d%s%d", bizCode, now.Year()%100, now.Month(), now.Day(), serialize, uid%100))
	}
}

func UuidByDefault() int64 {
	return Uuid(1024, 1)
}
