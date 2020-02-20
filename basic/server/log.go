package main

import (
    "fmt"
    "net/http"
    "reflect"
    "runtime"
)

func helloWorld(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello World!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
        fmt.Println("Handler function called - " + name)
        h(w, r)
    }
}

func main()  {
    http.HandleFunc("/hello", log(helloWorld))
    http.ListenAndServe(":8080", nil)
}
