package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)
}

func sum(intArr ...int) int {
	total := 0
	for _, v := range intArr {
		total += v
	}
	return total
}

// even takes in function that returns an int as a parameter.
//That function takes in a slice of ints as a param
func even(f func(intArr ...int) int, intList ...int) int {
	var y []int
	for _, v := range intList {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	//call function and return it
	return f(y...)
}
