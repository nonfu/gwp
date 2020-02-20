package main

import (
    "encoding/base64"
    "fmt"
    "net/http"
    "time"
)

func setMessage(w http.ResponseWriter, r *http.Request)  {
    msg := []byte("Hello World!")
    cookie := http.Cookie{
        Name: "flash_message",
        Value: base64.URLEncoding.EncodeToString(msg),
    }
    http.SetCookie(w, &cookie)
}

func getMessage(w http.ResponseWriter, r *http.Request)  {
    cookie, err := r.Cookie("flash_message")
    if err != nil {
        fmt.Fprintln(w, "No flash message found in Cookie")
    } else {
        delCookie := http.Cookie{
            Name: "flash_message",
            MaxAge: -1,
            Expires: time.Unix(1, 0),
        }
        http.SetCookie(w, &delCookie)
        message, _ := base64.URLEncoding.DecodeString(cookie.Value)
        fmt.Fprintln(w, string(message))
    }
}

func main() {
    http.HandleFunc("/set_message", setMessage)
    http.HandleFunc("/get_message", getMessage)
    http.ListenAndServe(":8080", nil)
}
