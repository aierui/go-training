package main

import (
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
)

type T struct {
	B  uint8 // is a byte
	I  int   // it is int32 on my x86 32 bit PC
	P  *int  // it is int32 on my x86 32 bit PC
	S  string
	SS []string
}

var p = fmt.Println

//In this case, the t := T{} can not measured by this method.

func memUsage(m1, m2 *runtime.MemStats) {
	p("Alloc:", m2.Alloc-m1.Alloc,
		"TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc,
		"HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}

func main() {

	str := "aierui"
	fmt.Printf("value => %v valueType => %T valueMemorySize =>%v \n", str, str, unsafe.Sizeof(str))

	var bigID1 int64
	bigID1 = 12345678910
	fmt.Printf("value => %v valueType => %T valueMemorySize =>%v \n", bigID1, bigID1, unsafe.Sizeof(bigID1))

	var bigID2 int64
	bigID2 = 1
	fmt.Printf("value => %v valueType => %T valueMemorySize =>%v \n", bigID2, bigID2, unsafe.Sizeof(bigID2))

	var tinyID int
	tinyID = 1
	fmt.Printf("value => %v valueType => %T valueMemorySize =>%v \n", tinyID, tinyID, unsafe.Sizeof(tinyID))

	var tinyID2 int
	tinyID2 = 999
	fmt.Printf("value => %v valueType => %T valueMemorySize =>%v \n", tinyID2, tinyID2, unsafe.Sizeof(tinyID2))

	var tinyID3 int8
	tinyID3 = 100
	fmt.Printf("value => %v valueType => %T valueMemorySize =>%v \n", tinyID3, tinyID3, unsafe.Sizeof(tinyID3))

	type Person struct {
		Name string `json:"name"`
		Age  uint   `json:"age"`
	}

	p1 := &Person{
		Name: "aierui",
		Age:  25,
	}

	fmt.Printf("value => %v valueType => %T valueMemorySize =>%v \n", p1, p1, unsafe.Sizeof(p1))

	/*
		value => aierui valueType => string valueMemorySize =>16
		value => 12345678910 valueType => int64 valueMemorySize =>8
		value => 1 valueType => int64 valueMemorySize =>8
		value => 1 valueType => int valueMemorySize =>8
		value => 999 valueType => int valueMemorySize =>8
		value => &{aierui 25} valueType => *main.Person valueMemorySize =>8
		value => 100 valueType => int8 valueMemorySize =>1
	*/

	const PtrSize = 32 << uintptr(^uintptr(0)>>63)
	p("PtrSize=", PtrSize)
	p("IntSize=", strconv.IntSize)
	var m1, m2, m3, m4, m5, m6 runtime.MemStats
	runtime.ReadMemStats(&m1)
	t := T{}
	runtime.ReadMemStats(&m2)
	p("sizeof(uint8)", unsafe.Sizeof(t.B),
		"offset=", unsafe.Offsetof(t.B))
	p("sizeof(int)", unsafe.Sizeof(t.I),
		"offset=", unsafe.Offsetof(t.I))
	p("sizeof(*int)", unsafe.Sizeof(t.P),
		"offset=", unsafe.Offsetof(t.P))
	p("sizeof(string)", unsafe.Sizeof(t.S),
		"offset=", unsafe.Offsetof(t.S))

	p("sizeof([]string)", unsafe.Sizeof(t.SS),
		"offset=", unsafe.Offsetof(t.SS))

	p("sizeof(T)", unsafe.Sizeof(t))

	memUsage(&m1, &m2)

	runtime.ReadMemStats(&m3)
	t2 := "abc"
	runtime.ReadMemStats(&m4)
	memUsage(&m3, &m4)
	//map will alloc memory in heap

	runtime.ReadMemStats(&m5)
	t3 := map[int]string{1: "x"}
	runtime.ReadMemStats(&m6)
	memUsage(&m5, &m6)
	fmt.Println(t2, t3) // prevent compiler error

	/*
		PtrSize= 64
		IntSize= 64
		sizeof(uint8) 1 offset= 0
		sizeof(int) 8 offset= 8
		sizeof(*int) 8 offset= 16
		sizeof(string) 16 offset= 24
		sizeof([]string) 24 offset= 40
		sizeof(T) 64
		Alloc: 0 TotalAlloc: 0 HeapAlloc: 0
		Alloc: 0 TotalAlloc: 0 HeapAlloc: 0
		Alloc: 256 TotalAlloc: 256 HeapAlloc: 256
		abc map[1:x]
	*/
}
