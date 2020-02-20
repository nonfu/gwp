package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
)

func process(w http.ResponseWriter, r *http.Request)  {
    r.ParseForm()
    fmt.Fprintln(w, r.PostForm)
}

func upload(w http.ResponseWriter, r *http.Request)  {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("upload.gtpl")
        t.Execute(w, nil)
    } else {
        file, _, err := r.FormFile("uploaded")
        if err == nil {
            data, err := ioutil.ReadAll(file)
            if err == nil {
                fmt.Fprintln(w, string(data))
            }
        }
    }
}

func main()  {
    http.HandleFunc("/form", process)
    http.HandleFunc("/upload", upload)
    http.ListenAndServe(":8080", nil)
}