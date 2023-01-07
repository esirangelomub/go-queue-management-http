package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)

	// Start a goroutine to process items in the queue
	go func() {
		for {
			item := <-queue
			fmt.Println("Processing item:", item)
		}
	}()

	// Add items to the queue every 10 seconds
	for i := 0; i < 10; i++ {
		queue <- i
		time.Sleep(time.Second * 10)
	}
}
