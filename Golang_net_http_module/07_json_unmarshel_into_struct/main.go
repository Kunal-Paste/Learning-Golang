package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type dataStructure struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	var data dataStructure

	bodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read body failed", err)
		return
	}

	if err := json.Unmarshal(bodyByte, &data); err != nil {
		fmt.Println("json unmarsh failed")
		return
	}

	fmt.Println(data.Fact, data.Length)

}
