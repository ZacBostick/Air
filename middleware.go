package main

import (
    "log"
    "net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logger() Middleware {
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            log.Printf("Logged connection from %s to %s", r.RemoteAddr, r.URL.Path)
            next(w, r)
        }
    }
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, m := range middlewares {
        f = m(f)
    }
    return f
}