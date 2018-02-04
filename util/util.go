package util

import (
	"time"
	"strconv"
	"encoding/json"
	"bytes"
	"reflect"
)

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func FormatInt(i int64) string {
	return strconv.FormatInt(int64(i), 10)
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer // or out := bytes.Buffer{}
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

// source: https://github.com/forestgiant/sliceutil/
// Contains checks if a slice contains an element
func Contains(s interface{}, e interface{}) bool {
	slice := convertSliceToInterface(s)

	for _, a := range slice {
		if a == e {
			return true
		}
	}
	return false
}

// source: https://github.com/forestgiant/sliceutil/
// convertSliceToInterface takes a slice passed in as an interface{}
// then converts the slice to a slice of interfaces
func convertSliceToInterface(s interface{}) (slice []interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice {
		return nil
	}

	length := v.Len()
	slice = make([]interface{}, length)
	for i := 0; i < length; i++ {
		slice[i] = v.Index(i).Interface()
	}

	return slice
}
