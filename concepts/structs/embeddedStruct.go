package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}
	v2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "silver",
		},
		luxury: true,
	}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.color)
	fmt.Println(v2.color)
}
