package main

import (
        "net/http"
        "io"
)

func main() {
        http.HandleFunc("/",foo)
        http.Handle("/flavicon.ico",http.NotFoundHandler())
        http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
        v := req.FormValue("q")
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        // change form method to get, back to post see what happens
        io.WriteString(w, `
        <form method="post">
          <input type="text" name="q">
          <input type="submit" value="Feed Me">
        </form>
        <br>`+v)
}
