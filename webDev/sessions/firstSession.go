package main

import (
        "net/http"
        "github.com/satori/go.uuid"
        "html/template"
)

type user struct {
        UserName string
        First string
        Last string
}

var tpl *template.Template
var dbUsers = map[string]user{} // user ID: user
var dbSessions = map[string]string{} // session ID: user ID

func init() {
        tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
        http.HandleFunc("/",foo)
        http.HandleFunc("/bar",bar)
        http.Handle("/favicon.ico", http.NotFoundHandler())
        http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
        // get cookie
        c, err := r.Cookie("session")
        if err != nil {
                sID, _  := uuid.NewV4()
                c = &http.Cookie{
                        Name: "session",
                        Value: sID.String(),
                }
                http.SetCookie(w,c)
        }

        // if user exists, get user
        var u user
        if un, ok := dbSessions[c.Value]; ok {
                u = dbUsers[un]
        }

        // process form
        if r.Method == http.MethodPost {
                un := r.FormValue("username")
                f := r.FormValue("firstname")
                l := r.FormValue("lastname")
                u = user{un, f, l}
                dbSessions[c.Value] = un
                dbUsers[un] = u
        }

        tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
        // get cookie
        c, err := r.Cookie("session")
        if err != nil {
                http.Redirect(w, r, "/", http.StatusSeeOther)
                return
        }

        un, ok := dbSessions[c.Value]
        if !ok {
                http.Redirect(w, r, "/", http.StatusSeeOther)
                return
        }
        u := dbUsers[un]
        tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
