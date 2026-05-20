package main

import "fmt"

func main() {
	views := []int{100, 200, 300, 400, 500}

	total := 0

	for i, v := range views {
		fmt.Println("day : ", i, " views : ", v)
		total += v
	}

	fmt.Println(total)
}
