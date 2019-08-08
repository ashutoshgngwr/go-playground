package main

import "fmt"

func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println("2 + 3 =", add(2, 3))
}
