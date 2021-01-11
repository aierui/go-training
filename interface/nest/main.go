package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	nInt := [][]int{
		{0, 1, 2},
		{4, 5, 6},
	}
	nIntRtn, err := json.Marshal(nInt)
	fmt.Println(err)
	fmt.Printf("nIntRtn Type:%T, Value:%v\n", nIntRtn, string(nIntRtn))

	/*
		<nil>
		nIntRtn Type:[]uint8, Value:[[0,1,2],[4,5,6]]
	*/

	nStr := [][]string{
		{"a1", "b1", "c1"},
		{"a2", "b2", "c2"},
	}

	nStrRtn, err := json.Marshal(nStr)
	fmt.Println(err)
	fmt.Printf("nStrRtn Type:%T, Value:%v\n", nStrRtn, string(nStrRtn))

	/*
		<nil>
		nStrRtn Type:[]uint8, Value:[["a1","b1","c1"],["a2","b2","c2"]]
	*/

	nBool := [][]bool{
		{true, false},
		{true, false, true},
	}

	nBoolRtn, err := json.Marshal(nBool)
	fmt.Println(err)
	fmt.Printf("nBoolRtn Type:%T, Value:%v\n", nBoolRtn, string(nBoolRtn))

	/*
		<nil>
		nBoolRtn Type:[]uint8, Value:[[true,false],[true,false,true]]
	*/

	nInterface := [][]interface{}{
		{"a1", "b1", "c1"},
		{"a2", "b2", "c2"},
		{1, 2, 3},
		{"q", "2", 3, false},
	}

	nInterfaceRtn, err := json.Marshal(nInterface)
	fmt.Println(err)
	fmt.Printf("nInterface Type:%T, Value:%v\n", nInterfaceRtn, string(nInterfaceRtn))

	/*
	   <nil>
	   nInterface Type:[]uint8, Value:[["a1","b1","c1"],["a2","b2","c2"],[1,2,3],["q","2",3,false]]
	*/
}
