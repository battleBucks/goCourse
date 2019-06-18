package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Begin goRoutines:", runtime.NumGoroutine())
	const numRoutines = 100
	counter := 0

        var mTex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func() {
                        mTex.Lock()
			v := counter
			v++
			counter = v
                        fmt.Println(counter)
                        mTex.Unlock()
			wg.Done()
		}()
		fmt.Println("Mid goRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("End goRoutines:", runtime.NumGoroutine())

        fmt.Println(counter)
}
