package main

import (
        "io"
        "net/http"
)

// declare functions to send to HandleFunc
// you don't need an underlying type to attach to
// ServeHTTP, create the function to send to HandleFunc
func dogFunc(res http.ResponseWriter, req *http.Request) {
        io.WriteString(res, "dog dog doggy")
}

// HandleFunc takes a function with a response writer and pointer  to a request
// as a parameter and that's what we have here
func catFunc(res http.ResponseWriter, req *http.Request) {
        io.WriteString(res, "cat cat catty")
}

func main() {

        // attach to default serve mux
        // and takes   (route, function)
        http.HandleFunc("/dog/", dogFunc)
        http.HandleFunc("/cat", catFunc)
        http.ListenAndServe(":8080", nil)
}

