package main

import "fmt"

func main() {
	var name = "kala"

	var ptr *string

	ptr = &name
	fmt.Println(*ptr)
	fmt.Println(&ptr)
}
