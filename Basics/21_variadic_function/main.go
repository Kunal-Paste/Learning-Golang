package main

import "fmt"

func addition(num ...int) int {

	total := 0

	for _, currentvalue := range num {
		total += currentvalue
	}

	return total
}

func main() {
	ans := addition(50 + 50 + 50 + 50)
	fmt.Println(ans)

	value := []int{1, 2, 3, 4, 5}
	fmt.Println(addition(value...))

	res := func(a int) int {
		return a * 2
	}

	fmt.Println(res(2))

	add := func(a int, b int) int { //iife
		return a + b
	}(10, 20)

	fmt.Println(add)

}
