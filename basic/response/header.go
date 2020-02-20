package main

import (
    "fmt"
    "net/http"
)

func statusCodeExample(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(501)
    fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Location", "https://xueyuanjun.com")
    w.WriteHeader(302)
}

func main()  {
    http.HandleFunc("/statusCode", statusCodeExample)
    http.HandleFunc("/redirect", headerExample)
    http.ListenAndServe(":8080", nil)
}
