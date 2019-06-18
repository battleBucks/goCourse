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
        fmt.Println("Here goes nothin'")
        arr := []string{
                "joe",
                "barrows",
                "pizza",
                "pho",
        }
        err := tpl.Execute(os.Stdout, arr)
        if err != nil {
                fmt.Println(err)
        }

}
