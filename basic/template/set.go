package main

import (
    "html/template"
    "net/http"
)

func setActionExample(w http.ResponseWriter, r *http.Request)  {
    t := template.Must(template.ParseFiles("set.html"))
    t.Execute(w, "golang")
}

func main()  {
    http.HandleFunc("/set_action", setActionExample)
    http.ListenAndServe(":8080", nil)
}
