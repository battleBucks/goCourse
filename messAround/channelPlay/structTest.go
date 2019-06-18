package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func setChannel(ch chan<- person, p person) {
	ch <- p
}

func getChannel(ch <-chan person) person {
	return <-ch
}
func main() {

	p1 := person{
		first: "Joe",
		last:  "Barrows",
		age:   39,
	}
	p2 := person{
		first: "Brady",
		last:  "Gross",
		age:   11,
	}
	//only a send channel designated by <-
	c := make(chan person, 2)
	//receive only:
	//c := make(<- chan int, 2)

	setChannel((chan<- person)(c), p1)
	setChannel((chan<- person)(c), p2)
	fmt.Printf("%T\n", c)

	val := getChannel((<-chan person)(c))
	fmt.Println("--------------")
	fmt.Printf("%v\n", val)
        fmt.Printf("First Name: %s\tLast Name: %s\tAge: %d\n", val.first,val.last,val.age)
	val = getChannel((<-chan person)(c))
        fmt.Printf("First Name: %s\tLast Name: %s\tAge: %d\n", val.first,val.last,val.age)
}
