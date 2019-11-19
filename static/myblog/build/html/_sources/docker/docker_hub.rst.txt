docker hub发布项目
#################################

把 docker hub 当成 git hub 理解用，哦哦
https://hub.docker.com/

上传镜像
-------------

::

	docker login
	docker commit -a "langxianzhe" -m "test” docker2.1 langxianzhe/docker2.1:0.1
	docker push langxianzhe/docker2.1:0.1
	dock login 输入账号密码

拉取镜像
---------------

::

	docker
	docker pull langxianzhe/docker2.1:0.1
	docker images
	docker run -it docker2.1:v.1