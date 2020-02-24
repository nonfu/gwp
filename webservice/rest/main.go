package main

import (
	"net/http"
)

func main() {
	/*db, err := sql.Open("mysql", "root:root@/gwp?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}*/

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post", HandleRequest(&FakePost{}))
	server.ListenAndServe()
}

type FakePost struct {
	Id      int
	Title   string
	Content string
	Author  string
}

func (post *FakePost) Retrieve(id int) (err error) {
	post.Id = id
	return
}

func (post *FakePost) Create() (err error) {
	return
}

func (post *FakePost) Update() (err error) {
	return
}

func (post *FakePost) Delete() (err error) {
	return
}