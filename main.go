package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get() // ❶
	fmt.Println("1")
	instance := myPool.Get() // ❶
	fmt.Println("2")
	myPool.Put(instance) // ❷
	fmt.Println("3")
	myPool.Get() // ❸
}
