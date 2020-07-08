package sonwflake

import (
	"fmt"
	"testing"
)

func TestNextId(t *testing.T) {
	num, sliceLen, dup := checkDuplicated(10000000)
	fmt.Printf("num %v sliceLen %v dup %v \n", num, sliceLen, dup)
}

func BenchmarkNewNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, dup := checkDuplicated(10000000)
		if dup > 0 {
			fmt.Printf("dup %v \n", dup)
		}
	}
}

func checkDuplicated(num int64) (int64, int64, int64) {
	w := NewNode(1, 1)
	r := make([]int64, 0)
	end := int64(num)
	var i int64
	for ; i < end; i++ {
		id := w.NextId()
		r = append(r, id)
	}
	s := int64(len(unique(r)))

	return num, s, num - s
}

func unique(m []int64) []int64 {
	s := make([]int64, 0)
	sMap := make(map[int64]int64)
	for _, value := range m {
		length := len(sMap)
		sMap[value] = 1
		if len(sMap) != length {
			s = append(s, value)
		}
	}

	return s
}
