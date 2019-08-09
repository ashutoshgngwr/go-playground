package main

import "fmt"

func sequence(start int) func() int {
	return func() int {
		defer func() { start++ }()
		return start
	} 
}

func main() {
	next := sequence(10)

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
