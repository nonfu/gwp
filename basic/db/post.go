package main

type Post struct {
    Id int
    Title string
    Content string
    Author string
    Comments []Comment
}

func Posts(limit int) (posts []Post, err error) {
    rows, err := Db.Query("select id, title, content, author from posts limit ?", limit)
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        post := Post{}
        err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author)
        if err != nil {
            return
        }
        posts = append(posts, post)
    }
    rows.Close()
    return
}

func GetPost(id int) (post Post, err error) {
    post = Post{}
    err = Db.QueryRow("select id, title, content, author from posts where id = ?", id).
        Scan(&post.Id, &post.Title, &post.Content, &post.Author)

    // 查询与之关联的 comments 记录
    rows, err := Db.Query("select id, content, author from comments where post_id = ?", post.Id)
    for rows.Next() {
        comment := Comment{Post: &post}
        err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
        if err != nil {
            return
        }
        post.Comments = append(post.Comments, comment)
    }
    rows.Close()
    return
}

func (post *Post) Create() (err error) {
    sql := "insert into posts (title, content, author) values (?, ?, ?)"
    stmt, err := Db.Prepare(sql)
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

func (post *Post) Update() (err error)  {
    _, err = Db.Exec("update posts set title = ?, content = ?, author = ? where id = ?",
        post.Title, post.Content, post.Author, post.Id)
    return
}

func (post *Post) Delete() (err error) {
    _, err = Db.Exec("delete from posts where id = ?", post.Id)
    return
}
