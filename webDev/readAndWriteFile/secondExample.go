package main

import (
        "net/http"
        "io"
        "fmt"
        "io/ioutil"
        "os"
        "path/filepath"
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

                //store on server
                dst, err := os.Create(filepath.Join("./user/", h.Filename))
                if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                }
                defer dst.Close()

                _, err = dst.Write(bs)
                if err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                }

        }

        //<form method="POST" enctype="text/plain">
        //<form method="POST" enctype="application/x-www-form-urlencoded">
        //<form method="POST" enctype="multipart/form-data">
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        io.WriteString(w, `
        <form method="POST" enctype="multipart/form-data">
        <input type="file" name="q">
        <input type="submit" >
        </form>
        <br>` +s)
}
