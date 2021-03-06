package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Page2  int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   10,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   20,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	jsonStr := `{"page": 1, "fruits": ["apple", "peach"] }`
	data := response2{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		panic(err)
	}
	fmt.Printf("data %#v\n", data)

	jsonStr1 := `{"num":6.13,"strs":["a","b"]}`
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr1), &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num, ok := dat["num"].(float64)
	fmt.Println(num)
	fmt.Println(ok)

	nums, ok := dat["nums"].(float64)
	fmt.Println(nums)
	fmt.Println(ok)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	m1 := make(map[int]string, 0)
	m1[0] = "2020-09-16 07:00:00"

	m12Json, _ := json.Marshal(m1)
	fmt.Printf("m12Json = %v\n", string(m12Json))

	m1Res := make(map[int]string, 0)
	json.Unmarshal(m12Json, &m1Res)
	fmt.Printf("m1Res = %#v\n", m1Res)

}
