package main

import (
        //"io"
        "log"
        "net/http"
        "html/template"
)

type Person struct {
        Name string
        Age int
}

var tpl *template.Template

func init() {
        tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func kids(w http.ResponseWriter, req *http.Request) {
        p := []Person{
                {
                        Name: "Kendall",
                        Age: 13,
                },
                {
                        Name: "Brady",
                        Age: 11,
                },
                {
                        Name: "Izzy",
                        Age: 9,
                },
                {
                        Name: "Margot",
                        Age: 7,
                },
                {
                        Name: "Charlotte",
                        Age: 7,
                },
                {
                        Name: "Tessa",
                        Age: 5,
                },
                {
                        Name: "Olivia",
                        Age: 4,
                },
        }
        err := req.ParseForm()
        if err != nil {
                log.Fatalln(err)
        }
        //io.WriteString(w, "List em up")
        tpl.Execute(w, p)
        //tpl.ExecuteTemplate(w, "index.html", req.Form)
}

func main() {
        http.HandleFunc("/", kids)
        http.Handle("flavicon.ico", http.NotFoundHandler())
        http.ListenAndServe(":8080", nil)
}
