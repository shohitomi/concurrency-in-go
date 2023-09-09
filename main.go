package main

import (
	"fmt"
	"time"
)

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream) // ❶
		for i := 1; i <= 5; i++ {
			intStream <- i
			time.Sleep(1 * time.Second)
		}
	}()

	for integer := range intStream { // ❷
		fmt.Printf("%v ", integer)
	}

	fmt.Println("done")
}
