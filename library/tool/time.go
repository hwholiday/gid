package tool

import "time"

func GetTimeUnix() int64 {
	return time.Now().Unix()
}
func GetTimeUnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}
