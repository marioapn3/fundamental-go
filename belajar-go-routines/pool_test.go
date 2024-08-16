package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoolTest(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Hello")
	pool.Put("World")

	for i := 0; i < 10; i++ {
		go func() {
			value := pool.Get()
			fmt.Println("Value: ", value)
			pool.Put(value)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
