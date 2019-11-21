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

time用法 ::

readTimeOut:=1
var a time.Duration
a = readTimeOut  提示错误

time.Duration 是 int64
a:=int64(readTimeOut)也是错误
 time.Duration(writeTimeOut)//这个 OK

