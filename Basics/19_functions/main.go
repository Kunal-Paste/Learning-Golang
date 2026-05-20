package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func productAndanswer(a int, b int) (int, string) {
	sum := a * b
	answer := "your answer is correct"

	return sum, answer
}

func main() {
	addition := add(10, 20)

	fmt.Println(addition)

	logic, ans := productAndanswer(10, 20)

	onlyLogic, _ := productAndanswer(20, 30)

	fmt.Println(logic, ans)
	fmt.Println(onlyLogic)
}
