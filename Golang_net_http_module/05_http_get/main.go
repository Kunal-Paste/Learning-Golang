package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	fmt.Println("status : ", res.Status)
	fmt.Println("statuscode : ", res.StatusCode)
}
