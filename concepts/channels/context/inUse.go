package main

import (
        "fmt"
        "context"
        "runtime"
        "time"
)

func main() {
        ctx, cancel := context.WithCancel(context.Background())
        //idiomatic to defer cancel() here

        fmt.Println("error check 1:", ctx.Err())
        fmt.Println("num go routines:", runtime.NumGoroutine())

        go func() {
                n := 0
                for {
                        select {
                        case <-ctx.Done():
                                return
                        default:
                                n++
                                time.Sleep(time.Millisecond * 200)
                                fmt.Println("working",n)
                        }
                }
        }()
        time.Sleep(time.Second * 2)
        fmt.Println("error check 2:", ctx.Err())
        fmt.Println("num go routines 2:", runtime.NumGoroutine())

        fmt.Println("about to cancel context")
        cancel()
        fmt.Println("cancelled context")

        time.Sleep(time.Second * 2)
        fmt.Println("error check 3:", ctx.Err())
        fmt.Println("num go routines 3:", runtime.NumGoroutine())
}
