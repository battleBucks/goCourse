package main

import (
	"fmt"
)

type person struct {
	first  string
	last   string
	gender string
	age    int
}

func main() {

	g1 := person{
		"Grayson",
		"Lane",
		"Male",
		8,
	}
	g1.info()

	func(x int) {
		fmt.Println("This is my anonymous function passing in an age:", x)
	}(g1.age)
}

func (p person) info() {
	fmt.Printf("Name: %s, %s\nGender: %s\nAge: %d\n", p.last, p.first, p.gender, p.age)
}
