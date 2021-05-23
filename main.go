package main

import (
	"fmt"
)

const (
	first  = iota + 6
	second = 2 << iota
)

func main() {
	fmt.Println("Hello, my dear friend")
	fmt.Println(first, second)
}
