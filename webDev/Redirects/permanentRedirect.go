package main

import (
        "fmt"
        "html/template"
        "net/http"
)

var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
        http.HandleFunc("/", foo)
        http.HandleFunc("/bar", bar)
        http.Handle("flavicon.ico", http.NotFoundHandler())
        http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
        fmt.Print("Your request method at foo: ", req.Method, "\n\n")
        fmt.Fprintf(res, "hello")
}
func bar(res http.ResponseWriter, req *http.Request) {
        fmt.Print("Your request method at bar: ", req.Method, "\n\n")
        http.Redirect(res, req, "/", http.StatusMovedPermanently)
}
