package main

import (
        "fmt"
        "net/http"
)

type hotdog int

// type handler is anything that implements this function:
// ServeHTTP which takes a response writer and a point to a request
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "anything to the write in here")
}

func main() {
        //this is what is handling
        var d hotdog
        http.ListenAndServe(":8080", d)
}
