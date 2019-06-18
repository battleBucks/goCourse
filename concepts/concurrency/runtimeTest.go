package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	first()
	second()
}

func first() {
	for i := 0; i < 10; i++ {
		fmt.Println("first:", i)
	}
}

func second() {
	for i := 0; i < 10; i++ {
		fmt.Println("second:", i)
	}
}
