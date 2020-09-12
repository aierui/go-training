package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, playground")

	data := make(map[string]interface{})

	data["item_id"] = "23230.00000"

	fmt.Printf("item_id Type %T item_id Value %v\n", data["item_id"], data["item_id"])

	//n, err := TransToInt(data["item_id"])

	//fmt.Println(err)

	itemId := TransToString(data["item_id"])
	fmt.Printf("itemId Type %T itemId Value %v\n", itemId, itemId)
	i0, err := strconv.ParseFloat(itemId, 64)

	fmt.Println(err)
	fmt.Printf("i0 Type %T i0 Value %v\n", i0, i0)

	i5 := TransToString(int(i0))
	fmt.Printf("i5 Type %T i5 Value %v\n", i5, i5)

	//i1, err := strconv.ParseInt(itemId, 10, 64)
	//// i, errInt := strconv.ParseInt(v.(string), 10, 64)
	//fmt.Println(err)

	//i2 := TransToString(i1)
	//
	//m := TransToString(data["item_id"])
	//
	//z := TransToString(n)
	//
	//fmt.Printf("i1 Type %T i1 Value %v\n", i1, i1)
	//fmt.Printf("i2 Type %T i2 Value %v\n", i2, i2)
	//fmt.Printf("n Type %T n Value %v\n", n, n)
	//fmt.Printf("m Type %T m Value %v\n", m, m)
	//fmt.Printf("z Type %T z Value %v\n", z, z)
}

func TransToInt(data interface{}) (res int, err error) {
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

func TransToString(data interface{}) (res string) {
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
