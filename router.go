package main

import (
    "net/http"
    "strings"
)

type Router struct {
    mux *http.ServeMux
}

func NewRouter() *Router {
    return &Router{mux: http.NewServeMux()}
}

func (router *Router) HandleFunc(method, path string, handler http.HandlerFunc) {
    fullPath := method + " " + path
    router.mux.HandleFunc(fullPath, func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        handler(w, r)
    })
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.Method + " " + r.URL.Path
    if _, pattern := router.mux.Handler(r); pattern != "" && strings.HasPrefix(path, pattern) {
        router.mux.ServeHTTP(w, r)
    } else {
        http.NotFound(w, r)
    }
}