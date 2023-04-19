package main

import "fmt"

type persion struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 persion
	var pers2 persion

	pers1.name = "kala"
	pers1.age = 25
	pers1.job = "enganeer"
	pers1.salary = 100000

	pers2.name = "madhu"
	pers2.age = 30
	pers2.job = "principal"
	pers2.salary = 200000

	printPersion(pers1)
	printPersion(pers2)
}
func printPersion(pers persion) {
	fmt.Println("name:", pers.name)
	fmt.Println("age:", pers.age)
	fmt.Println("job:", pers.job)
	fmt.Println("salary:", pers.salary)
}
