package main

import "fmt"

type Aurthor struct {
	name   string
	age    int
	salary int
	book   string
	place  string
}
type Book struct {
	name  string
	pages int
}

func main() {
	var aur Aurthor
	var aur1 Aurthor
	var aur2 Aurthor

	var b1 Book
	var b2 Book
	var b3 Book

	aur.name = "valmiki"
	aur.age = 30
	aur.salary = 500000
	aur.book = "The Ramayan"
	aur.place = "varanasi"
	fmt.Println(aur)

	aur1 = Aurthor{"Ram", 24, 200000, "molla", "Bhadrachalam"}
	fmt.Println(aur1)

	aur2 = Aurthor{
		name:   "kali",
		age:    34,
		salary: 300000,
		book:   "Bharatham",
		place:  "Thirupathi",
	}
	fmt.Println(aur2)

	b1.name = "Bhagavatham"
	b1.pages = 1000

	fmt.Println(b1)

	b2 = Book{"khuran", 900}
	fmt.Println(b2)

	b3 = Book{
		name:  "Bible",
		pages: 1300,
	}
	fmt.Println(b3)

}
