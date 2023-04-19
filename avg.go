package main

import "fmt"
func main(){
	var x [5]int
	x[0] = 90
	x[1] = 67
	x[2] = 76
	x[3] = 80
	x[4] = 67

	total := 0
	for i := 0; i < 5; i++ {
		total += x[i]

		fmt.Println(total)
	}


func average() {
	marks := []int{34, 45, 35, 56, 66}
	sum := 0
	for _, v := range marks {
		sum += v
	fmt.Println("average",sum/len(marks))
}
}
}