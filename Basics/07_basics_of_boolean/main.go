package main

import (
	"fmt"
)

func main() {
	isLogedin := true
	isAdmin := false
	hasSubscription := true

	canHandledashboard := isAdmin && hasSubscription
	canDeletepost := isAdmin || (isLogedin && hasSubscription)

	fmt.Println(canDeletepost, canHandledashboard)

	age := 23
	isAdult := age > 18
	fmt.Println(isAdult)
}
