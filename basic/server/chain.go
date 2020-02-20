package main

import (
    "fmt"
    "net/http"
)

type HelloWorldHandler struct {}

func (hw *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func logHandler(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Handler called: %T\n", h)
        h.ServeHTTP(w, r)
    })
}

func main() {
    handler := &HelloWorldHandler{}
    http.Handle("/hello", logHandler(handler))
    http.ListenAndServe(":8080", nil)
}
