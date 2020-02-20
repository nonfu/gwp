package main

import (
	"database/sql"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/gwp?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post", HandleRequest(&Post{Db: db}))
	server.ListenAndServe()
}