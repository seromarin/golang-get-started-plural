package main

import "fmt"

func main() {
	firstName := "Sebastian"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Arthur"
	fmt.Println(ptr, *ptr)
}
