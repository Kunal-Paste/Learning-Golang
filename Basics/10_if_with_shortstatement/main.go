package main

import (
	"fmt"
)

func main() {
	price := 10
	products := 50

	if total := price * products; total > 300 {
		fmt.Println("discount applicable")
	} else {
		fmt.Println("buy more products for discount!!")
	}
}
