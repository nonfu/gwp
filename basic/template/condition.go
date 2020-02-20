package main

import (
    "html/template"
    "math/rand"
    "net/http"
    "time"
)

func process(w http.ResponseWriter, r *http.Request)  {
    t := template.Must(template.ParseFiles("condition.html"))
    rand.Seed(time.Now().Unix())
    t.Execute(w, rand.Intn(10) > 5)
}

func main()  {
    http.HandleFunc("/condition", process)
    http.ListenAndServe(":8080", nil)
}
