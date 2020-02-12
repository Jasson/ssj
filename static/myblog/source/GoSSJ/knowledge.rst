Go 遇到的错误
#############



类型转换
--------

类型转换::

	if !(ok && token.Valid) {
      return nil, errors.New("非法的token")
  	 }

  	 jwtToken := &JwtToken{
  	    Uid: claims["uid"].(string),
    	  Sid: claims["sid"].(string),
      	Iat: int64(claims["iat"].(float64)),
      	Exp:int64(claims["exp"].(float64)),
   		}
   		return jwtToken, nil
	}

整型转换为字符串 类型转换丢失数据::

  stirng(u.user_id)丢数据
  用下面方法解决
  userId  :=  fmt.Sprintf("%d", u.UserId)



时间相关
--------
time用法::

	readTimeOut:=1
	var a time.Duration
	a = readTimeOut  提示错误
	time.Duration 是 int64
	a:=int64(readTimeOut)也是错误
	 time.Duration(writeTimeOut)//这个 OK


map与数字相关的用法
-------------------

查看 map debug 内容
第一次 debug 看到 map 显示的内容困惑了我好长一段时间，因为前面有个索引数字 0-N，我误认为是数组

.. code-block:: go

    package main
    import "log"

    func main()  {
       lxw := make(map[string]string)
       lxw["aaa"]="xxx"
       lxw["value"]="valu"
       log.Println(lxw)
    }

.. image:: /images/mapstring.png


mongo-go-driver 连接错误
-------------------------

连接地址::

  10.128.132.16:27017

错误信息::
  
  err:= server selection error: server selection timeout, current topology: { Type: ReplicaSetNoPrimary, Servers: [{ Addr: 10.140.2.18:27017, Type: Unknown, State: Connected, Average RTT: 0, Last error: connection() : dial tcp 10.140.2.18:27017: i/o timeout }, ] }

解决方法 增加 connect=direct 参数::
  
  10.128.132.16:27017/?connect=direct
  Regarding your comment on why using mongo shell, or PyMongo works.

  This is due to the difference in connection mode. When specifying a single node, i.e. mongodb://node1:27017 in shell or PyMongo, server discovery are not being made. Instead it will attempt to connect to that single node (not as part as a replica set). The catch is that you need to connect to the primary node of the replica set to write (you have to know which one). If you would like to connect as a replica set, you have to define the replica set name.

  In contrast to the mongo-go-driver, by default it would perform server discovery and attempt to connect as a replica set. If you would like to connect as a single node, then you need to specify connect=direct in the connection URI.


`参考文档  <https://ask.csdn.net/questions/1019794>`_ 