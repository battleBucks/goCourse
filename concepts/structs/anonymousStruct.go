package main

import (
	"fmt"
)

func main() {
	a1 := struct {
		instruments []string
		members     map[string]string
	}{
		instruments: []string{"piano", "guitar", "saxophone"},
		members:     map[string]string{"Brady": "saxophone"},
	}
	fmt.Println(a1)
}
