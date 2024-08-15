package main

import (
    "embed"
    "net/http"
    "log"
    "io/fs"
)

// 使用embed嵌入静态文件目录
//go:embed build/*
var embededFiles embed.FS

func main() {
    // 创建一个http.FileSystem，从嵌入的文件系统中读取文件
    staticFiles, err := fs.Sub(embededFiles, "build")
    if err != nil {
        log.Fatal(err)
    }
    fs := http.FS(staticFiles)

    // 提供静态文件服务
    http.Handle("/", http.FileServer(fs))

    log.Println("服务器正在运行，监听端口 :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}