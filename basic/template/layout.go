package main

import (
    "html/template"
    "math/rand"
    "net/http"
    "time"
)

func layoutExample(w http.ResponseWriter, r *http.Request)  {
    rand.Seed(time.Now().Unix())
    var t *template.Template
    if rand.Intn(10) > 5 {
        t = template.Must(template.ParseFiles("layout.html", "hello_blue.html"))
    } else {
        t = template.Must(template.ParseFiles("layout.html"))
    }
    t.ExecuteTemplate(w, "layout", "")
}

func main()  {
    http.HandleFunc("/layout", layoutExample)
    http.ListenAndServe(":8080", nil)
}
