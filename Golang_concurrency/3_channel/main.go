package main

import (
	"fmt"
	"time"
)

func main() {
	type User struct {
		ID   int
		Name string
	}

	ch := make(chan User)

	go func() {
		time.Sleep(200 * time.Millisecond)

		ch <- User{ID: 100, Name: "Kunal"}
	}()

	fmt.Println("main : waiting to recieve user")

	u := <-ch

	fmt.Println("the recived data in main :", u, u.ID, u.Name)
}
