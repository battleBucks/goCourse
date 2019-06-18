package main

import (
        "io"
        "log"
        "text/template"
        "net/http"
)

func dog(res http.ResponseWriter, req *http.Request){
        err := req.ParseForm()
        if err != nil {
                log.Fatalln(err)
        }
        io.WriteString(res, "dog dog doggy")
}
func home(res http.ResponseWriter, req *http.Request){
        err := req.ParseForm()
        if err != nil {
                log.Fatalln(err)
        }
        io.WriteString(res, "home, holmes")
}
func me(res http.ResponseWriter, req *http.Request){
        err := req.ParseForm()
        if err != nil {
                log.Fatalln(err)
        }
        err = tpl.ExecuteTemplate(res,"index.gohtml", "Joe")
}

var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){

        http.Handle("/dog/",http.HandlerFunc(dog))
        http.Handle("/me/",http.HandlerFunc(me))
        http.Handle("/",http.HandlerFunc(home))
        http.ListenAndServe(":8080",nil)
}
