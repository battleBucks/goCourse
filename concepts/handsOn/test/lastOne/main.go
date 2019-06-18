package main

import (
        "fmt"
        "goCourse/concepts/handsOn/test/lastOne/dog"
)

type canine struct {
        name string
        age int
}

func main() {
        patty := canine{
                name: "Patty",
                age: dog.Years(10),
        }
        fmt.Println(patty)
        fmt.Println(dog.YearsTwo(20))
}
