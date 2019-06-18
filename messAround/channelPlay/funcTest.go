package main

import "fmt"

func setChannel(ch chan <- int, num int) {
        ch <- num
}

func getChannel(ch <-chan int)(int) {
        return <- ch
}
func main() {

        //only a send channel designated by <-
	c := make(chan int, 2)
        //receive only:
	//c := make(<- chan int, 2)

        fmt.Printf("%T\n", c)
        setChannel((chan<-int)(c),42)
        setChannel((chan<-int)(c),43)
        fmt.Printf("%T\n", c)

	//fmt.Println(<-c)
	//fmt.Println(<-c)
        val := getChannel((<- chan int)(c))
        fmt.Println("--------------")
        fmt.Printf("%d\n", val)
        val = getChannel((<- chan int)(c))
        fmt.Printf("%d\n", val)
}
