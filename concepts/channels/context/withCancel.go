package main

import (
        "fmt"
        "context"
)

func main() {
        ctx := context.Background()

        fmt.Println("context:\t", ctx)
        fmt.Println("context err:\t", ctx.Err())
        fmt.Printf("context type:\t%T\n", ctx)
        fmt.Println("-----------")

        ctx, _ = context.WithCancel(ctx)
        fmt.Println("context:\t", ctx)
        fmt.Println("context err:\t", ctx.Err())
        fmt.Printf("context type:\t%T\n", ctx)
        fmt.Println("-----------")
}
