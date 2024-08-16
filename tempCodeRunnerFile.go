package main

import "fmt"

func main() {
	// array[low:high] => slice (low inclusive, high exclusive)
	// array[low:] => slice from low to end
	// array[:high] => slice from start to high
	// array[:] => slice all array
	names := [...]string{"Eko", "Kurniawan", "Khannedy"}
	slice := names[0:2]

	// slice[0] = "Budi"
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	day := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	slice1 := day[1:3]
	fmt.Println(slice1)

	slice2 := day[0:3]
	fmt.Println(slice2)

	// Append
	slice3 := append(slice2, "Hari")
	fmt.Println(slice3)

	// Make penjelasan lebih lanjut
	slice4 := make([]string, 2)
	slice4[0] = "Eko"
	slice4[1] = "Kurniawan"
	fmt.Println(slice4)

	slice4 = append(slice4, "Khannedy")

	fmt.Println(slice4)

	// Copy

	destination := make([]string, 2)
	source := []string{"Eko", "Kurniawan", "Khannedy"}

	n := copy(destination, source)
	fmt.Println(destination)

}
