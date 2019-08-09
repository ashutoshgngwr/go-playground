package main

import "fmt"

func acceptsVarArgs(args ...int) {
	fmt.Println(args)
}

func main() {
	acceptsVarArgs(9, 4, 5, 6, 6, 8, 8, 6, 7)
}