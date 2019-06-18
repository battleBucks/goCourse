package main

import "fmt"

func main() {

	c := make(chan int)

	//send
	go sendIt(c)

	//receive
	getIt(c)

	fmt.Println("about to exit")
}

//send
func sendIt(c chan<- int) {
	c <- 42
}

//receive
func getIt(c <-chan int) {
	fmt.Println(<-c)
}
