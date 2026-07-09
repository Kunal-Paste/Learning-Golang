package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Value string
	Err   error
}

func Worker(url string, wg *sync.WaitGroup, resultChan chan Data) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 50)

	fmt.Printf("image processed : %s\n", url)

	resultChan <- Data{
		Value: url,
		Err:   nil,
	}
}

func main() {
	var wg sync.WaitGroup

	resultChan := make(chan Data, 5)

	startTime := time.Now()

	wg.Add(5)
	go Worker("image1.png_url", &wg, resultChan)
	go Worker("image2.png_url", &wg, resultChan)
	go Worker("image3.png_url", &wg, resultChan)
	go Worker("image4.png_url", &wg, resultChan)
	go Worker("image5.png_url", &wg, resultChan)

	wg.Wait()

	close(resultChan)
	for result := range resultChan {
		fmt.Printf("received : %v\n", result)
	}

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
