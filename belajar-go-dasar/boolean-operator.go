package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian > 75
	var lulusAbsensi = absensi > 75

	var lulus = lulusUjian && lulusAbsensi

	fmt.Println(lulus)

	var nilaiUjian = 80
	var nilaiAbsensi = 80

	var lulusUjian2 = nilaiUjian > 75

	var lulusAbsensi2 = nilaiAbsensi > 75

	var lulus2 = lulusUjian2 || lulusAbsensi2

	fmt.Println(lulus2)

}
