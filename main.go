package main

import (
	"video/fib"

	"fmt"
)

func main() {

	result := fib.Fib(10)
	fmt.Println("Fib(10) =", result)

}
