package main

import (
    "html/template"
    "net/http"
)

func xssAttackExample(w http.ResponseWriter, r *http.Request)  {
    t := template.Must(template.ParseFiles("xss.html"))
    if r.Method == "GET" {
        t.Execute(w, nil)
    } else {
        t.Execute(w, template.HTML(r.FormValue("comment")))
    }
}

func main()  {
    http.HandleFunc("/xss", xssAttackExample)
    http.ListenAndServe(":8080", nil)
}
