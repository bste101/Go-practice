package main

import "fmt"

func info(name string) {
	// name := "Golang"
	fmt.Printf("My name is %s\n", name)
}
func introduce(name string, age int, bio string) {
	fmt.Printf("Hello, My name is %s.I'm %d years old,Now I'm %s", name, age, bio)
}

func main() {
	info("Golang")
	info("Best")
	introduce("Best", 21, "Developer")
}
