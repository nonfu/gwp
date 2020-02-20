package main

import (
"bytes"
"encoding/gob"
"fmt"
"io/ioutil"
)

type Article struct {
    Id int
    Title string
    Content string
    Author string
}

func store(data interface{}, filename string)  {
    buffer := new(bytes.Buffer)
    encoder := gob.NewEncoder(buffer)
    err := encoder.Encode(data)
    if err != nil {
        panic(err)
    }
    err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
    if err != nil {
        panic(err)
    }
}

func load(data interface{}, filename string) {
    raw, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    buffer := bytes.NewBuffer(raw)
    dec := gob.NewDecoder(buffer)
    err = dec.Decode(data)
    if err != nil {
        panic(err)
    }
}

func main()  {
    article := Article{Id: 1, Title: "Article 1", Content: "Hello World!", Author: "学院君"}
    store(article, "article1")
    var articleRead Article
    load(&articleRead, "article1")
    fmt.Println(articleRead)
}
