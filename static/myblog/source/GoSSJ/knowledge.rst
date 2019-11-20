Go 遇到的错误
########################################



类型转换
----------------------------------------------------------------------------------

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

