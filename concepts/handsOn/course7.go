package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).age = (*p).age + 1
}

func main() {
	i := 42
	fmt.Printf("Value is %d\n", i)
	fmt.Printf("Address is %v\n", &i)

	p := person{
		first: "Joe",
		last:  "Barrows",
		age:   39,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
} //end main
