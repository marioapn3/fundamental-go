package main

import "fmt"

type Address struct {
	Street string
	City   string
}

func main() {
	adress1 := Address{"Jalan A", "Jakarta"}
	adress2 := adress1

	adress2.City = "Bandung"

	fmt.Println(adress1)
	fmt.Println(adress2)

	// opeartor &
	adress3 := Address{"Jalan B", "Jakarta"}
	adress4 := &adress3

	adress4.City = "Bandung"

	fmt.Println(adress3)
	fmt.Println(adress4)
}
