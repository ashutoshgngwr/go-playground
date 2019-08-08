package main

import "fmt"

func main() {
	var m = map[int]int {
		1: 10,
		10: 100,
		100: 1000,
		1000: 10000,
	}

	fmt.Println(m)
	delete(m, 10)
	fmt.Println("delete(m, 10) =", m)
	fmt.Println("m[10] = ", m[10])
	_, ok := m[10]
	fmt.Println("v, ok := m[10], ok =", ok)
}
