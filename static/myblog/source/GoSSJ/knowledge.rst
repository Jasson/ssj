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



