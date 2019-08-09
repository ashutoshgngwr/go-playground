package main

import "fmt"
import "time"

func main() {
	steve := make(chan int, 1)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("Sending job", i, "...")
			time.Sleep(time.Millisecond * 300)
			steve <- i
		}
		close(steve)
	}()

	for {
		j, more := <- steve
		if more {
			fmt.Println("Received job", j, ".")
		} else {
			fmt.Println("No more jobs! exiting...")
			break
		}
	}
}
