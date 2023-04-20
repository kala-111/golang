package main

import "fmt"

func main() {
	fmt.Println(sum(1, 4, 3))

}
func sum(kala ...int) int {
	total := 0
	for _, v := range kala {
		total += v
	}
	return total
}
