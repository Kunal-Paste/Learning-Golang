package main

import (
	"fmt"
)

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("MONDAY")
	case 2:
		fmt.Println("TUESDAY")
	case 3:
		fmt.Println("WEDNESDAY")
	default:
		fmt.Println("ANY DAY")
	}
}
