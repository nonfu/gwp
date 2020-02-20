package main

import (
    "html/template"
    "net/http"
)

func includeActionExample(w http.ResponseWriter, r *http.Request)  {
    t := template.Must(template.ParseFiles("t1.html", "t2.html"))
    t.Execute(w, "Hello World!")
}

func main()  {
    http.HandleFunc("/include", includeActionExample)
    http.ListenAndServe(":8080", nil)
}
