package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
		fmt.Println("Test Done")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(3 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

// only in digunakan untuk mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"
}

// only out digunakan untuk menerima data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Hello"
	channel <- "World"
	channel <- "Go"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Buffered Channel")
}

func TestGoBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Hello"
		channel <- "World"
		channel <- "Go"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Buffered Channel")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Data %d", i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Range Channel")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
