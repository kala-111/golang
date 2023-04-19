package main

import "fmt"

func main() {
	var arr1 = [8]string{"hello", "world", "welcome"}
	var arr2 = []int{3, 5, 6, 7}
	var arr3 = []float32{2.3, 4.5}
	var arr4 = []bool{}
	var arr5 = []int{4, 6, 3, 2, 8}
	var arr6 = [6]int{2, 4, 5}
	arr7 := [5]int{0: 20, 3: 40, 4: 50}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5[3])
	arr6[2] = 40
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(len(arr1))
	fmt.Println(len(arr4))
	fmt.Println(len(arr3))
	fmt.Println(len(arr5))

}
