#!/usr/bin/env bash
#rm -rf config
#cp -a ../config ./
#以下命令在项目的 k8s 目录下执行
#生成二进制包
GOARCH=amd64 GOOS=linux go build -o ssj ../main.go
#制作镜像
docker build -t registry.cn-beijing.aliyuncs.com/langzhe/ssj:v1.1 .
#上传镜像
docker push registry.cn-beijing.aliyuncs.com/langzhe/ssj:v1.1
