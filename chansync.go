package main

import (
	"fmt"
	"time"
)

func runner(done chan bool) {
	fmt.Println("runner: running...")
	time.Sleep(time.Second * 5)
	fmt.Println("runner: finished!")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	fmt.Println("main: starting runner...")
	go runner(done)

	fmt.Println("main: waiting for runner to exit")
	<- done
	fmt.Println("main: exiting...")
}
