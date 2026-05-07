package main

import "fmt"

func main() {
	result := []string{"Kunal", "Paste"}
	fmt.Println(result, result[0], result[len(result)-1])

	var num []int
	num = append(num, 10)
	num = append(num, 20)
	num = append(num, 30)

	fmt.Println(num)
}
