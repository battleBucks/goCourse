package main

import (
        "net/http"
        "html/template"
        "log"
)

var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
        http.HandleFunc("/", dog)
        http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
        http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
        err := tpl.Execute(w,nil)
        if err != nil {
                log.Fatalln("template didn't execute: ", err)
        }
}
