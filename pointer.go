package main

import "fmt"

func main() {
	d := 3
	var e *int = &d
	fmt.Printf("type of e %T\n", e)
	fmt.Println("address of d is", e)
}
