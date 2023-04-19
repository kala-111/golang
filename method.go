package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary int
}

func (e Employee) displaysalary() {
	fmt.Printf("Salary of %s is %d", e.name, e.salary)
}
func main() {
	emp1 := Employee{
		name:   "ram",
		age:    22,
		salary: 44000,
	}
	emp1.displaysalary()
}
