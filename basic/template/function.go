package main

import (
    "html/template"
    "net/http"
    "time"
)

func formatDate(t time.Time) string {
    layout := "2006-01-02 15:04:05"
    return t.Format(layout)
}

func customFunctionExample(w http.ResponseWriter, r *http.Request)  {
    funcMap := template.FuncMap{
        "fdate": formatDate,
    }
    t := template.New("function.html").Funcs(funcMap)
    t, _ = t.ParseFiles("function.html")
    t.Execute(w, time.Now())
}

func main()  {
    http.HandleFunc("/date_format", customFunctionExample)
    http.ListenAndServe(":8080", nil)
}
