TeamCity 之旅
######################

TeamCity提供一系列特性可以让团队快速实现持续集成：IDE工具集成、各种消息通知、各种报表、项目的管理、分布式的编译等等，所有的这些，都是让你的团队快速享有持续集成带来的效率提升、高质量的软件保障。
使用 TeamCity，你能够在几分钟之内为你的项目配置一个构建服务器，它内建了持续单元测试，代码质量分析和早期的构建问题分析报告，你甚至可以在IDE进行。
TeamCity 提供平滑的学习曲线，你可以逐步的学习经它的高级特性和功能，你很快就能加强你发布管理实践。 本次发布，在可用性作了大量的改进，更新的IDE 插件支持 CVS和SVN，另外还包括一些之前版本不具备的企业级的特性。

.. figure:: /images/teamcity.png
   :width: 180%


安装
------------------------------------------------

用Docker Image 安装 Server端::
	
	docker run -it --name teamcity-server-instance  \
    	-v /opt/teamcity/data:/data/teamcity_server/datadir \
    	-v /opt/teamcity/logs:/opt/teamcity/logs  \
    	-p 8111:8111 \
    	jetbrains/teamcity-server 

    出现此提示时表示初步安装成功，接下来进入后台进行配置

	Startup confirmation is required. Open TeamCity web page in the browser. Server is running at http://localhost:8111

数据库选择::

	支持 Mysql, Oracle, PostgreSQL,HSQLDB,这里选用的是默认数据.

.. figure:: /images/teamcity.png
   :width: 180%

初次进入会提示创建项目，后续需要主动创建项目.这里提供了简便的用户密码方式授权，同时也提供了 Oauth 方式

.. figure:: /images/teamcityinit.png
   :width: 180%


.. figure:: /images/teamcitygit.png
   :width: 180%

根据提示选择 Proceed,此步骤会根据工程目录自动选择可能的Build Steps


.. figure:: /images/teacitybutilstep.png
   :width: 180%

首次接触误以为 直接run可以了。会提示没有可用的 agent.

.. figure:: /images/teamcityagent.png
   :width: 180%

.. figure:: /images/teamcityagent1.png
   :width: 180%

安装 Agent，每一台要发布的机器都要安装 agent,有没有批量安装方法呢？用 K8S 环境下如何部署？::

	docker run -p 9090:9090 -it -e SERVER_URL="http://120.27.8.120:8111"  \
    -v /data/teamcity_agent/conf:/data/teamcity_agent/conf \
    -v /var/run/docker.sock:/var/run/docker.sock  \
    -v /opt/buildagent/work:/opt/buildagent/work \
    -v /opt/buildagent/temp:/opt/buildagent/temp \
    -v /opt/buildagent/tools:/opt/buildagent/tools \
    -v /opt/buildagent/plugins:/opt/buildagent/plugins \
    -v /opt/buildagent/system:/opt/buildagent/system \
    jetbrains/teamcity-agent

 注意::

 	Image 安装要注意制定本地目录，否则 Agent 虽然能 connected但是 build的时候找不到。为了省去不必要的麻烦端口号都有点 9090。


.. figure:: /images/teamcityagent2.png
   :width: 180%


angent安装出现 authorization token 时，复制 token回到 web 管理界面进行 agent 授权。


.. figure:: /images/teamcityagent3.png
   :width: 180%


.. figure:: /images/teamcityagent4.png
   :width: 180% 

授权完会进入自动build


.. figure:: /images/teamcityagent5.png
   :width: 180% 


.. figure:: /images/teamcityagent6.png
   :width: 180% 



到目前为止 整个 TeamCity 就安装完了::
	
	上面用到了 Token 授权 Server 访问 Agent,在安装 Server 时也曾出现过用 Token 代替用户和密码登陆，后续再要求创建管理员账号。

配置自动化发布
----------------------

设置 Build Steps ,此例子总共6 步::
	
	build, login, push, stop rm, delete image, run.
	Login主要是为了登陆阿里的容器仓库，实际上在非容器化部署的 agent 中只执行一次即可，主机会记住登陆信息。 


.. figure:: /images/teamcityconfig1.png
   :width: 180% 

Buid::

	创建 Build,比较简单 Runner Type 选择 Docker 即可。 给 Step 命名为 Build(随便命名可读即可)，Docker command 选择 build,
	Path to file默认 TeamCity 会读取根目录下的 Dockerfile，也可以指定其他目录的 Dockerfile。
	Image name:tag 根基实际情况指定。这里是用的阿里的registry.cn-beijing.aliyuncs.com/langzhe/ssj。


.. figure:: /images/teamcitybuild.png
   :width: 180% 


Login::

	授权允许访问仓库,Docker command 中没有 login，这里选择 other,在 Command name 中输入 login.
	在 Additional arguments 中数据用户名密码和地址。

.. figure:: /images/teamcitylogin.png
   :width: 180% 

Push  把镜像推到阿里镜像仓库，这个配置比较简单

.. figure:: /images/teamcitypush.png
   :width: 180% 

stop rm::
	测试的时候发现经常 Image 不更新，这里在 run 时，先把容器停掉并删除，同时把 Image 删除，run 的时候用的 latest.需要在主机上执行命令Runner Type要选择 CommandLine ,Run:选择 Custom script并输入以下脚本：

.. figure:: /images/teamcitystop.png
   :width: 180% 

删除 Image，与第四步类似，脚本如下::

	docker rmi `docker images |grep ssj |awk '{print $3}'` 


.. figure:: /images/teamcityremoveimage.png
   :width: 180% 

第六步 run,与第二步类似::
	
	DockerComand选择 Other,在 Command name 中输入 run.
	Additional arguments for the command:输入-p 80:80 -d  registry.cn-beijing.aliyuncs.com/langzhe/ssj

.. figure:: /images/teamcityrun.png
   :width: 180% 

然后，执行 Run，这个过程需要反复调试。可以禁用一些 Steps.

最后一定要设置 Trigers

.. figure:: /images/teamcitytrigger1.png
   :width: 180% 

.. figure:: /images/teamcitytrigger2.png
   :width: 180% 









