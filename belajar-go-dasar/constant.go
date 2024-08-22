package main

import "fmt"

func main() {
	const firstName string = "John"
	const lastName = "Doe"
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)

	const (
		firstName1 string = "John"
		lastName1         = "Doe"
	)

	fmt.Printf("Hello, %s %s!\n", firstName1, lastName1)
}
