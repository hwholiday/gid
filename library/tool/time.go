package tool

import "time"

func GetTimeUnix() int64 {
	return time.Now().Unix()
}
