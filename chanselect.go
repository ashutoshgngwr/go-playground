package main

import (
	"fmt"
	"time"
)


func main() {
	tik := make(chan string, 1)
	tak := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 7)
		tik <- "tik tik tik"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		tak <- "tak tak tak"
	}()

	fmt.Println("ta ta ta")

	for i := 0; i < 10; i++ {
		select {
		case msg := <- tik:
			fmt.Println(msg)
		case msg := <- tak:
			fmt.Println(msg)
		case <- time.After(1 * time.Second):
			fmt.Println("timed out...")
		}
	}
}
