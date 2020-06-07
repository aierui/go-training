package main

func main() {
	s := "123"
	ps := &s
	b := []byte(*ps)
	c := []byte(s)
	pb := &b

	s += "4"
	*ps += "5"
	b[1] = '0'

	println(c)
	println(*ps)
	println(string(*pb))
}
