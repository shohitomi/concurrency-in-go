package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // ❶
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	time.Sleep(3 * time.Second)
	close(begin) // ❷
	wg.Wait()
}
