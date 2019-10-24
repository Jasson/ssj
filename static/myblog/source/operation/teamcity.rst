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