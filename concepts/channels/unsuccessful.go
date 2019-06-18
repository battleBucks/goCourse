package main

import "fmt"

func main() {
	//makes channel and sets buffer channel to allow N values to sit in there
	//to run, change value of second arg
	c := make(chan int, 1)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
