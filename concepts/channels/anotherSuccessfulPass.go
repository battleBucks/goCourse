package main

import "fmt"

func main() {
	//makes channel and sets buffer channel to allow N values to sit in there
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
