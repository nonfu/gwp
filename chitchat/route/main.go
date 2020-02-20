package route

import (
    "github.com/nonfu/gwp/chitchat/data"
    "net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
    threads, err := data.Threads()
    if err != nil {
        error_message(writer, request, "Cannot get threads")
    } else {
        _, err := session(writer, request)
        if err != nil {
            generateHTML(writer, threads, "layout", "public.navbar", "index")
        } else {
            generateHTML(writer, threads, "layout", "private.navbar", "index")
        }
    }
}

func Err(writer http.ResponseWriter, request *http.Request)  {
    vals := request.URL.Query()
    _, err := session(writer, request)
    if err != nil {
        generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
    } else {
        generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
    }
}