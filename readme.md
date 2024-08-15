#假设你有一个名为 main.go 的 Go 文件，以下命令将编译它为适用于 Linux 64 位的可执行文件：
GOOS=linux GOARCH=amd64 go build -o reactgo main.go

GOOS=linux：指定目标操作系统为 Linux。
GOARCH=amd64：指定目标架构为 64 位（AMD64，即 x86_64）。
-o reactgo：指定输出的可执行文件名为 reactgo。

dpkg-deb --build reactgo
 dpkg -i reactgo.deb

 可以使用-ldflags参数来去除符号信息，减少可执行文件的大小。
 go build -ldflags "-s -w" -o reactgo main.go

 
在Linux上，你可以使用以下命令构建静态链接的可执行文件。
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o reactgo main.go

示例：从Linux编译为Windows
GOOS=windows GOARCH=amd64 go build -o reactgo.exe main.go