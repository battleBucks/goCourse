package main

import ("fmt")

const (
    a = 42
    b int = 42
)
const (
    e = 2017 + iota
    f = 2017 + iota
    g = 2017 + iota
    h = 2017 + iota
)

func main() {
    fmt.Printf("Decimal: %d\n",a)
    fmt.Printf("Binary: %b\n",a)
    fmt.Printf("Hex: %x\n",a)
    fmt.Printf("typed constant b is %d\n", b)
    c := a << 1
    fmt.Printf("Decimal: %d\n",c)
    fmt.Printf("Binary: %b\n",c)
    fmt.Printf("Hex: %x\n",c)

    d := `This
is a
                raw string`
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
    fmt.Println(g)
    fmt.Println(h)
}

