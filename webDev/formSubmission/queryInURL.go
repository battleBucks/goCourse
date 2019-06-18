package main

import (
        "net/http"
        "io"
)

func main() {
        http.HandleFunc("/", foo)
        http.Handle("/favicon.ico", http.NotFoundHandler())
        http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
        v := req.FormValue("q")
        io.WriteString(w, "my search: "+v)
}
