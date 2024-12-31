package main

import "fmt"

func main() {
	me := "Golang"
	fmt.Println("You are %s\n", me)
	var x int
	x = 10

	var y *int
	y = &x
	fmt.Println(&x)
	fmt.Println(y)
	fmt.Println(*y)

	var addr *string = &me
	fmt.Println(addr)
}
