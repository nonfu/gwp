package main

import (
    "fmt"
    "github.com/nonfu/gwp/chitchat/route"
    "net/http"
    "time"
)

func main()  {
    fmt.Println("ChitChat", route.Version(), "started at", route.Config.Address)

    // 创建路由器
    mux := http.NewServeMux()

    // 处理静态资源
    files := http.FileServer(http.Dir(route.Config.Static))
    mux.Handle("/static/", http.StripPrefix("/static", files))

    // 首页及错误路由
    mux.HandleFunc("/", route.Index)
    mux.HandleFunc("/err", route.Err)

    // 认证相关路由
    mux.HandleFunc("/login", route.Login)
    mux.HandleFunc("/logout", route.Logout)
    mux.HandleFunc("/signup", route.Signup)
    mux.HandleFunc("/signup_account", route.SignupAccount)
    mux.HandleFunc("/authenticate", route.Authenticate)

    // 主题相关路由
    mux.HandleFunc("/thread/new", route.NewThread)
    mux.HandleFunc("/thread/create", route.CreateThread)
    mux.HandleFunc("/thread/post", route.PostThread)
    mux.HandleFunc("/thread/read", route.ReadThread)

    // 启动 HTTP 服务器
    server := &http.Server{
        Addr:           route.Config.Address,
        Handler:        mux,
        ReadTimeout:    time.Duration(route.Config.ReadTimeout * int64(time.Second)),
        WriteTimeout:   time.Duration(route.Config.WriteTimeout * int64(time.Second)),
        MaxHeaderBytes: 1 << 20,
    }
    server.ListenAndServe()
}
