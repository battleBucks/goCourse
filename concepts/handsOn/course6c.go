package main

import (
	"fmt"
)

func main() {
	g := func(x string) {
		fmt.Println("Just kicking out some output with ", x)
	}
	call(g, "stringy")

	inc := incrementer()
	for i := 0; i < 3; i++ {
		fmt.Println(inc())
	}
}

func out(x string) {
	fmt.Println("Just kicking out some output with ", x)
}

func call(f func(y string), s string) {
	fmt.Println("\tCalling out in a minute with ", s, "as a param")
	f(s)
}
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
