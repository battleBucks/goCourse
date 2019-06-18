package main

import (
        "html/template"
        "log"
        "net/http"
)

type hotdog int

// type handler is anything that implements this function:
// ServeHTTP which takes a response writer and a point to a request
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
        err := req.ParseForm()
        if err != nil {
                log.Fatalln(err)
        }

        //field attached to req called form
        // saying give me the form
        tpl.ExecuteTemplate(w, "parseFromIndex.gohtml", req.Form)
}

var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseFiles("parseFromIndex.gohtml"))
}

func main() {
        //this is what is handling
        var d hotdog
        http.ListenAndServe(":8080", d)
}
