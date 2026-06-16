package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// goroutine A

	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("goroutine A finished the simulated api at : ", time.Since(start))
	}()

	// goroutine B
	go func() {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("goroutine B finished the simulated api at : ", time.Since(start))
	}()

	fmt.Println("main func started two goroutine at time : ", time.Since(start))

	fmt.Println("main doing step 1 :", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main doing step 2 :", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main doing step 3 :", time.Since(start))

	time.Sleep(500 * time.Millisecond)

	fmt.Println("main exiting at : ", time.Since(start))
}
