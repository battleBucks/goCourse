package main

import (
	"fmt"
)

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	p1 := person{
		first:       "joe",
		last:        "barrows",
		favIceCream: []string{"Vanilla", "Peanut Butter Cup"},
	}
	p2 := person{
		first:       "jess",
		last:        "thaler",
		favIceCream: []string{"Black Cherry + Vanilla twist", "Orange sherbert"},
	}
	fmt.Println(p1)

	personMap := map[string]person{}
	personMap[p1.last] = p1
	personMap[p2.last] = p2

	for _, record := range personMap {
		toPrint := fmt.Sprintf("%s %s's favorite ice cream flavors are: ", record.first, record.last)
		for i, v := range record.favIceCream {
			if i < len(record.favIceCream)-1 {
				toPrint += fmt.Sprintf("%s, ", v)
			} else {
				toPrint += fmt.Sprintf("%s", v)
			}
		}
		fmt.Println(toPrint)
	}
}
