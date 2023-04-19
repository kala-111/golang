package main

import "fmt"

func main() {
	time := 6
	if time < 4 {
		fmt.Println("good morning")
		if time > 6 {
			fmt.Println("good evening")
			if time >= 5 {
				fmt.Println("good")
			}
		}
	} else {
		fmt.Println("good afternoon")
	}
}
