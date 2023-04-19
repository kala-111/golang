package main

import "fmt"

func main() {
	marks := map[string]int{
		"roopa": 90,
		"megha": 35,
		"mahe":  60,
		"krish": 100,
	}
	highestmarks := marks["roopa"]
	var studentname string
	for name, marks := range marks {
		if marks > highestmarks {
			highestmarks = marks
			studentname = name
		}
	}
	fmt.Println(studentname, marks)
}
