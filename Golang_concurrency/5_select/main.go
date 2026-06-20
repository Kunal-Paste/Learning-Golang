package main

import (
	"fmt"
	"time"
)

func main() {
	resultCh := make(chan string)

	go func() {
		time.Sleep(40 * time.Millisecond)
		resultCh <- "worker success"
	}()

	// timeout channel
	timeCh := time.After(200 * time.Millisecond)

	select {
	case res := <-resultCh:
		fmt.Println("main: go result", res)

	case <-timeCh:
		fmt.Println("main : timeout happned stop waiting")
	}

	fmt.Println("work is now done")

}
