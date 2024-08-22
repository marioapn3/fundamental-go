package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func parameter(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func returnFunc(firstName string, lastName string) string {
	return "Hello " + firstName + " " + lastName
}

func multipleReturn(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func namedReturn(firstName string, lastName string) (first string, last string) {
	first = firstName
	last = lastName
	return first, last
}

// variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	sayHello()
	parameter("Eko", "Khannedy")
	data := returnFunc("Budi", "Nugraha")
	fmt.Println(data)

	firstName, lastName := multipleReturn("Budi", "Nugraha")
	fmt.Println(firstName, lastName)

	// jika kita hanya ingin mengambil salah satu return value maka kita bisa menggunakan _ (underscore)
	fn1, _ := multipleReturn("Budi", "Nugraha")
	fmt.Println(fn1)

	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)
}
