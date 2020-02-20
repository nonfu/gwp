package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (post *Post) Retrieve(id int) (err error) {
	err = post.Db.QueryRow("select id, title, content, author from posts where id = ?", id).
		Scan(&post.Id, &post.Title, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	sql := "insert into posts (title, content, author) values (?, ?, ?)"
	stmt, err := post.Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(post.Title, post.Content, post.Author)
	if err != nil {
		return
	}

	postId, _ := res.LastInsertId()
	post.Id = int(postId)
	return
}

func (post *Post) Update() (err error) {
	_, err = post.Db.Exec("update posts set title = ?, content = ?, author = ? where id = ?",
		post.Title, post.Content, post.Author, post.Id)
	return
}

func (post *Post) Delete() (err error) {
	_, err = post.Db.Exec("delete from posts where id = ?", post.Id)
	return
}
