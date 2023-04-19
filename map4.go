package main

import "fmt"

func main() {

	fmt.Println(evenNumbers(50))

}
func evenNumbers(x int) int {

	even := []int{}
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			even = append(even, i)
		}

	}
	return even

}
