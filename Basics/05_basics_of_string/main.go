package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Kunal"
	lastName := "Paste"
	fullName := firstName + " " + lastName

	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))

}
