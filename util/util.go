package util

import (
	"time"
	"strconv"
)

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func FormatInt(i int64) string {
	return strconv.FormatInt(int64(i), 10)
}
