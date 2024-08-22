package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// for dengan statement
	for i := 0; i < 10; i++ {
		fmt.Println("Angka", i)
	}

	name := []string{"Eko", "Kurniawan", "Khannedy"}
	for index, name := range name {
		fmt.Println("Index", index, "=", name)
	}
}
