package main

import "fmt"

func main() {
	StuMarks := make(map[string]int)
	StuMarks["Alice"] = 90
	StuMarks["Balu"] = 87
	StuMarks["Charle"] = 88

	Highest := StuMarks["Balu"]

	for SName, Marks := range StuMarks {
		if Marks == Highest {
			fmt.Println("the student got Highest marks", SName, Marks)
		}
	}

}
