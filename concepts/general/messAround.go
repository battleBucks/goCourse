package main

import (
	"fmt"
	"strings"
)

type properties struct {
	name  string
	model string
	job   string
}

type robot struct {
	properties
}

func main() {
	markOne := robot{
		properties: properties{
			"Saturday",
			"m4t1000",
			"happyGreeter",
		},
	}
	var instruct string
	fmt.Scanln(&instruct)

	if strings.ToLower(instruct) == "hello" || strings.ToLower(instruct) == "hi" || strings.ToLower(instruct) == "hey" {
		markOne.greeting()
	} else {
		fmt.Printf("I didn't quite get that...\n")
	}

	fmt.Printf("Would you like to hear about me?\n")

	instruct = ""
	fmt.Scanln(&instruct)

	if strings.ToLower(instruct) == "yes" || strings.ToLower(instruct) == "yeah" || strings.ToLower(instruct) == "yup" {
		markOne.introduction()
		fmt.Printf("How'd I do? Goodbye!\n")
	} else {
		fmt.Printf("OK! Goodbye!\n")
	}

}

func (r robot) greeting() {
	fmt.Printf("Hello friend! This is my first job!\n")
	instruct := ""
	fmt.Scanln(&instruct)
}
func (r robot) introduction() {
	fmt.Printf("Hello! My name is %s, my model is %s and my job is %s!\n", r.name, r.model, r.job)
	instruct := ""
	fmt.Scanln(&instruct)
}
