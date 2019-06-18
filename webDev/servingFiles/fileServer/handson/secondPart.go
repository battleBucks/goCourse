package main

import (
        "net/http"
)

func main() {
        log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
