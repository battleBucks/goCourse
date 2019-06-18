package main

import (
        "fmt"
        "net/http"
)

type hotdog int

// type handler is anything that implements this function:
// ServeHTTP which takes a response writer and a point to a request
func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {

        res.Header().Set("Barrows-Key","This is from Joe")
        res.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprintln(res, "<h1>Any code you want in this func</h1>")
}

/*
var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseFiles("parseFromIndex.gohtml"))
}
*/
func main() {
        //this is what is handling
        var d hotdog
        http.ListenAndServe(":8080", d)
}
