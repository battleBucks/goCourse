package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Printf("x is equal to %d before call\n", x)
	passIt(&x)
	fmt.Printf("x is equal to %d after call\n", x)
}
func passIt(y *int) {
	fmt.Printf("y is equal to %d upon entry to method\n", *y)
	*y = 42
	fmt.Printf("y is equal to %d before leaving method\n", *y)
}
