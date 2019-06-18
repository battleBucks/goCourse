package main

import (
        "fmt"
        "runtime"
        "sync"
        "sync/atomic"
)

func main() {
        fmt.Println("Starting up!")

        fmt.Println("CPUs",runtime.NumCPU())
        fmt.Println("go routines",runtime.NumGoroutine())
        var counter int64

        const gs = 100
        var wg sync.WaitGroup
        wg.Add(gs)

        for i := 0; i < gs; i++ {
                go func(){
                        atomic.AddInt64(&counter, 1)
                        runtime.Gosched()
                        fmt.Println("Counter\t", atomic.LoadInt64(&counter))
                        wg.Done()
                }()
                fmt.Println("inside! go routines",runtime.NumGoroutine())
        }
        wg.Wait()
        fmt.Println("after! go routines",runtime.NumGoroutine())
        fmt.Println("Count: ",counter)
}
