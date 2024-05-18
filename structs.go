package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address Address
}

type employee struct {
	User
	Address
}

type Address struct {
	City  string
	State string
}

func (o User) print() {
	fmt.Printf("Name: %s\nAge:%d\nAddress.City:%s\n", o.Name, o.Age, o.Address)
}

func (o employee) print() {
	fmt.Printf("Name: %s\nAge:%d\nAddress.City:%s\n", o.User.Name, o.User.Age, o.User.Address)
}

type PrintInterface interface {
	print()
}

func ApplyPrint(o PrintInterface) {
	o.print()
}

func handlerStructs() {
	u := User{
		Name: "John Doe",
		Age:  30,
	}

	fmt.Println(u.Name, u.Age)

	u.Name = "Jane Doe"
	u.Age = 25

	fmt.Println(u.Name, u.Age)

	u.Address = Address{
		City:  "Oeiras",
		State: "Piauí",
	}

	fmt.Println(u.Address.City, u.Address.State)

	u.Address = Address{
		State: "Rio de Janeiro",
	}

	fmt.Println(u.Address.City, u.Address.State)

	user2 := User{"JaneDoe", 30, Address{
		City:  "Oeiras",
		State: "Piauí",
	}}

	fmt.Println(user2.Name, user2.Age, user2.Address.City, user2.Address.State)

	e := employee{
		User: User{
			Name: "JaneDoe",
			Age:  30,
		},
		Address: Address{
			City:  "Oeiras",
			State: "Piauí",
		},
	}

	user2.Address.City = "Rio de Janeiro"
	e.Address.City = "Rio de Janeiro"

	fmt.Println(e.User.Name, e.User.Age, e.Address.City, e.Address.State)

	// Métodos
	user2.print()
	e.print()
	ApplyPrint(user2)
	ApplyPrint(e)
}
