package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("background: ", background)

	todo := context.TODO()
	fmt.Println("todo: ", todo)
}

func TestContextWithValue(t *testing.T) {
	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")

	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")

	ctxF := context.WithValue(ctxC, "f", "F")
	ctxG := context.WithValue(ctxF, "g", "G")

	fmt.Println("ctxA: ", ctxA)
	fmt.Println("ctxB: ", ctxB)
	fmt.Println("ctxC: ", ctxC)
	fmt.Println("ctxD: ", ctxD)
	fmt.Println("ctxE: ", ctxE)
	fmt.Println("ctxF: ", ctxF)
	fmt.Println("ctxG: ", ctxG)

	fmt.Println("ctxA: ", ctxA.Value("b"))

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				// simulasi kalau proses nya lama
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	cancel()

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println(n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}

func TestContextWithDealine(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println(n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}
