package main

import "fmt"

func init() {
	fmt.Println("init()")
}

func beforeInit() int {
	fmt.Println("beforeInit()")
	return 5
}

var x = beforeInit()

func main() {
	fmt.Println("main()")
}