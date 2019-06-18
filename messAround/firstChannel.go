package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int, 2)

	go func() {
		ch1 <- 42
		ch1 <- 45
		ch1 <- 48
	}()
	ch2 <- 43
	ch2 <- 44

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
