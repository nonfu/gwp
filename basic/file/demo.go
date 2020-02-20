package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main()  {
    data := []byte("Hello World!")
    err := ioutil.WriteFile("data1", data, 0644)
    if err != nil {
        panic(err)
    }
    read1, _ := ioutil.ReadFile("data1")
    fmt.Println(string(read1))

    file1, _ := os.Create("data2")
    defer file1.Close()

    bytes, _ := file1.Write(data)  // 将数据写入file1并返回写入的字节数
    fmt.Printf("wrote %d bytes to file\n", bytes)

    file2, _ := os.Open("data2")
    defer file2.Close()

    read2 := make([]byte, len(data))
    bytes, _ = file2.Read(read2) // 从file2读取数据到read2并返回读取的字节数
    fmt.Printf("read %d bytes from file\n", bytes)
    fmt.Println(string(read2))
}
