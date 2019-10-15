docker-compose.yml 解释
#############################################

简要说明：::
	
	  3 networks:
	  4   basic:
	  5 
	  6 services:
	  7 
	  8   world:
	  9     container_name: world
	 10     image: go-web 指定镜像，如不存在，则利用下面 build，根据当前目录下的 Dockefile 编译一个镜像，Dockefile 下面设置自动获取过来
	 11     build: 
	 12         context: ./
	 13     ports:
	 14       - "80:80"
	 15     volumes:
	 16       - ./app/go/world:/go/src/app:rw
	 17     networks:
	 18       - basic
	  3 networks:
	  4   basic:
	  5 
	  6 services:
	  7 
	  8   world:
	  9     container_name: world
	 10     image: go-web
	 11     build: 
	 12         context: ./
	 13     ports:
	 14       - "80:80"
	 15     volumes:
	 16       - ./app/go/world:/go/src/app:rw
	 17     networks:
	 18       - basic
