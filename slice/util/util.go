package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// HasElem
func HasElem(s interface{}, elem interface{}) bool {
	v := reflect.ValueOf(s)
	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

// ToInt convert interface data into int type
func ToInt(data interface{}) (res int, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = int(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = int(val.Uint())
	case float64:
		res = int(data.(float64))
	case float32:
		res = int(data.(float32))
	case string:
		res, err = strconv.Atoi(strings.TrimSpace(data.(string)))
	case []byte:
		res, err = strconv.Atoi(strings.TrimSpace(string(data.([]byte))))
	case json.Number:
		var resInt64 int64
		resInt64, err = data.(json.Number).Int64()
		res = int(resInt64)
	default:
		res, err = strconv.Atoi(fmt.Sprintf("%v", data))
	}

	return
}

// ToInt64 convert interface data into int64 type
func ToInt64(data interface{}) (res int64, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		res = int64(val.Uint())
	case float64:
		res = int64(data.(float64))
	case float32:
		res = int64(data.(float32))
	case json.Number:
		res, err = data.(json.Number).Int64()
	case string:
		res, err = strconv.ParseInt(strings.TrimSpace(data.(string)), 10, 64)
	case []byte:
		res, err = strconv.ParseInt(strings.TrimSpace(string(data.([]byte))), 10, 64)
	default:
		res, err = strconv.ParseInt(fmt.Sprintf("%v", data), 10, 64)
	}

	return
}

// ToString convert interface data into string type
func ToString(data interface{}) (res string, err error) {
	switch v := data.(type) {
	case float64:
		res = strconv.FormatFloat(data.(float64), 'f', 6, 64)
	case float32:
		res = strconv.FormatFloat(float64(data.(float32)), 'f', 6, 32)
	case int:
		res = strconv.FormatInt(int64(data.(int)), 10)
	case int64:
		res = strconv.FormatInt(data.(int64), 10)
	case uint:
		res = strconv.FormatUint(uint64(data.(uint)), 10)
	case uint64:
		res = strconv.FormatUint(data.(uint64), 10)
	case uint32:
		res = strconv.FormatUint(uint64(data.(uint32)), 10)
	case json.Number:
		res = data.(json.Number).String()
	case string:
		res = data.(string)
	case []byte:
		res = string(v)
	default:
		res = ""
	}

	return
}
