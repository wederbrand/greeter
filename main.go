package main

import (
	"fmt"

	"github.com/wederbrand/greeting"
)

func main() {
	greeting := greeting.greeting()
	fmt.Println("hello", greeting)
}