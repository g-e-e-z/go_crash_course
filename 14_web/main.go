package main

import (
    "fmt"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello Weborld</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello About</h1>")
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/about", about)
    fmt.Println("Server starting at 3000")
    http.ListenAndServe(":3000", nil)
}
