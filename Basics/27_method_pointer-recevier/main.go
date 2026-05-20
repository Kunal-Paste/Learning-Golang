package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	u1 := User{name: "Kunal", age: 12}
	fmt.Println(u1, u1.age)

	u1.Birthday()

	fmt.Println("after updation : ", u1, u1.age)
}

func (u *User) Birthday() {
	u.age++
}
