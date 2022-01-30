package main

import (
	"fmt"

	"github.com/motapratik/GolangTestingDemo/function"
)

func main() {
	fmt.Println(function.Greeting("Gopher"))
	fmt.Println(function.Addition(2, 2))
	fmt.Println(function.DisplayMsg("Hello"))
}
