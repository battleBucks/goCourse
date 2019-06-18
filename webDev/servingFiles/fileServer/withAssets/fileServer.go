package main

import (
        "io"
        "net/http"
)

func main() {
        http.HandleFunc("/", dog)
        //handle takes a route and a handler
        // so anything at path 'resources' is going to use the handler provided in second arg
        http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
        http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        io.WriteString(w, `<img src="/resources/toby.jpg">`)
}
