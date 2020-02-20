package main

import (
    "fmt"
    "net/http"
)

func headers (w http.ResponseWriter, r *http.Request) {
    h := r.Header
    fmt.Fprintln(w, h)
}

func main()  {
    http.HandleFunc("/headers", headers)
    http.ListenAndServe(":8080", nil)
}
