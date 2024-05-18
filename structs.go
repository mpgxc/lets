package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address Address
}

type Employee struct {
	User
	Address
}

type Address struct {
	City  string
	State string
}

func (o User) Print() {
	fmt.Printf("Name: %s\nAge:%d\nAddress.City:%s\n", o.Name, o.Age, o.Address)
}

func (o Employee) Print() {
	fmt.Printf("Name: %s\nAge:%d\nAddress.City:%s\n", o.Name, o.Age, o.Address)
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

	employee := Employee{
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
	employee.Address.City = "Rio de Janeiro"

	fmt.Println(employee.Name, employee.Age, employee.City, employee.State)

	// Métodos
	user2.Print()
	employee.Print()
}
