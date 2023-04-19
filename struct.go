package main

import "fmt"

type employee struct {
	firstname string
	lastname  string
	age       int
	salary    int
}

func main() {
	emp1 := employee{
		firstname: "ram",
		lastname:  "raj",
		age:       23,
		salary:    44000,
	}
	emp2 := employee{
		firstname: "gagan",
		lastname:  "deep",
		age:       22,
		salary:    45000,
	}
	fmt.Println("First name:", emp1.firstname)
	fmt.Println("Lastname:", emp1.lastname)
	fmt.Println("Age:", emp1.age)
	fmt.Println("Salary:", emp1.salary)
	fmt.Println("firstname:", emp2.firstname)
	fmt.Println("lastname:", emp2.lastname)
	fmt.Println("age:", emp2.age)
	fmt.Println("salary:", emp2.salary)
}
