package main

import (
    "testing"
    "time"
)

func TestPrint(t *testing.T)  {
    print()
}

func TestGoPrint(t *testing.T)  {
    goPrint()
    time.Sleep(1 * time.Millisecond)  // 延迟10毫秒退出，等待所有子协程执行完毕
}

func BenchmarkPrint(b *testing.B) {
    for i := 0; i < b.N; i++ {
        print()
    }
}

func BenchmarkGoPrint(b *testing.B)  {
    for i := 0; i < b.N; i++ {
        goPrint()
    }
}
