package main

func main() {

	// incorrect
	var m map[string]int //A
	m["a"] = 1
	if v := m["b"]; v != nil { //B
		println(v)
	}


	// correct
	m := make(map[string]int)
	m["a"] = 1
	if v, ok := m["b"]; ok {
		println(v)
	}

}
