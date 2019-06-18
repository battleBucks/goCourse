package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)
}

func receive(even, odd <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("Value is even:", v)
		case v := <-odd:
			fmt.Println("Value is odd:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("NOT OK:", i, "OK VAL", ok)
			} else {
				fmt.Println("NOT OK:", i, "OK VAL", ok)
			}
			return
		}
	}
}
func send(even, odd chan<- int, q chan<- bool) {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	//close(even)
	//close(odd)
	close(q)
}
