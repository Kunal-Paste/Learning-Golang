package main

import "fmt"

func divideAndsum(a int, b int) (Kunal int, john int) {
	Kunal = a / b
	john = a + b

	return Kunal, john
}

func main() {
	d, a := divideAndsum(10, 10)
	fmt.Println(d, a)
}
