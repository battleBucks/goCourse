package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)
}

func receive(even, odd, q <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Value is even:", v)
		case v := <-odd:
			fmt.Println("Value is odd:", v)
		case v := <-q:
			fmt.Println("quit channel:", v)
			return
		}
	}
}
func send(even, odd, q chan<- int) {

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	//close(even)
	//close(odd)
	q <- 0
}
