package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Mario",
		"address": "Semarang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	book := make(map[string]string)

	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko Kurniawan Khannedy"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
}
