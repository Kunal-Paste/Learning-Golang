package main

import (
	"fmt"
)

func main() {
	var city string
	city = "Pune"

	subscriber := 5000 // short form of declaring data type
	subscriber += 7000

	likes, comments := 1000, 30

	fmt.Println(city, subscriber, likes, comments)
}
