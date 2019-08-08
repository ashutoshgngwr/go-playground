package main

import "fmt"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c ", i)

		// no need to break after each case
		switch i {
		case 'a', 'e', 'i', 'o', 'u':
			fmt.Print("v")
		default:
			fmt.Print("c")
		}

		fmt.Println()
	}
}
