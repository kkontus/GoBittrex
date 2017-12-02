package util

import (
	"time"
	"strconv"
	"encoding/json"
	"bytes"
)

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func FormatInt(i int64) string {
	return strconv.FormatInt(int64(i), 10)
}

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer // or out := bytes.Buffer{}
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
