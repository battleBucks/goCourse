package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	var wg sync.WaitGroup

        fmt.Println(runtime.NumGoroutine())
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("First routine")
			wg.Done()
		}()
		go func() {
			fmt.Println("Second routine")
			wg.Done()
		}()
	}
	wg.Wait()
        fmt.Println(runtime.NumGoroutine())
}
