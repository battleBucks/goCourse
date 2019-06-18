package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go populate(c)
	expose(c)
	fmt.Printf("------------\nDONE!\n------------\n")
}

func populate(send chan<- int) {
	for i := 0; i < 100; i++ {
		send <- i
	}
	close(send)
}

func expose(receive <-chan int) {
	for v := range receive {
		fmt.Println(v)
	}
}
