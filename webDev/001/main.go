package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer file.Close()
	err = tpl.Execute(file, nil)
	if err != nil {
		log.Fatalln(err)
	}

} //end main
