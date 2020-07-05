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
