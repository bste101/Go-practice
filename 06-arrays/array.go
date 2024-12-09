package main

import "fmt"

func main() {
	var langs = [4]string{"golang", "python", "javascript"}
	fmt.Printf("langs: %v\n", langs)

	n := langs[1]
	fmt.Printf("%#v\n", n)
}
