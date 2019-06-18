package main

import "fmt"

func main() {
	// make(type, length, capacity)
	x := make([]int, 10, 11)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

//makes size 11
	x = append(x, 3424)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
