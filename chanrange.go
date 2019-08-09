package main

import "fmt"
import "time"

func main() {
	msg := make(chan int, 1)
	go func() {
		for i := 0; i < 20; i++ {
			msg <- i
			time.Sleep(time.Millisecond * 200)
		}
		close(msg)
	}()

	for i := range msg {
		fmt.Println("msg", i)
	}
}
