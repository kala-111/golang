package main

import "fmt"

func main() {
	arr1 := []int{4, 5, 7, 8, 2, 1, 9}
	myslice1 := arr1[1:4]
	fmt.Printf("myslice1 =%v\n", myslice1)
	fmt.Printf("length =%v\n", len(myslice1))
	fmt.Printf("capacity =%v\n", cap(myslice1))
	myslice2 := make([]int, 5, 7)
	fmt.Printf("myslice2 =%v\n", myslice2)

	fmt.Printf("length =%v\n", len(myslice2))
	fmt.Printf("capacity =%v\n", cap(myslice2))
	myslice2 = append(myslice2, 55, 22, 44)
	fmt.Printf("myslice2 =%v\n", myslice2)
	fmt.Printf("length =%v\n", len(myslice2))
	fmt.Printf("capacity =%v\n", cap(myslice2))
	myslice3 := []int{4, 5, 44, 85, 77, 80}
	fmt.Printf("myslice3 =%v\n", myslice3)
	myslice4 := []int{11, 33, 44, 55, 78}
	myslice5 := append(myslice3, myslice4...)
	fmt.Printf("myslice5 =%v\n", myslice5)
	arr2 := []int{11, 22, 33, 44, 55, 66, 77}
	myslice7 := arr2[2:5]
	fmt.Printf("myslice7 =%v\n", myslice7)
	myslice6 := make([]int, 34, 43)
	fmt.Printf("myslice6 =%v\n", myslice6)
	fmt.Printf("length =%v\n", len(myslice6))
	fmt.Printf("capacity =%v\n", cap(myslice6))
	myslice6 = append(myslice6, 65, 34)
	fmt.Printf("myslice6 =%v\n", myslice6)

	fmt.Printf("length =%v\n", len(myslice6))
	fmt.Printf("capacity =%v\n", cap(myslice6))

}
