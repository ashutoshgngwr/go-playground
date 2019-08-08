package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("swap(2, 3) =", swap(2, 3))
}
