package main

import (
        "fmt"
)

type person struct {
        First string
        Last string
        Age int
}

func (p *person) Speak() {
        fmt.Println("Hello! I'm", p.First)
}

type human interface{
        Speak()
}

func saySomething(h human) {
        h.Speak()
}

func main() {
        fmt.Println("Beginning")
        p1 := person{
                First: "Joe",
                Last: "Barrows",
                Age: 39,
        }
        fmt.Println(p1)
        p1.Speak()
        saySomething(&p1)
}
