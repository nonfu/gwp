package main

import (
    "html/template"
    "net/http"
)

func parseFiles(w http.ResponseWriter, r *http.Request)  {
    t := template.Must(template.ParseFiles("tmpl.html"))
    t.Execute(w, "Hello World!")
}

func parseString(w http.ResponseWriter, r *http.Request)  {
    tmpl := `<!DOCTYPE html> <html>
        <head>
            <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
            <title>Go Web Programming</title>
        </head>
        <body>
            {{ . }}
        </body> 
    </html>`

    t := template.New("tmpl.html")
    t.Parse(tmpl)
    t.Execute(w, "Hello World!")
}

func main()  {
    http.HandleFunc("/template", parseFiles)
    http.HandleFunc("/string", parseFiles)
    http.ListenAndServe(":8080", nil)
}
