package main

import (
    "fmt"
)

func main()  {
    post := Post{Title: "Post 1", Content: "Hello World!", Author: "学院君"}

    // 创建记录
    fmt.Println(post)
    post.Create()
    fmt.Println(post)

    // 获取单条记录
    readPost, _ := GetPost(post.Id)
    fmt.Println(readPost)

    // 更新记录
    readPost.Content = "Hello Golang!"
    readPost.Author = "nonfu"
    readPost.Update()

    // 获取列表
    posts, _ := Posts(5)
    fmt.Println(posts)

    // 删除记录
    readPost.Delete()
}
