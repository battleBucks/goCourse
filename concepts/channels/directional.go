package main

import "fmt"

func main() {

	//only a send channel designated by <-
	c := make(chan<- int, 2)
	//receive only:
	//c := make(<- chan int, 2)

	c <- 42
	c <- 43

	//fmt.Println(<-c)
	//fmt.Println(<-c)
	fmt.Println("--------------")
	fmt.Printf("%T\n", c)
}
