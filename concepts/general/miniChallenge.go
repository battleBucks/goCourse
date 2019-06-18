package main

import ("fmt")

const (
    start = 33
    end = 122
)

func main() {
    for i:= start; i <= end; i++ {
        fmt.Printf("Number: %d is character: %#U\n", i, i)
        //fmt.Printf("Number: %d is character: %s\n", i, string(i))
    }
}

