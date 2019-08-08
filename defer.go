package main

import "fmt"

// defer LIFOs
func main() {
	defer fmt.Println("Deferred output 0")

	fmt.Println("Normal output 1")

	for i := 1; i < 5; i++ {
		defer fmt.Println("Deferred output", i)
	}

	fmt.Println("Normal output 2")
}
