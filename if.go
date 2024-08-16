package main

import "fmt"

func main() {
	name := "Mario"

	if name == "Mario" {
		fmt.Println("Halo Mario")
	} else if name == "Budi" {
		fmt.Println("Halo BudiBudi")
	} else {
		fmt.Println("Halo Tidak Diketahui")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
