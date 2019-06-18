package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// like below or add go routine
	//c:= make(chan int, 1)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
