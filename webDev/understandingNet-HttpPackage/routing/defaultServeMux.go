package main

import (
        "io"
        "net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
        io.WriteString(res, "dog dog doggy")
}

type hotcat int

func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
        io.WriteString(res, "cat cat catty")
}

func main() {
        var d hotdog
        var c hotcat

        /*
        NewServeMux can be passed into ListenAndServe like this:
        mux := http.NewServeMux()
        mux.Handle("/dog/", d)
        mux.Handle("/cat/", c)
        http.ListenAndServe(":8080", mux)
        */
        // or the defaultServeMux is passed in with nil
        // using http Handle
        http.Handle("/dog/", d)
        http.Handle("/cat", c)
        http.ListenAndServe(":8080", nil)
}

