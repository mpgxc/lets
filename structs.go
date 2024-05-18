package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	City  string
	State string
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

	user.Address = Address{
		City:  "Oeiras",
		State: "Piauí",
	}

	fmt.Println(user.Address.City, user.Address.State)

	user.Address = Address{
		State: "Rio de Janeiro",
	}

	fmt.Println(user.Address.City, user.Address.State)

	user2 := User{"JaneDoe", 30, Address{
		City:  "Oeiras",
		State: "Piauí",
	}}

	fmt.Println(user2.Name, user2.Age, user2.Address.City, user2.Address.State)
}
