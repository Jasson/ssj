Go mapinterface-interface-json 转换
########################################################

此笔记来自于对 interface 如何转换为 struct 疑问,以至于研究了一下 Go 的反射，甚至基于反射去实现一些功能（也不能说错误就是造轮子了，没必要，还是不必要的轮子）

网上查到有人用 把读出的 interface转为 Json，在把 Jason 转为 map-interface



来自 `这里有详细的转换例子 <https://blog.csdn.net/xiaoquantouer/article/details/80233177>`_

已知的 interface 转为 map-interface或 struct 用类型断言转换即可