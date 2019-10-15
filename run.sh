#!/usr/bin/env bash

// docker stop `docker ps -a |grep ssj |awk '{print $1}'` && docker rm `docker ps -a |grep ssj |awk '{print $1}'`
// docker run -p 80:80 -d registry.cn-beijing.aliyuncs.com/langzhe/ssj:latest
 docker logs `docker ps  |grep ssj |awk '{print $1}'`
