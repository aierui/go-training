package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasElem(t *testing.T) {
	ast := assert.New(t)

	sliceString := []string{"apple", "iPhone", "watch"}
	ast.Equal(true, HasElem(sliceString, "apple"))
	ast.Equal(false, HasElem(sliceString, "a"))

	sliceInt := []int{1, 2, 3, 4, 5}
	ast.Equal(true, HasElem(sliceInt, 1))
	ast.Equal(false, HasElem(sliceInt, 10))

	sliceInterface := []interface{}{1, 2, "apple", "iPhone"}

	ast.Equal(true, HasElem(sliceInterface, 1))
	ast.Equal(true, HasElem(sliceInterface, "apple"))
	ast.Equal(false, HasElem(sliceInterface, "maybe"))
}

func TestToInt(t *testing.T) {
	ast := assert.New(t)

	data1 := map[string]interface{}{
		"int": "1212",
	}
	rtn, err := ToInt(data1["int"])
	if err != nil {
		t.Fatalf("TestToInt err %v\n", err)
	}

	rtn, err = ToInt(int8(12))
	if err != nil {
		t.Fatalf("TestToInt err %v\n", err)
	}
	ast.Equal(12, rtn)

	rtn, err = ToInt(int32(3212))
	if err != nil {
		t.Fatalf("TestToInt err %v\n", err)
	}
	ast.Equal(3212, rtn)

	rtn, err = ToInt(int64(311212))
	if err != nil {
		t.Fatalf("TestToInt err %v\n", err)
	}
	ast.Equal(311212, rtn)

	rtn, err = ToInt(float64(311.12))
	if err != nil {
		t.Fatalf("TestToInt err %v\n", err)
	}
	ast.Equal(311, rtn)
}
