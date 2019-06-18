package main

import (
        "fmt"
)

type Person struct {
        Name string
        Age int
}

func main() {
        fmt.Println("hello")
        p := []Person{
                {
                        Name: "Brady",
                        Age: 11,
                },
                {
                        Name: "Kendall",
                        Age: 13,
                },
        }
        for _, val := range p {
                fmt.Println(val.Name)
        }
}
