package main

import (
    "html/template"
    "net/http"
)

func iterationActionExample(w http.ResponseWriter, r *http.Request)  {
    t := template.Must(template.ParseFiles("iteration.html"))
    daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
    t.Execute(w, daysOfWeek)
}

func main() {
    http.HandleFunc("/iteration", iterationActionExample)
    http.ListenAndServe(":8080", nil)
}
