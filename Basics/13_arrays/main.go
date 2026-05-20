package main

import "fmt"

func main() {
	var marks [3]int
	marks[0] = 10
	marks[1] = 20
	marks[2] = 30

	fmt.Println(marks)

	// array literals
	res := [5]string{"Kunal", "shubham", "raj", "prafull", "aniket"}

	fmt.Println(len(res))
}
