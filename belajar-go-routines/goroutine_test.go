package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Test Done")

	time.Sleep(1 * time.Second)
}

func DisplayNumbe(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		go DisplayNumbe(i)
	}

	time.Sleep(5 * time.Second)
}
