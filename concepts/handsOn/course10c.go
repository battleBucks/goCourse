package main

import (
	"fmt"
)

func main() {

	q := make(chan int)
	c := gen(q)
	recieve(c, q)

	fmt.Printf("-----------\n")
	fmt.Printf("about to exit...\n")

}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 100
	}()

	return c
}

func recieve(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Printf("Value %d is from 'c'\n", v)
		case v := <-q:
			fmt.Printf("Value %d is from 'q'\n", v)
			return
		}
	}
}
