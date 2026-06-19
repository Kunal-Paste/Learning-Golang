package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()

		fmt.Println("task 1")
		time.Sleep(250 * time.Millisecond)
		fmt.Println("task 1 is now done")

	}()

	go func() {
		defer wg.Done()

		fmt.Println("task 2")
		time.Sleep(150 * time.Millisecond)
		fmt.Println("task 2 is now done")

	}()

	go func() {
		defer wg.Done()

		fmt.Println("task 3")
		time.Sleep(199 * time.Millisecond)
		fmt.Println("task 3 is now done")

	}()

	fmt.Println("main waiting for tasks to execute")

	wg.Wait()

	fmt.Println("all tasks are finished !")
}
