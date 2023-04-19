package main

import "fmt"

func main() {
	months := 12

	switch months {
	case 1:
		fmt.Println("january")
	case 2:
		fmt.Println("february")
	case 3:
		fmt.Println("march")
	case 4:
		fmt.Println("april")
	case 5:
		fmt.Println("may")
	case 6:
		fmt.Println("june")
	case 7:
		fmt.Println("july")
	case 8:
		fmt.Println("august")
	case 9:
		fmt.Println("september")
	case 10:
		fmt.Println("october")
	case 11:
		fmt.Println("november")
	case 12:
		fmt.Println("december")

	default:
		fmt.Println("something else")
	}
	if months == 5 {
		fmt.Println("months")
	} else {
		fmt.Println("something esle")
	}
}
