package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// struct method
func (p Person) sayHello() {
	fmt.Println("Hello", p.Name)
}

func main() {
	var eko Person
	eko.Name = "Eko"
	eko.Age = 30

	joko := Person{
		Name: "Joko",
		Age:  35,
	}

	fmt.Println(eko)
	fmt.Println(joko)
}
