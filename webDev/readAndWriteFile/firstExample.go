package main

import (
        "net/http"
        "io"
        "fmt"
        "io/ioutil"
)

func main() {
        http.HandleFunc("/", foo)
        http.Handle("/favicon.ico", http.NotFoundHandler())
        http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
        var s string
        fmt.Println(req.Method)
        // as it looks, only goes in block if it's a post
        if req.Method == http.MethodPost {
                // open
                f, h, err := req.FormFile("q")
                if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                }
                defer f.Close()

                // check out the info
                fmt.Println("\nfile:", f, "\nheader:", h, "\nerr:", err)

                //read
                bs, err := ioutil.ReadAll(f)
                if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                }
                s = string(bs)
        }

        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        io.WriteString(w, `
        <form method="POST" enctype="multipart/form-data">
        <input type="file" name="q">
        <input type="submit">
        </form>
        <br>` +s)
}
