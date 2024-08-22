package main

import "fmt"

func main() {
	var firstName string = "John"
	// membuat variabel lastName tanpa tipe data
	lastName := "Doe"
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)

	var (
		firstName1 string = "John"
		lastName1  string = "Doe"
	)

	fmt.Printf("Hello, %s %s!\n", firstName1, lastName1)
}
