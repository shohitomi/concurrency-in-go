package main

import (
	"fmt"
	"time"
)

func main() {
	stringStream := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		stringStream <- "Hello channels!" // ❶
	}()
	fmt.Println(<-stringStream) //
	fmt.Println("Hello")
}
