package main

import "fmt"

func main() {
	jb := []string{"Joe", "Barrows", "pizza"}
	jt := []string{"Jess", "Thaler", "Riggies"}
	gl := []string{"Gray", "Lane", "tacos"}

	xp := [][]string{jb, jt, gl}
	fmt.Println(xp)
}
