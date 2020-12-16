package main

import "fmt"

func main() {

	m := make(map[string]interface{}, 0)

	m["k1"] = "v1"
	m["k2"] = 2

	fmt.Println(m)

	err := setValue(m)
	fmt.Println(err)
	fmt.Println(m)

	arr := []string{"1", "2"}
	fmt.Println(arr)

	err = setSlice(arr)
	fmt.Println(err)
	fmt.Println(arr)

	/*
		map[k1:v1 k2:2]
		<nil>
		map[k1:v1 k2:2 set_k1:set_v1]
		[1 2]
		<nil>
		[1 2]
	*/
}

func setValue(m map[string]interface{}) error {

	m["set_k1"] = "set_v1"

	return nil
}

func setSlice(arr []string) error {

	arr = append(arr, "set_slice1", "set_slice2")

	return nil
}
