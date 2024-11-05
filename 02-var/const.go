package main

import (
	"fmt"
)

func main() {
	// ตัวอย่างค่าเปลี่ยนได้
	var day string = "Monday"
	fmt.Println("day:", day)

	day = "วันจันทร์"
	fmt.Println("day:", day)

	// constant
	const day_const string = "Monday"
	fmt.Println("day:", day_const)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
}
