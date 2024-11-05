package main

import (
	"fmt"
)

/*
	The zero value is:
	0 for numeric types,
	false for the boolean type
	"" (the empty string) for strings.
*/

/* basic type
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

// var name string = "Golang มาแล้ว!!!"

func main() {
	// var name string = "Golang มาแล้ว!!!"
	// var name = "Golang มาแล้ว!!!"
	name := "Golang มาแล้ว!!!" // ใช้นอกฟังก์ชั่นไม่ได้

	fmt.Printf("name: %v\n", name)
	fmt.Printf("type: %T\n", name)
}
