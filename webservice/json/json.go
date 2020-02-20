package main

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

type Post struct {
    Id       int       `json:"id"`
    Content  string    `json:"content"`
    Author   Author    `json:"author"`
    Comments []Comment `json:"comments"`
}

type Author struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

type Comment struct {
    Id      int    `json:"id"`
    Content string `json:"content"`
    Author  string `json:"author"`
}

func parseWithUnmarshal() {
    jsonFile, err := os.Open("post.json")
    if err != nil {
        fmt.Println("Error opening JSON file:", err)
        return
    }
    defer jsonFile.Close()

    jsonData, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        fmt.Println("Error reading JSON data:", err)
        return
    }

    var post Post
    json.Unmarshal(jsonData, &post)
    fmt.Println(post)
}

func parseWithDecoder() {
    jsonFile, err := os.Open("post.json")
    if err != nil {
        fmt.Println("Error opening JSON file:", err)
        return
    }
    defer jsonFile.Close()

    decoder := json.NewDecoder(jsonFile)
    for {
        var post Post
        err := decoder.Decode(&post)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Error decoding JSON:", err)
            return
        }
        fmt.Println(post)
    }
}

func createWithMarshal() {
    post := Post{
        Id: 1,
        Content: "Hello World!",
        Author: Author{Id: 2, Name: "Sau Sheong",},
        Comments: []Comment{
            Comment{Id: 3, Content: "Have a great day!", Author: "Adam",},
            Comment{Id: 4, Content: "How are you today?", Author: "Betty",},
        },
    }

    jsonData, err := json.MarshalIndent(&post, "", "\t")
    if err != nil {
        fmt.Println("Error marshalling to JSON:", err)
        return
    }
    err = ioutil.WriteFile("post2.json", jsonData, 0644)
    if err != nil {
        fmt.Println("Error writing JSON to file:", err)
        return
    }
}

func createWithEncoder() {
    post := Post{
        Id: 1,
        Content: "Hello World!",
        Author: Author{Id: 2, Name: "Sau Sheong",},
        Comments: []Comment{
            Comment{Id: 3, Content: "Have a great day!", Author: "Adam",},
            Comment{Id: 4, Content: "How are you today?", Author: "Betty",},
        },
    }

    jsonFile, err := os.Create("post3.json")
    if err != nil {
        fmt.Println("Error creating json file:", err)
        return
    }

    encoder := json.NewEncoder(jsonFile)
    encoder.SetIndent("", "\t") // 缩进，这一步不是必须的
    err = encoder.Encode(&post)
    if err != nil {
        fmt.Println("Error encoding json to file:", err)
        return
    }
}

func main() {
    parseWithUnmarshal()
    parseWithDecoder()
    createWithMarshal()
    createWithEncoder()
}
