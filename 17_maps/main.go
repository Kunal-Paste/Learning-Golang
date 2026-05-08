package main

import "fmt"

func main() {
	data := map[string]int{
		"Kunal":   20000,
		"Shubham": 20000,
	}

	fmt.Println(data)

	var score map[string]int

	fmt.Println(score, score["a"])

	score = make(map[string]int)
	score["marks"] = 90

	fmt.Println(score)

	users := map[string]string{
		"Kunal":   "Pune",
		"Shubham": "Mumbai",
		"Aniket":  "Satara",
	}

	fmt.Println(users)
	delete(users, "Aniket")
	delete(users, "PP") // no error throw
	fmt.Println(users)
}
