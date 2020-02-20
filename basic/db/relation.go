package main

import (
"fmt"
)

func main()  {
    post := Post{Title: "Post 1", Content: "Hello World", Author: "学院君"}
    post.Create()

    comment := Comment{Content: "Test", Author: "学院君小号", Post: &post}
    comment.Create()

    readPost, _ := GetPost(post.Id)
    fmt.Println(readPost)
    fmt.Println(readPost.Comments)
    fmt.Println(readPost.Comments[0].Post)
}
