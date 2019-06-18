package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		ans := i % 4
		fmt.Printf("%d\n", ans)
	}
}
