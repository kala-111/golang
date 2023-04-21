package main

import "fmt"

func main() {
	marks := []int{23, 43, 54, 25, 54, 56}
	sum, highest := f(marks)
	defer fmt.Println("sum:", sum, "highest:", highest)
	panic("Its an error")

	fmt.Println("highest:", highest)

}
func f(marks []int) (int, int) {
	sum := 0
	for _, v := range marks {
		sum += v
	}
	highest := marks[0]
	for _, v := range marks {
		if v > highest {
			highest = v
		}
	}
	return sum, highest
}
