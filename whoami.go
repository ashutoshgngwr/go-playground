package main

import "fmt"

func main() {
	whoami(1)
	whoami("some")
	whoami(2.2)
	whoami(false)
}

func whoami(i interface{}) {
	fmt.Print("whoami: ")
	switch t := i.(type) {
	default:
		fmt.Printf("%T", t)
	}

	fmt.Println()
}
