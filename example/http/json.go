package main

import (
	"encoding/json"
	"fmt"
)

func jsonHandling() {
	type User struct {
		Name string
		Age  int
	}

	jsonString, _ := json.Marshal(User{"tester", 29})
	fmt.Println(string(jsonString))

	var user User
	_ = json.Unmarshal(jsonString, &user)

	fmt.Println(user.Name, user.Age)
}
