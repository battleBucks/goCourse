package main

import "fmt"

func main() {

	c := make(chan int)

	//send
	go sendIt(c)

	//close(c)
	//receive
	getIt(c)

	fmt.Println("about to exit")
}

//send
func sendIt(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

//receive
func getIt(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
