package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := "5"

	val, err := parseLevel(input)

	if err != nil {
		return err
	}

	fmt.Println("selected level is", val)

	return nil

}

func parseLevel(s string) (int, error) {
	//(value,error)
	// nil error -> success
	// not nil ->failure

	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("level must be number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("number must be 1 to 5")
	}

	return n, nil
}
