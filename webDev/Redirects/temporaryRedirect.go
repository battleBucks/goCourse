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
        http.HandleFunc("/barred", barred)
        http.Handle("flavicon.ico", http.NotFoundHandler())
        http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
        fmt.Print("Your request method at foo: ", req.Method, "\n\n")
        fmt.Fprintf(res, "hello")
}
func bar(res http.ResponseWriter, req *http.Request) {
        fmt.Print("Your request method at bar: ", req.Method, "\n\n")
        //res.Header().Set("Location", "/")
        //res.WriteHeader(http.StatusSeeOther)

        // or use redirect from http package
        // with a writer, reader, location, and code
        http.Redirect(res, req, "/", http.StatusSeeOther)
}
func barred(res http.ResponseWriter, req *http.Request) {
        fmt.Print("Your request method at barRED: ", req.Method, "\n\n")
        tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
