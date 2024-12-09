package main

import (
	"fmt"
)

var name string = "Best"

// func plus(a int, b int) int {
// 	return a + b
// }

var plus = func(a int, b int) int { return a + b }

func say(n string) {
	fmt.Printf("My name is %s\n", n)
}

func cal(op func(int, int) int) {
	r := op(4, 5)
	fmt.Printf("result : %v\n", r)
}
func main() {
	fmt.Printf("value: %v\n", name)
	fmt.Printf("type: %T\n", name)
	say(name)
	p := plus(1, 2)
	fmt.Println("1 + 2 =", p)
	fmt.Printf("type: %T\n", plus)

	cal(plus)
	minus := func(a int, b int) int { return a - b }
	cal(minus)
}
