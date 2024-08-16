package main

import "fmt"

func main() {
	// konversi int
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi uint
	var nilaiU32 uint32 = 100000
	var nilaiU64 uint64 = uint64(nilaiU32)
	var nilaiU8 uint8 = uint8(nilaiU32)
	fmt.Println(nilaiU32)
	fmt.Println(nilaiU64)
	fmt.Println(nilaiU8)

	var name = "John Wick"
	var e byte = name[0]
	var eString string = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
	
}

