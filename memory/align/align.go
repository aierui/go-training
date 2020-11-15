package align

type Ag struct {
	arr [2]int8  // 2
	bl  bool     // 1 padding 5
	sl  []int16  // 24
	ptr *int64   // 8
	st  struct { // 16
		str string
	}
	m map[string]int16
	i interface{}
}
