package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func forStruct() {
	var users [3]User

	users[0] = User{"tester1", 99}

	for _, user := range users {
		fmt.Println(user)
	}
}
