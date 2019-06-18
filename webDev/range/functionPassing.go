package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("ftpl.gohtml").Funcs(fm).ParseFiles("ftpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
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
	mlk := sage{
		Name:  "Ghandi",
		Motto: "Be the change",
	}
	list := []sage{buddha, ghandi, mlk}
	err := tpl.Execute(os.Stdout, list)
	if err != nil {
		log.Fatalln(err)
	}

} //end main
