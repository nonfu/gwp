package main

import (
    "fmt"
    "net/http"
)

type IndexHandler struct {}

func (h *IndexHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Default.")
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "World!")
}

func main() {
    index := &IndexHandler{}
    hello := &HelloHandler{}
    world := &WorldHandler{}

    http.Handle("/", index)
    http.Handle("/hello", hello)
    http.Handle("/hello/", hello)
    http.Handle("/world", world)

    http.ListenAndServe(":8080", nil)
}
