package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {
	f, err := os.Open("names.txt")
	check(err)
	defer f.Close()

	//returns a byte slice and an error
	bs, err := ioutil.ReadAll(f)
	check(err)
	fmt.Println(string(bs))
}
