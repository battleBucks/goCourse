package main

import (
        "fmt"
        "log"
        "strconv"
        "io"
        "net/http"
)

func main() {
        http.HandleFunc("/",jb)
        http.Handle("/favicon.iso", http.NotFoundHandler())
        http.ListenAndServe(":8080",nil)
}

func jb(res http.ResponseWriter, req *http.Request) {
        c, err := req.Cookie("liveOne")

        if err != nil {
                if err == http.ErrNoCookie {
                        c = &http.Cookie{
                                Name: "liveOne",
                                Value: "0",
                        }
                } else {
                        log.Println("ERROR: ", err)
                }
        }
        total, err := strconv.Atoi(c.Value)
        total++
        c.Value = strconv.Itoa(total)

        http.SetCookie(res, c)
        response := fmt.Sprintf("Total hits: %s", c.Value)
        io.WriteString(res, response)
}
