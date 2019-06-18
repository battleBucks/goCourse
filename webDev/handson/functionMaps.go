package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tpl = template.Must(template.New("fm.gohtml").Funcs(fm).ParseFiles("templates/fm.gohtml"))
}

func main() {
	data := "lower case letters"

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}
}
