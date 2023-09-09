package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // ❶
	go func() {
		defer wg.Done() // ❷
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1) // ❶
	go func() {
		defer wg.Done() // ❷
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // ❸
	fmt.Println("All goroutines complete.")
}
