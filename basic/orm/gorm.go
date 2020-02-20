package main

import (
"fmt"
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"time"
)

type Topic struct {
    Id int
    Title string
    Content string
    Author string `sql:"not null"`
    Comments []Comment
    CreatedAt time.Time
}

type Comment struct {
    Id int
    Content string
    Author string `sql:"not null"`
    TopicId int `sql:"index"`
    CreatedAt time.Time
}

var DbConn *gorm.DB

func init()  {
    var err error
    DbConn, err = gorm.Open("mysql", "root:root@/gwp?charset=utf8mb4&parseTime=true")
    if err != nil {
        panic(err)
    }
    DbConn.AutoMigrate(&Topic{}, &Comment{})
}

func main()  {
    topic := Topic{Title: "Test 1", Content: "Test Content", Author: "学院君"}
    fmt.Println(topic)

    DbConn.Create(&topic)
    fmt.Println(topic)

    comment := Comment{Content: "Test Comment", Author: "学院君小号"}
    DbConn.Model(&topic).Association("Comments").Append(comment)

    var readTopic Topic
    DbConn.Where("author = ?", "学院君").First(&readTopic)

    var comments []Comment
    DbConn.Model(&readTopic).Related(&comments)
    fmt.Println(comments[0])
}
