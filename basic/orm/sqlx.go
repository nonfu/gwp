package main

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)

type Post struct {
    Id int
    Title string
    Content string
    AuthorName string `db:author`
}

var Db *sqlx.DB

func init() {
    var err error
    Db, err = sqlx.Open("mysql", "root:root@/gwp?charset=utf8mb4&parseTime=true")
    if err != nil {
        panic(err)
    }
}

func GetPost(id int) (post Post, err error)  {
    post = Post{}
    err = Db.QueryRowx("select id, title, content, author from posts where id = ?", id).StructScan(&post)
    return
}

func (post *Post) Create() (err error) {
    res, err := Db.Exec("insert into posts (title, content, author) values (?, ?, ?)",
        post.Title, post.Content, post.AuthorName)
    if err != nil {
        return
    }
    postId, _ := res.LastInsertId()
    post.Id = int(postId)
    return
}

func main() {
    post := Post{Title: "sqlx", Content: "Sqlx Demo", AuthorName: "学院君"}
    post.Create()
    fmt.Println(post)

    GetPost(post.Id)
    fmt.Println(post)
}
