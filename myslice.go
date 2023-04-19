package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %v\n", len(numbers))
	fmt.Printf("capacity = %v\n", cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]
	fmt.Printf("neededNumbers %v\n", neededNumbers)
	numbersCopy := make([]int, len(neededNumbers))

	fmt.Printf("numberscopy =%v\n", numbersCopy)
	fmt.Printf("capacity =%v\n", cap(numbersCopy))
	fmt.Printf("capacity= %v\n", cap(numbersCopy))
}
