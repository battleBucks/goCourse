package main

import ("fmt")

var a int
var b string
var c bool
type hotDiggety int

func main() {

    fmt.Printf("a = %d and is of type: %T\n",a, a)
    fmt.Printf("b = %s and is of type: %T\n",b, b)
    fmt.Printf("c = %t and is of type: %T\n",c, c)

    a = 42
    b = "James Bond"
    c = true

    fmt.Printf("a = %d and is of type: %T\n",a, a)
    fmt.Printf("b = %s and is of type: %T\n",b, b)
    fmt.Printf("c = %t and is of type: %T\n",c, c)

    s := fmt.Sprintf("a, b, c is %d, %s, %t!\n",a, b, c)
    fmt.Println(s)

    var x hotDiggety
    fmt.Printf("x = %v\tx is type %T\n",x,x)
    x = 42
    fmt.Printf("x = %v\tx is type %T\n",x,x)
    var y int
    y = int(x)
    fmt.Println("last value is: ", y)
}

