package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() {
		for i := 0; i < 10000000000; i++ {}
		msg <- "DONE!"
	}()

	fmt.Println("waiting to receive message...")
	fmt.Println(<- msg)
}
