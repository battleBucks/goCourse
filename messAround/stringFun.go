package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go sendIt(ch)

	receiveIt(ch)

}

func sendIt(ch chan<- string) {

	for i := 0; i < 10; i++ {
		ch <- string('A' + i)
	}
	close(ch)
}

func receiveIt(ch <-chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}
