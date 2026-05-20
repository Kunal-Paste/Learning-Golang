package main

import (
	"fmt"
)

func main() {
	view1 := 1000
	view2 := 2000
	totalview := view1 + view2

	likes := 10
	likes++
	likes++

	averageview := totalview / 2

	fmt.Println(totalview, likes, averageview)

	rating1 := 4.5
	rating2 := 4.8

	averagerating := (rating1 + rating2) / 2

	fmt.Println(averagerating)
}
