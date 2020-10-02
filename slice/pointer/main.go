package main

import "fmt"

func main() {
	rtn := appendSlice()
	fmt.Printf("users %#v\n", rtn)
}

type User struct {
	Age     int          `json:"age"`
	Data    string       `json:"data"`
	Unknown map[int]bool `json:"unknown"`
}

func appendSlice() (users []*User) {
	unknown := map[int]bool{
		1: true,
	}

	user1 := User{
		Age:     1,
		Data:    "test1",
		Unknown: unknown,
	}

	user2 := User{
		Age:     2,
		Data:    "test2",
		Unknown: unknown,
	}

	users = append(users, &user1)
	users = append(users, &user2)

	return users
}
