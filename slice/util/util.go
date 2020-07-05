package util

import "reflect"

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
