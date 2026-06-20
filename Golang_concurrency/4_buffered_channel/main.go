package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan string, 2)

	go func() {
		fmt.Println("producer: sending job 1")
		jobs <- "job-1"

		fmt.Println("producer: sending job 2")
		jobs <- "job-2"

		fmt.Println("producer: sending job 3, but this will wait until consumer reads")
		jobs <- "job-3"

		fmt.Println("producer: sent all jobs")
		close(jobs)
	}()

	for jobs := range jobs {
		fmt.Println("consumer got ", jobs)
		time.Sleep(300 * time.Millisecond)
		fmt.Println("consumer finished", jobs)
	}

	fmt.Println("main: all jobs are completed")
}
