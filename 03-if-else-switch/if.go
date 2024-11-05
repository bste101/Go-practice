package main

import "fmt"

func main() {
	num := 10

	if num%2 == 0 {
		fmt.Println("Even Number")
	} else if num == 31 {
		fmt.Println("ไม่บอก")
	} else {
		fmt.Println("Odd Number")
	}
}
