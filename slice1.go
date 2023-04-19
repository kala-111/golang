package main

import "fmt"

func main() {
	StudentAge := map[int]string{
		19: "raj",
		14: "vani",
		16: "rupa",
		10: "vidhya",
	}
	fmt.Println(StudentAge)
	for Age, Sname := range StudentAge {
		if Age >= 12 {

			fmt.Println(Sname)
		}

	}
}
