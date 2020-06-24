package main

func main() {
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ",", j, " ")
			break Loop
		}
		println()
	}
}
