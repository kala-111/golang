package main

import "fmt"

func main() {
	var i = "rama"
	var j = "chandra"
	fmt.Println(i, "and", j, "are good friend")
	fmt.Printf("%v and %v are classmates \n", i, j)
	fmt.Printf("%T and %v \n", i, i)
	fmt.Printf("%T and %v", j, j)
}
