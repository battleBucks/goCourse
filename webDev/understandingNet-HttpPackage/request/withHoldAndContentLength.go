package main

import (
        "html/template"
        "log"
        "net/http"
        "net/url"
)

type hotdog int

// type handler is anything that implements this function:
// ServeHTTP which takes a response writer and a point to a request
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        err := r.ParseForm()
        if err != nil {
                log.Fatalln(err)
        }

        data := struct {
                Method          string
                URL             *url.URL
                Submissions     map[string][]string
                Header          http.Header
                Host            string
                ContentLength   int64
        }{
                r.Method,
                r.URL,
                r.Form,
                r.Header,
                r.Host,
                r.ContentLength,
        }
        //field attached to req called form
        // saying give me the form
        tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
        //this is what is handling
        var d hotdog
        http.ListenAndServe(":8080", d)
}
