package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func handlerStructs() {
	user := User{
		Name: "John Doe",
		Age:  30,
	}

	fmt.Println(user.Name, user.Age)

	user.Name = "Jane Doe"
	user.Age = 25

	fmt.Println(user.Name, user.Age)
}
