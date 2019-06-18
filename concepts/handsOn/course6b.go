package main

import "fmt"
import "math"

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	a := c.radius * c.radius * math.Pi
	return a
}
func (s square) area() float64 {
	a := s.side * s.side
	return a
}
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		side: 15,
	}
	c := circle{
		radius: 12.345,
	}
	info(s)
	a := s.area()
	fmt.Println("Printing area of square: ", a)
	info(c)
	b := c.area()
	fmt.Println("Printing area of circle: ", b)

        func (x float64) {
		fmt.Println("Printing square side with my anonymous function...", x)
	}(s.side)
	
	f := func(x float64) {
		fmt.Println("Printing circle radius with my function assigned to a variable...", x)
}
	f(c.radius)
}
