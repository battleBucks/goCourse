package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	data := map[string]string{
		"Joe": "Pizza",
		"Jess": "Riggies",
		"Gray": "Tacos",
		"Brady": "Baby Jesus",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}
}
