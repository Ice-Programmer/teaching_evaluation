package utils

import "time"

func GetNowSecs() int64 {
	return time.Now().Unix()
}
