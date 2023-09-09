package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})    // ❶
	queue := make([]interface{}, 0, 10) // ❷

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // ❽
		queue = queue[1:] // ❾
		fmt.Println("Removed from queue")
		c.L.Unlock() // ❿
		c.Signal()   // ⓫
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            // ❸
		for len(queue) == 2 { // ❹
			c.Wait() // ❺
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // ❻
		c.L.Unlock()                        // ❼
	}
}
