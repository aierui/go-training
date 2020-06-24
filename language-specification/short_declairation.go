package main



func main () {

	//Assuming x is declared and y is not declared, which clauses below are correct?


	x, _ := f() //incorrect
	x, _ = f() //correct
	x, y := f() //correct
	x, y = f() //incorrect
}