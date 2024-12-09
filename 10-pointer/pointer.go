package main

import "fmt"

func main() {
	me := "Golang"
	fmt.Println("You are %s\n", me)

	var addr *string= &me
	fmt.Println(addr)
}
