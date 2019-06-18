package main

import (
        "fmt"
        "net/http"
        "log"
)

func main() {
        http.HandleFunc("/",setACookie)
        http.HandleFunc("/read",read)
        http.HandleFunc("/cookies",moreCookies)
        http.Handle("/flavicon.ico",http.NotFoundHandler())
        http.ListenAndServe(":8080", nil)
}

func setACookie(res http.ResponseWriter, req *http.Request) {
        http.SetCookie(res, &http.Cookie{
                Name: "firstCookie",
                Value: "firstCookieValue",
        })
}
func moreCookies(res http.ResponseWriter, req *http.Request) {
        http.SetCookie(res, &http.Cookie{
                Name: "secondCookie",
                Value: "secondCookieValue",
        })
}
func read(res http.ResponseWriter, req *http.Request) {
        c, err := req.Cookie("firstCookie")
        if err != nil {
                log.Println(err)
        } else {
                fmt.Fprintln(res, "Cookie one: ", c)
        }
        c, err = req.Cookie("secondCookie")
        if err != nil {
                log.Println(err)
        } else {
                fmt.Fprintln(res, "Cookie two: ", c)
        }
}
