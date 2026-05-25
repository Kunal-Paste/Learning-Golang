package main

import (
	"fmt"
	"go-module/internal/greet"
)

func main() {
	msg := greet.Hello("Kunal")

	fmt.Println(msg)
}
