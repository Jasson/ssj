Go mapinterface-struct-assertions与断言
########################################################

类型转换
使我纠结了好久的断言


Go 中类型断言 VS Java,Erlang断言
---------------------------------------------

Go 里面叫类型断言不是断言，Java 和 Erlang编程中常用断言的地方都是在单元测试中,都有醒目的关键字“assert”参与。

当然在 Erlang 里面 guard 也叫断言，好多函数及 BIF 函数都是叫断言，常用就不自知（is_binary(),is_float(),node()，hd()等）


Go 类型断言
--------------------

在go 里面的类型断言与在 Java 和 Erlang 里面的断言相比不一回事。
Go 的断言 这样叫先入为主错误的叫法。原名叫“类型断言（Typeassertions）”
类型断言 提供了访问接口值底层具体值的方式。
它的关键就是判断类，并把值赋予新变量

 ::

  A type assertion provides access to an interface value's underlying concrete value.
  t := i.(T)


该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
第一次看到这里（https://tour.go-zh.org/methods/15） 例子也看了，也写了，实际使用的时候压根不记得这回事。
重要的作用就是把 interface 转换成你想要的数据。 好比 Java Object class 转换成子类，类型断言粗暴的翻译没有表达出原来的含义，表达出类型转换回去好了
表达式 T(v) 将值 v 转换为类型 T。

应用-实践
------------

写 Erlang 写久了都忘记还有类型转换这回事了。特别是定义函数时压根不考虑类型，直接定义变量使用就行，传 record还是 list，aotm, tuple，只要使用的时候能清楚的知道就行，感觉 record不爽直接换成传 list 也是可以的，当然调用地方要传递参数也要修改成 list。这都不重要。
重要的是返回结果可以是任意类型的，返回直接使用就行,如果你愿意一个函数可以返回 list，另一些情况返回 record，完全是允许的。“所见即所得”，说一不二，有什么说（返回）什么，一直用一直爽。直到遇到Go 一下子爽不起来了。返回的都是 interface 需要转换。
Go 实践中遇到
第一个问题: interface 如何转换为 struct
第一次遇到没搞定，查到了 interface 转 map-interface
把变量放入 context

::

  sid := make(map[string]interface{})
  sid["user_id"] = "langxw1"
  sid["sid"] = "langsid"
  content.Set("sid", sid)
 
  //log.Println("sid:=", sid)
  if val, ok := content.Get("sid"); ok {
    log.Println("val:=", val)
    sidMap, err := cast.ToStringMapE(val)
    if err != nil {
       log.Println(err)
     }
    userId := sidMap["user_id"].(string)
    log.Println("userId:=", userId)
  }
  再通过 userId := sidMap["user_id"].(string) 取出 user_id
  这里已经用了过的断言 type assertion，当初无知
 
  实际上完全 可以用"sidMap := val.(map[string]interface{})”取代 cast.ToStringMapE
  进一步了解 用 map-interface 可读写比较差，当初是无知采用的，现在回过够来重新写
  type SID struct {
   UserId string `json:"user_id"`
   Sid    string `json:"sid"`
  }

  sid := SID{"langxw", "sid"}
  content.Set("sid", sid)

  if val, ok := content.Get("sid"); ok {
     sidStruct := val.(SID)
     log.Println("sid:=", sidStruct)
  }


哎，砍柴不误磨刀工，没有磨刀就去砍柴了，导致路途坎坷，沾了一身泥，不得不回头重新来
前期因为解决 interface 转换 map-interface 和 struct 看了好多反射的内容，及找到了 https://github.com/goinggo/mapstructure 
mapstructure is a Go library for decoding generic map values to structures and vice versa, while providing helpful error handling.
This library is most useful when decoding values from some data stream (JSON, Gob, etc.) where you don't quite know the structure of the underlying data until you read a part of it. You can therefore read a map[string]interface{} and use this library to decode it into the proper underlying native Go structure.

 