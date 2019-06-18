package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("stpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	ghandi := sage{
		Name:  "Ghandi",
		Motto: "Be the change",
	}
	list := []sage{buddha, ghandi}
/*
	data := struct {
		Peeps []sage
	}{
		list,
	}
*/
	err := tpl.Execute(os.Stdout, list)
	if err != nil {
		log.Fatalln(err)
	}

} //end main
