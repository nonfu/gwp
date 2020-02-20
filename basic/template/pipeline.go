package main

import (
    "html/template"
    "net/http"
)

func pipelineExample(w http.ResponseWriter, r *http.Request)  {
    t := template.Must(template.ParseFiles("pipeline.html"))
    t.Execute(w, 12.3456)
}

func main()  {
    http.HandleFunc("/pipeline", pipelineExample)
    http.ListenAndServe(":8080", nil)
}
