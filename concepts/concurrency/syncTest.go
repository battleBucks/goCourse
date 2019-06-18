package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitForIt sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	waitForIt.Add(1) // for adding to wait for one go routine
	go first()       // creating a go routine that shoots off at this point
	second()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	waitForIt.Wait() // waiting for go routine to finish
}

func first() {
	for i := 0; i < 10; i++ {
		fmt.Println("first:", i)
	}
	waitForIt.Done() // routine is finished
}

func second() {
	for i := 0; i < 10; i++ {
		fmt.Println("second:", i)
	}
}
