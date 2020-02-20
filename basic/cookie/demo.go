package main

import (
    "fmt"
    "net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request)  {
    c1 := http.Cookie{
        Name: "first_cookie",
        Value: "Go Web Programming",
        HttpOnly: true,
    }
    c2 := http.Cookie{
        Name: "second_cookie",
        Value: "Manning Publications Co",
        HttpOnly: true,
    }
    //w.Header().Set("Set-Cookie", c1.String())
    //w.Header().Add("Set-Cookie", c2.String())
    http.SetCookie(w, &c1)
    http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request)  {
    h := r.Header["Cookie"]
    c1, _ := r.Cookie("first_cookie")
    cs := r.Cookies()

    fmt.Fprintln(w, h)
    fmt.Fprintln(w, c1)
    fmt.Fprintln(w, cs)
}

func main()  {
    http.HandleFunc("/setcookie", setCookie)
    http.HandleFunc("/getcookie", getCookie)
    http.ListenAndServe(":8080", nil)
}
