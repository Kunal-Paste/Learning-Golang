package main

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

func main() {
	u := User{name: "Kunal", age: 10, email: "paste@gmail.com"}

	fmt.Println(u.Intro())
}

// this method will receive a copy of user
func (u User) Intro() string {
	return fmt.Sprintf("Hi i am %s", u.name)
}
