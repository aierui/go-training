package main

import "fmt"

//var container = []string{"zero", "one", "two"}

type Key struct {
	Path, Country string
}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])

	/*
		The element is "one".
	*/

	hits := make(map[string]map[string]int)
	//n := hits["/doc/"]["au"]
	add(hits, "/doc/", "au")
	fmt.Println(hits)

	/*
		map[/doc/:map[au:1]]
	*/

	hits2 := make(map[Key]int)

	hits2[Key{"/", "vn"}]++
	hits2[Key{"/ref/spec", "ch"}] += 28
	fmt.Println(hits2)
	fmt.Println(len(hits2))
	n := hits2[Key{"/", "a"}] // read
	fmt.Println(n)

	m := hits2[Key{"/ref/spec", "ch"}] // read
	fmt.Println(m)

	/*
		map[{/ vn}:1 {/ref/spec ch}:28]
		2
		0
		28
	*/
}

func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}
