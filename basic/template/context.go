package main

import (
    "html/template"
    "net/http"
)

func contextExample(w http.ResponseWriter, r *http.Request)  {
    t := template.Must(template.ParseFiles("context.html"))
    content := `I asked: <i>"What's up?"</i>`
    t.Execute(w, content)
}

func main()  {
    http.HandleFunc("/context", contextExample)
    http.ListenAndServe(":8080", nil)
}
