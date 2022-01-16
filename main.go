package main

import (
	"fmt"

	"github.com/wederbrand/greeting"
)

func main() {
	greeting := greeting.Greeting()
	fmt.Println("hello", greeting)
}