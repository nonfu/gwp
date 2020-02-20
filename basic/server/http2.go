package main

import (
    "fmt"
    "golang.org/x/net/http2"
    "net/http"
)

type CustomHandler struct {}

func (h *CustomHandler) ServeHTTP (w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "Hello World!")
}

func main()  {
    handler := CustomHandler{}
    server := http.Server{
        Handler: &handler,
        Addr: "127.0.0.1:8080",
    }
    http2.ConfigureServer(&server, &http2.Server{})
    server.ListenAndServeTLS("cert.pem", "key.pem")
}
