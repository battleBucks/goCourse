package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := person{
		First: "Joe",
		Last:  "Barrows",
		Age:   39,
	}
	p2 := person{
		First: "Brady",
		Last:  "Gross",
		Age:   12,
	}
	pl := []person{
		p1,
		p2,
	}
	fmt.Println(pl)
	b, err := json.Marshal(pl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// json string is what Ive just printed
	js := `[{"First":"Joe","Last":"Barrows","Age":39},{"First":"Brady","Last":"Gross","Age":12}]`

	// need to convert to bytes

	bs := []byte(js)
	//also need a struct which will be above
	people := []person{}
	err = json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
