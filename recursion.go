package main

import "fmt"

func factorial(n int64) int64 {
	if n < 1 {
		return 1
	}

	return factorial(n - 1) * n
}

func main() {
	fmt.Println("factorial(10):", factorial(10))
	fmt.Println("factorial(60):", factorial(60)) // :D
}
