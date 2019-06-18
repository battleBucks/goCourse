package main

import (
        "fmt"
        "runtime"
        "sync"
)

func main() {
        fmt.Println("Starting up!")

        fmt.Println("CPUs",runtime.NumCPU())
        fmt.Println("go routines",runtime.NumGoroutine())
        counter := 0

        const gs = 100
        var wg sync.WaitGroup
        wg.Add(gs)

        var mTex sync.Mutex

        for i := 0; i < gs; i++ {
                go func(){
                        mTex.Lock()
                        v := counter
                        runtime.Gosched()
                        v++
                        counter = v
                        mTex.Unlock()
                        wg.Done()
                }()
                fmt.Println("inside! go routines",runtime.NumGoroutine())
        }
        wg.Wait()
        fmt.Println("after! go routines",runtime.NumGoroutine())
        fmt.Println("Count: ",counter)
}
