package main

import (
    "fmt"
    "net/http"
)

func main() {
    router := NewRouter()
    
    // Use Chain to apply Logger middleware to the handler
    router.HandleFunc("GET", "/", Chain(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to Air!")
    }, Logger()))

    fmt.Println("Starting Air server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        fmt.Println("Server failed to start:", err)
    }
}

