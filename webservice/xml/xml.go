package main

import (
    "encoding/xml"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

type Post struct { //#A
    XMLName xml.Name `xml:"post"`
    Id string `xml:"id,attr"`
    Content string `xml:"content"`
    Author Author `xml:"author"`
    Xml string `xml:",innerxml"` // 用于匹配 post 元素中原生 XML 内容
    Comments []Comment `xml:"comments>comment"`
}

type Author struct {
    Id string `xml:"id,attr"`
    Name string `xml:",chardata"`
}

type Comment struct {
    Id string `xml:"id,attr"`
    Content string `xml:",chardata"`
    Author Author `xml:"author"`
}

func parseSmallXML(r io.Reader, post *Post) (err error)  {
    xmlData, err := ioutil.ReadAll(r)
    if err != nil {
        fmt.Println("Error reading XML data:", err)
        return
    }

    xml.Unmarshal(xmlData, post)
    return
}

func parseLargeXML(r io.Reader)  {
    decoder := xml.NewDecoder(r)
    for {
        t, err := decoder.Token()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Error decoding XML into tokens:", err)
            return
        }
        switch se := t.(type) {
        case xml.StartElement:
            if se.Name.Local == "comment" {
                var comment Comment
                decoder.DecodeElement(&comment, &se)
            }
        }
    }
}

func main()  {
    /*xmlFile, err := os.Open("post2.xml")
    if err != nil {
        fmt.Println("Error opening XML file:", err)
        return
    }
    defer xmlFile.Close()*/

    // 使用 Unmarshal 解析小文件
    //var post Post
    //parseSmallXML(xmlFile, &post)
    //fmt.Println(post)

    // 使用 Decoder 解析大文件
    //parseLargeXML(xmlFile)

    author := Author{Id: "1", Name: "学院君"}
    var comments []Comment
    comment1 := Comment{Id:"1", Content:"Comment1", Author: author}
    comment2 := Comment{Id:"2", Content:"Comment2", Author: author}
    comments = append(comments, comment1, comment2)

    post := Post{
        Id: "1",
        Content: "Hello XML",
        Author: author,
        Comments: comments,
    }

    /*xmlData, err := xml.MarshalIndent(&post, "", "\t")
    if err != nil {
        fmt.Println("Error marshalling to XML:", err)
        return
    }

    err = ioutil.WriteFile("post3.xml", []byte(xml.Header + string(xmlData)), 0644)
    if err != nil {
        fmt.Println("Error writing XML to file:", err)
        return
    }*/

    xmlFile, err := os.Create("post4.xml")
    if err != nil {
        fmt.Println("Error creating XML file:", err)
        return
    }
    xmlFile.WriteString(xml.Header)

    encoder := xml.NewEncoder(xmlFile)
    encoder.Indent("", "\t") // 缩进，这一步不是必须的
    err = encoder.Encode(&post)
    if err != nil {
        fmt.Println("Error encoding XML to file:", err)
        return
    }
}

