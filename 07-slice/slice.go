package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//         0  1  2  3  4
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 6)
	y := len(x)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	z := x[1:3]
	fmt.Printf("z: %v\n", z)
	name := "กขคง"
	println(utf8.RuneCountInString(name))
}
