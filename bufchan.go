package main

import "fmt"

func main() {
	sig := make(chan int, 10)

	for i := 0; i < 10; i++ {
		sig <- i
	}

	fmt.Println(<- sig)
	fmt.Println(<- sig)
}
