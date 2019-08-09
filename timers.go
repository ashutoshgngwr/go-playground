package main

import "fmt"
import "time"

func main() {
	timer := time.NewTimer(time.Second * 2)
	lock := make(chan bool)
	
	go func() {
		<- timer.C
		fmt.Println("Timer finished")
		lock <- true
	}()

	<- lock
	if stopped := timer.Stop(); stopped {
		fmt.Println("Timer stopped")
	}
}
