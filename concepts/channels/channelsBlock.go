package main

import "fmt"

func main() {
	//makes channel
	c := make(chan int)

	//tries to put 42 on but it locks - needs to pass batton
	//(send and receive need to happen at same time)
	//c <- 42
	//fmt.Println(<-c)

	//here down works, to test block uncomment above line and comment below
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
