//时间格式化
package utils

import (
	"time"
)

//当前时间格式为：2006-1-2
func GetDate() string {
	return time.Now().Format("2006-1-2")
}

//当前时间格式为：2006-1-2 15:04:05
func GetDateTime() string {
	return time.Now().Format("2006-1-2 15:04:05")
}

//当前时间的UNIXTIME 秒
func GetUnix() int64 {
	return time.Now().Unix()
}

//当前时间的UNIXTIME 毫秒
func GetMilliUnix() int64 {
	return time.Now().UnixNano() / 1e6
}

//当前时间的UNIXTIME 纳秒
func GetNanoUnix() int64 {
	return time.Now().UnixNano()
}

//时间戳转时间类型time.Unix 再格式化成时间字符串Format()
func Unix2Date(sec int64) string {
	return time.Unix(sec, 0).Format("2006-1-2")
}

func Unix2DateTime(sec int64) string {
	return time.Unix(sec, 0).Format("2006-1-2 15:04:05")
}

//时间字符串转时间类型time.Parse 再转化成别的格式
func Date2Unix(date string) int64 {
	t,_ := time.Parse("2006-1-2 15:04:05", date)
	return t.Unix()
}