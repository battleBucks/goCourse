package main

import "fmt"

func main() {
	m := map[string]int{
		"Joe":  39,
		"Jess": 36,
		"Gray": 8,
	}
	fmt.Println(m)
	fmt.Println(m["Joe"])

	v, ok := m["Bebop"]
	fmt.Println(v)
	fmt.Println(ok)
	m["George"] = 2

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Joe")
	fmt.Println(m)
}
