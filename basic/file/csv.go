package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

type Post struct {
    Id int
    Title string
    Content string
    Author string
}

func main()  {
    csvFile, err := os.Create("posts.csv")
    if err != nil {
        panic(err)
    }
    defer csvFile.Close()

    allPosts := []Post{
        Post{Id: 1, Title: "Book 1", Content: "Hello World!", Author: "Sau Sheong"},
        Post{Id: 2, Title: "Book 2", Content: "Bonjour Monde!", Author: "Pierre"},
        Post{Id: 3, Title: "Book 3", Content: "Hola Mundo!", Author: "Pedro"},
        Post{Id: 4, Title: "Book 4", Content: "Greetings Earthlings!", Author: "Sau Sheong"},
    }

    writer := csv.NewWriter(csvFile)
    for _, post := range allPosts {
        line := []string{
            strconv.Itoa(post.Id),
            post.Title,
            post.Content,
            post.Author,
        }
        err := writer.Write(line)
        if err != nil {
            panic(err)
        }
    }
    writer.Flush()  // 确保内存中的缓冲数据已经写入文件，以便接下来读取

    file, err := os.Open("posts.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.FieldsPerRecord = -1  // 校验每一行记录是否每个字段都存在，设置为-1则不做这个校验
    record, err := reader.ReadAll()
    if err != nil {
        panic(err)
    }

    var posts []Post
    for _, item := range record {
        id, _ := strconv.ParseInt(item[0], 0, 0)
        post := Post{Id: int(id), Title: item[1], Content: item[2], Author: item[3]}
        posts = append(posts, post)
    }

    fmt.Println(posts[0].Id)
    fmt.Println(posts[0].Title)
    fmt.Println(posts[0].Content)
    fmt.Println(posts[0].Author)
}
