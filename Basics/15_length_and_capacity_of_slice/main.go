package main

import (
	"fmt"
)

func main() {
	result := make([]int, 0, 5)
	fmt.Println(result, len(result), cap(result))

	result = append(result, 10)
	fmt.Println("after appending 10", result, len(result), cap(result))

	result = append(result, 20, 30)
	fmt.Println("after appending 20, 30", result, len(result), cap(result))

	result = append(result, 40, 50, 60)
	fmt.Println("after appending 40, 50, 60", result, len(result), cap(result))

	result = append(result, 70, 80)
	fmt.Println("after appending 70, 80", result, len(result), cap(result))

	result = append(result, 90, 100, 110)
	fmt.Println("after appending 90, 100, 110", result, len(result), cap(result))

	//so cap is like a dynamic array which gets doubled if the elements exceds the length

	todo := []string{"to youtube", "learn golang"}

	moreTasks := []string{"learn japanese"}

	todo = append(todo, moreTasks...)

	fmt.Println(todo)
}
