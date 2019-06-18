package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)

	check(err)

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	check(err)
	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	check(err)

	fmt.Println(answer1, answer2, answer3)
}
