package reflect_declare

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Age  int    `json:"age" tag1:"age" tag2:"age"`
	Name string `json:"name" tag1:"name"`
	City int    `json:"city" tag2:"city"`
}

func (c Person) PrintName() {
	fmt.Printf("Person Name: %v\n", c.Name)
}

func (c Person) SetName(v string) {
	c.Name = v
}

func TestReflectDeclare(t *testing.T) {

	h1 := Person{
		Age:  25,
		Name: "aierui",
		City: 1,
	}

	h1Type := reflect.TypeOf(h1)

	fmt.Printf("h1Type.Name %v\n", h1Type.Name())
	fmt.Printf("h1Type.Align %v\n", h1Type.Align())
	fmt.Printf("h1Type.FieldAlign %v\n", h1Type.FieldAlign())
	fmt.Printf("h1Type.Method %v\n", h1Type.Method(0))
	fmt.Printf("h1Type.NumMethod %v\n", h1Type.NumMethod())

	method, ok := h1Type.MethodByName("SetName")
	fmt.Printf("h1Type.MethodByName method %#v, ok %v\n", method, ok)

	fmt.Printf("h1Type.PkgPath %v\n", h1Type.PkgPath())
	fmt.Printf("h1Type.String %v\n", h1Type.String())
	fmt.Printf("h1Type.Kind %v\n", h1Type.Kind())

	/*
		h1Type.Name Person
		h1Type.Align 8
		h1Type.FieldAlign 8
		h1Type.Method {PrintName  func(reflect_declare.Person) <func(reflect_declare.Person) Value> 0}
		h1Type.NumMethod 2
		h1Type.MethodByName method reflect.Method{Name:"SetName", PkgPath:"", Type:(*reflect.rtype)(0xc0000582a0), Func:reflect.Value{typ:(*reflect.rtype)(0xc0000582a0), ptr:(unsafe.Pointer)(0xc00000e070), flag:0x13}, Index:1}, ok true
		h1Type.PkgPath github.com/aierui/go-training/reflect/declare
		h1Type.String reflect_declare.Person
		h1Type.Kind struct
	*/

	h2 := &Person{
		Age:  25,
		Name: "aierui",
		City: 1,
	}

	h2Type := reflect.TypeOf(h2)

	fmt.Printf("h2Type.Name %v\n", h2Type.Name())
	fmt.Printf("h2Type.Align %v\n", h2Type.Align())
	fmt.Printf("h2Type.FieldAlign %v\n", h2Type.FieldAlign())
	fmt.Printf("h2Type.Method %v\n", h2Type.Method(0))
	fmt.Printf("h2Type.NumMethod %v\n", h2Type.NumMethod())

	method2, ok2 := h2Type.MethodByName("SetName")
	fmt.Printf("h2Type.MethodByName method %#v, ok %v\n", method2, ok2)

	fmt.Printf("h2Type.PkgPath %v\n", h2Type.PkgPath())
	fmt.Printf("h2Type.String %v\n", h2Type.String())
	fmt.Printf("h2Type.Kind %v\n", h2Type.Kind())

	/*
		h2Type.Name
		h2Type.Align 8
		h2Type.FieldAlign 8
		h2Type.Method {PrintName  func(*reflect_declare.Person) <func(*reflect_declare.Person) Value> 0}
		h2Type.NumMethod 2
		h2Type.MethodByName method reflect.Method{Name:"SetName", PkgPath:"", Type:(*reflect.rtype)(0xc000118360), Func:reflect.Value{typ:(*reflect.rtype)(0xc000118360), ptr:(unsafe.Pointer)(0xc0001160a0), flag:0x13}, Index:1}, ok true
		h2Type.PkgPath
		h2Type.String *reflect_declare.Person
		h2Type.Kind ptr
	*/

}
