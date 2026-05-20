package main

import "fmt"

type User struct {
	id    int
	name  string
	email string
	age   int
}

func main() {

	u1 := User{

		id:    1,
		name:  "Kunal",
		email: "paste@gmail.com",
		age:   10,
	}

	fmt.Println(u1, u1.name, u1.age)

	u2 := User{
		name: "Shubham",
		age:  10,
	}

	// struct is mutable

	fmt.Println(u2, u2.age)

	u2.age = 11

	fmt.Println(u2, u2.age)
}
