package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    http.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Here are the recipes")
    })

    fmt.Println("Server is listening on :8080...")
    http.ListenAndServe(":8080", nil)
}
