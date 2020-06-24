package main



package main

type S struct {
	name string
}

// Cannot assign to m["x"].name
func main() {
	m := map[string]S{"x": S{"one"}}
	m["x"].name = "two"
}



func main2() {
	m := map[string]*S{"x": &S{"one"}}
	m["x"].name = "two"
}