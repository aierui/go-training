package main


type S struct {
	m string
}

func f() *S {
	//return __ //A

	return &S{"foo"};
}

func main() {
	//p := __    //B
	p := *f()
	print(p.m) //print "foo"
}
