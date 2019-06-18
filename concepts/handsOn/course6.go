package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi, my name is poopypants.... I mean %s %s and I am %d years old\n", p.first, p.last, p.age)
}
func (p person) nameAndAge() (string, int) {
	return p.first, p.age
}
func (p person) years() int {
	return p.age
}
func variadicExercise(nums ...int) int {
	sum := arrExercise(nums)
	return sum
}
func arrExercise(nums []int) int {
	sum := 0
	for _, x := range nums {
		sum = sum + x
	}
	return sum
}
func toBeDeferred(str string) {
	fmt.Println(str)
}
func returningFunc() func() {
	return func() {
		fmt.Println("I'm deep in a function")
	}
}

func main() {

	defer toBeDeferred("is this at the end?\n")
	p := person{
		"grayson",
		"lane",
		8,
	}
	fmt.Printf("%v\n", p)

	name, years := p.nameAndAge()
	justAge := p.years()
	if years == justAge {
		fmt.Println("printing person details:")
		fmt.Println(name)
		fmt.Println(years)
	}

	fmt.Println("\nAttempting to unfurl and sum stuff")
	arr := []int{1, 2, 3, 4, 5}
	sum := variadicExercise(arr...)
	fmt.Printf("Expecting 15 here: %d\n", sum)
	arr = append(arr, 6)
	sum = arrExercise(arr)
	fmt.Printf("Expecting 21 here: %d\n", sum)

	fmt.Println("Calling speak method now...")
	p.speak()

	funk := returningFunc()
	fmt.Println("about to call the func that was assigned to a variable")
	funk()
}
