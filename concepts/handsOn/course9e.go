package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("Begin goRoutines:", runtime.NumGoroutine())
        numRoutines := 100
	var counter int64

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
                        fmt.Prinln(atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Mid goRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("End goRoutines:", runtime.NumGoroutine())

	fmt.Println(counter)
}
