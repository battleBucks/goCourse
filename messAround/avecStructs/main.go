package main

import (
        "fmt"
        "os"
        "text/template"
)

type Person struct {
        First string
        Last string
        Age int
}

var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseFiles("template.gohtml"))
}
func main() {
        p1 := Person{
                First: "Joe",
                Last: "Barrows",
                Age: 39,
        }
        err := tpl.Execute(os.Stdout, p1)
        if err != nil {
                fmt.Println(err)
        }
}
