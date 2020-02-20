package main

import (
    "encoding/json"
    "net/http"
)

type Post struct {
    User string
    Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "application/json")
    post := Post{
        User: "学院君",
        Threads: []string{"first", "second", "third"},
    }
    data, _ := json.Marshal(post)
    w.Write(data)
}

func main() {
    http.HandleFunc("/json", jsonExample)
    http.ListenAndServe(":8080", nil)
}
