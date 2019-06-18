package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println("Value: ", a)
	fmt.Println("Address in memory: ", &a)
	fmt.Printf("Value type: %T\n", a)
	fmt.Printf("Address in memory type: %T\n", &a)

	fmt.Println("assigning pointer a to pointer b")
	var b *int = &a
	fmt.Println("Value of b: ", b)
	fmt.Println("Value of b with *: ", *b)
	fmt.Println("updating value in b")
	*b = 12
	fmt.Println("printing value in a")
	fmt.Println("Value of a: ", a)
}
