package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			fmt.Println(i, "strating...")
			time.Sleep(time.Second * 5)
			fmt.Println(i, "finishing...")
			wg.Done()
		}(i, &wg)
	}

	fmt.Println("waiting for all workers to finish...")
	wg.Wait()

	fmt.Println("all workers finished!")
}
