package main

import "fmt"

func main() {
	marks := map[string]int{
		"Kunal":   20,
		"Shubham": 0, // valid zero
	}

	fmt.Println(marks["Kunal"])
	fmt.Println(marks["Shubham"])
	fmt.Println(marks["PP"]) // iit has not a valid zero value;

	valS, okS := marks["Shubham"]
	valP, okP := marks["PP"]

	fmt.Println(valS, okS)
	fmt.Println(valP, okP)

	if val, validator := marks["PP"]; validator {
		fmt.Println(val, "value is present")
	} else {
		fmt.Println("value is not present")
	}

	if val, i := marks["Shubham"]; i {
		fmt.Println(val, "value is present")
	} else {
		fmt.Println("value is not present")
	}
}
