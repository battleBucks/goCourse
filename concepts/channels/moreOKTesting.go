package main

import "fmt"

func main() {

	c := make(chan string)

	go func() {
		c <- "Joe"
		close(c)
	}()

	v, ok := <-c
	fmt.Println("v:", v, "\tok:", ok)
	v, ok = <-c
	fmt.Println("v:", v, "\tok:", ok)
}
