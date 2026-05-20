package main

import "fmt"

func main() {
	score := 10
	fmt.Println("before score : ", score)

	addScore(&score)
	fmt.Println("after updation : ", score)
}

func addScore(score *int) {
	*score += 10
}
