docker 发布 第一个 GO
==============================

1.编写一个 main.go
---------------------------------
所有的内容都在同一个目录下::

        package main
        import (
                "fmt"                                                                                                                                                                
                "log"
                "net/http"
        )
        func sayHello(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "hello world")  
        }
        func main() {
                http.HandleFunc("/", sayHello) //注册URI路径与相应的处理函数
                log.Println("【默认项目】服务启动成功 监听端口 80")
                er := http.ListenAndServe("0.0.0.0:80", nil)
                if er != nil {
                        log.Fatal("ListenAndServe: ", er)
                }
        }

编译(我是 mac 环境需要制定运行平台)::

        GOARCH=amd64 GOOS=linux go build main.go）

2.创建一个 Dockerfile
---------------------------------

内容如下::

        FROM golang

        MAINTAINER Razil "zc6496359"

        WORKDIR $GOPATH/src/godocker

        ADD . $GOPATH/src/godocker

        RUN go build main.go

        EXPOSE 80

        ENTRYPOINT ["./main”]

参数解释::
        FROM -> 母镜像 

        MAINTAINER -> 维护者信息 

        WORKDIR -> 工作目录 

        ADD -> 将文件复制到镜像中 

        RUN -> 执行操作（就跟在终端执行语句一样） 

        EXPOSE -> 暴露端口 

        ENTRYPOINT -> 程序入口

3.编译镜像
----------------------------------
docker build -t zcdocker .

出现Successfully build … 说明成功

4.运行并创建一个容器
---------------------------------

docker run -p 80:80 -d zcdocker
-p   本机端口:镜像端口 
-d    后台运行

5.测试 OK

curl -v http://localhost:8080