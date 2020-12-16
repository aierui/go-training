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

	m2 := make(map[string]int, 0)
	m2V := m2["k2"]
	fmt.Printf("m2V %T->%v\n", m2V, m2V)

	m3 := make(map[string]string, 0)
	m3V := m3["k3"]
	fmt.Printf("m3V %T->%v\n", m3V, m3V)

	m4 := make(map[string]interface{}, 0)
	m4V := m4["k4"]
	fmt.Printf("m4V %T->%v\n", m4V, m4V)

	m5 := make(map[string]float64, 0)
	m5V := m5["k5"]
	fmt.Printf("m5V %T->%v\n", m5V, m5V)

	/*
		m2V int->0
		m3V string->
		m4V <nil>-><nil>
		m5V float64->0
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
