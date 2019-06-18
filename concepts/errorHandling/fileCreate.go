package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {
	f, err := os.Create("names.txt")
	check(err)
	defer f.Close()

	r := strings.NewReader("Wasssssuuuup")

	io.Copy(f, r)
}
