package main

import "fmt"

func main() {
	name := "Mario"

	switch name {
	case "Mario":
		fmt.Println("Halo Mario")
	case "Budi":
		fmt.Println("Halo Budi")
	default:
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch with condition
	score := 90

	switch {
	case score > 90:
		fmt.Println("A")
	case score > 80:
		fmt.Println("B")
	case score > 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
