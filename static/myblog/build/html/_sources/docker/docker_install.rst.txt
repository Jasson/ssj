Docker 安装
####################

以下文字是之前收集到印象笔记，没有记录出处，若有异议联系我删除(wchat:langxianzhe)：

2017 年 3 月开始 docker 在原来的基础上分为两个分支版本: Docker CE 和 Docker EE。
Docker CE 即社区免费版，Docker EE 即企业版，强调安全，但需付费使用。
本文介绍 Docker CE 的安装使用。
移除旧的版本::

	$ sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-selinux \
                  docker-engine-selinux \
                  docker-engine

安装一些必要的系统工具::

	sudo yum install -y yum-utils device-mapper-persistent-data lvm2

添加软件源信息::

	sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

更新 yum 缓存::

	sudo yum makecache fast

安装 Docker-ce::

	sudo yum -y install docker-ce

启动 Docker 后台服务::

	sudo systemctl start docker

测试运行 hello-world::

	[root@runoob ~]# docker run hello-world


由于本地没有hello-world这个镜像，所以会下载一个hello-world的镜像，并在容器内运行。


镜像加速
------------------

鉴于国内网络问题，后续拉取 Docker 镜像十分缓慢，我们可以需要配置加速器来解决，我使用的是网易的镜像地址：http://hub-mirror.c.163.com。
新版的 Docker 使用 /etc/docker/daemon.json（Linux） 或者 %programdata%\docker\config\daemon.json（Windows） 来配置 Daemon。
请在该配置文件中加入（没有该文件的话，请先建一个)::

	{
	  "registry-mirrors": ["http://hub-mirror.c.163.com"]
	}


删除 Docker CE
---------------------

执行以下命令来删除 Docker CE::
	$ sudo yum remove docker-ce
	$ sudo rm -rf /var/lib/docker

docker
-----------------------

运行 docker 容器::

	docker run -t -i ubuntu:14.04 /bin/bash
	docker run -i -t  quirky_stallman /bin/bash
	docker run -p 80:80 -d zcdocker

查看容器 log::

	docker logs world

删除 Image::

	docker rmi删除 Imgage

删除 容器::

	docker rm 删除容器


利用容器 ID 或名称进入 容器::

	docker exec -it da8283abd851 sh 
	docker exec -it quirky_stallman sh


批量删除 imgage::

	docker rmi `docker images |grep 781MB`


批量删除 容器::

	docker rm `docker ps -a |awk '{print $1}' | grep [0-9a-z]`

进入容器 shell::
	
	docker exec -it 3412c530d267 sh


停止状态为Exited的容器::

	docker ps -a | grep "Exited" | awk '{print $1 }'|xargs docker stop

删除状态为Exited的容器::

	docker ps -a | grep "Exited" | awk '{print $1 }'|xargs docker rm

删除 tag名称为none的image::

	docker images|grep none|awk '{print $3 }'|xargs docker rmi
	docker rmi $(docker images|grep none|awk '{print $3 }')


杀死所有正在运行的容器::

	docker kill $(docker ps -a -q)

删除所有已经停止的容器::

	docker rm $(docker ps -a -q)

删除所有未打 dangling 标签的镜像::

	docker rmi $(docker images -q -f dangling=true)

删除所有镜像::

	docker rmi $(docker images -q)


为这些命令创建别名

# ~/.bash_aliases::

	# 杀死所有正在运行的容器.
	alias dockerkill='docker kill $(docker ps -a -q)'

	# 删除所有已经停止的容器.
	alias dockercleanc='docker rm $(docker ps -a -q)'

	# 删除所有未打标签的镜像.
	alias dockercleani='docker rmi $(docker images -q -f dangling=true)'

	# 删除所有已经停止的容器和未打标签的镜像.
	alias dockerclean='dockercleanc || true && dockercleani'


Q&A
----------------------------

镜像删除失败::

	第一个镜像直接docker rmi 74f5aea45cf6就删除成功了，然而后面两个镜像都是成对出现，直接docker rmi删除失败，报错信息如下：
	Error response from daemon:
	conflict: unable to delete b3cd25f64a2a (must be forced) - image 
	is referenced in multiple repositories

	解决方案:

	首先docker rmi时指定名称，而非image id，然后再执行docker rmi -f image idj即可
	docker rmi paddlepaddle/paddle:1.1.0-gpu-cuda8.0-cudnn7
	docker rmi -f b3cd25f64a2a

