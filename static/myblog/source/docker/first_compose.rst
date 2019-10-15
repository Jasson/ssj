docker-compose 发布 第一个 GO
######################################

1.编写一个 main.go
----------------------

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

  GOARCH=amd64 GOOS=linux go build main.go

2.创建一个 Dockerfile
--------------------------
内容如下：::

  FROM golang                                                                                                                                                                  
  MAINTAINER langxw "langxw"
  WORKDIR $GOPATH/src/godocker
  COPY . $GOPATH/src/godocker
  EXPOSE 80
  ENTRYPOINT ["./build.sh"]


3.创建  build.sh
-----------------------
::

  #!/usr/bin/env bash
  cd $GOPATH/src/godocker && ./main

3.编译镜像：
------------

docker build -t docker2.1 .

出现Successfully build … 说明成功

4.创建一个docker-compose.yml 
---------------------------------
::

    1 version: '2'                                 
    2 
    3 networks:
    4   basic:
    5 
    6 services:
    7 
    8   docker2.1:
    9     container_name: docker2.1
   10     image: docker2.1
   11     ports:
   12       - "80:80"
   13     networks:
   14       - basic

4.运行并创建一个容器
----------------------
docker-compose -f docker-compose.yml up -d docker2.1


5.测试 OK
----------

curl -v http://localhost:8080