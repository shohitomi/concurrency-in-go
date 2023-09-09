package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5) // ❶
		go func() {                       // ❷
			defer close(resultStream) // ❸
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream // ❹
	}

	resultStream := chanOwner()
	for result := range resultStream { // ❺
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
