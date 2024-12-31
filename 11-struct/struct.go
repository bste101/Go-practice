package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Hello() string {
	return "Hello " + p.Name
}

func main() {
	x := Person{Name: "Best", Age: 21}
	fmt.Printf("Hello, %s\n", x.Name)

	// println((Hello(x)))
	println(x.Hello())
}
