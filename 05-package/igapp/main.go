package main

import (
	"fmt"

	"github.com/bste101/igapp/time"
	"github.com/bste101/igapp/user"
)

func main() {
	user.Info("Best", "Cool", 21)
	t := time.Today()
	fmt.Printf("Today is %v\n", t)
}
