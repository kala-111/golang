package main

import "fmt"

func main() {
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("odd weekday")
	case 2, 4:
		fmt.Println("even weekday")
	case 7, 8:
		fmt.Println("weekend")
	default:
		fmt.Println("invalid day")

	}
}
