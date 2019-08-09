package main

import (
	"time"
	"fmt"
)

func createWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Exec #%d\n", j)
		time.Sleep(time.Second)
		fmt.Printf("Finish #%d\n", j)
		results <- j
	}
	close(results)
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < 5; i++ {
		go createWorker(i, jobs, results)
	}

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for range results {	}
}
