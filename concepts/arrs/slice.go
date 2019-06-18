package main

import (
	"fmt"
)

func main() {
	//Composite literal:
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	//everything from one on
	fmt.Println(x[1:])
	//from one(inclusive) to 4(not inclusive)
	fmt.Println(x[1:4])

	//append
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)
	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)
	//removing two elements at positions 2 and 3
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
