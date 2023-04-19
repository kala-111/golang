package main

import "fmt"

func main() {
	fmt.Println(factorial(10))
}
func factorial(x int) int {
	if x == 1 || x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
