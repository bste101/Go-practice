package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday

	switch today {
	case time.Monday:
		fmt.Println("today is Monday")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Printf("สวัสดีวัน %v!!!\n", today)
	}
}
