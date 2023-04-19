package main

import "fmt"

func main() {
	fmt.Println(0)
	fmt.Println(fibonaci(15))

}
func fibonaci(x int) int {
	if x <= 1 {
		return 1
	}
	return fibonaci(x-1) + fibonaci(x-2)
}
