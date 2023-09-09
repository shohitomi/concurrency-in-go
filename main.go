package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c) // ❶
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: // ❷
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
