package main

import "net/http"

func writeExample(w http.ResponseWriter, r *http.Request)  {
    str := `<html> 
    <head>
        <title>Go Web Programming</title>
    </head> 
    <body>
        <h1>Hello World</h1>
    </body> 
</html>`
    w.Write([]byte(str))
}

func main()  {
    http.HandleFunc("/write", writeExample)
    http.ListenAndServe(":8080", nil)
}
