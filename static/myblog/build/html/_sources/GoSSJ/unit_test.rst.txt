Go 单元测试
############

之前写 Erlang 的单元测试代码顺风顺水，理所当然以为 Golang 的也很简单，事实也是很简单，实际学习过程中比预期的繁琐，花的时间也多。网上搜了好多例子基本都是前边一律中规中矩的 Hello World 轻松上手，一旦涉及项目应用就开始挠头捉耳了。各种达不到预期。

下图右击 执行某个测试函数

.. figure:: /images/goruntest.png
   :width: 180%


如何一个命令执行所有Go的测试代码？
-----------------------------------------

我去，就么一个问题折磨了我好久，原因还是对 Go的不是熟悉，没有系统认真学学 Go,只是上手就用。想起了自己常说的话“路没有捷径，看似是捷径，行走过程中会渐渐变成弯道”
网上搜并没有很容易搜出结果来.

运行所有的单元测试命令::

    go test ./...
    The second, called package list mode, occurs when go test is invoked
    with explicit package arguments (for example 'go test math', 'go
    test ./...', and even 'go test .'). In this mode, go test compiles
    and tests each of the packages listed on the command line. If a
    package test passes, go test prints only the final 'ok' summary
    line. If a package test fails, go test prints the full test output.
    If invoked with the -bench or -v flag, go test prints the full
    output even for passing package tests, in order to display the
    requested benchmark results or verbose logging.
    An import path is a pattern if it includes one or more "..." wildcards,
    each of which can match any string, including the empty string and
    strings containing slashes. Such a pattern expands to all package
    directories found in the GOPATH trees with names matching the
    patterns.

    To make common patterns more convenient, there are two special cases.
    First, /... at the end of the pattern can match an empty string,

运行当前目录单元测试命令::

    go test ./
    go test -v ./xxx/xxx/*go


Go 单元测试基本知识
-------------------------

Golang单元测试对文件名和方法名，参数都有很严格的要求。::

    1.单元测试文件名必须以xxx_test.go命名
    2. 方法必须是TestXxx开头
    3. 方法参数必须 t *testing.T
    4. 测试文件和被测试文件必须在一个包中

标记单元测试结果
-----------------------

终止 t.FailNow()

只标记错误不终止 t.Fail()


Go 单元测试进阶命令
------------------------

命令1：go help test

命令2: go help packages


补充：ginkgo 工具
------------------------

研究单元测试时走了好多弯路，其中查到 ginkgo工具。也花了不少时间。搞明白后还是很简单的。以下代码是样例copy 后是跑不起来的

.. code-block:: go
   :emphasize-lines: 3,5

        package data_analysis

        import (
            "fmt"
            . "github.com/onsi/ginkgo"
            . "github.com/onsi/gomega"
            "octopus_web/conf"
            "testing"
        )

        func TestCart(t *testing.T) {
            RegisterFailHandler(Fail)
            RunSpecs(t, "DataAnalysis Suite")
        }
        var _ = Describe("./Routers/Api/V2/DataAnalysis/DataModels", func() {
            Context("initially", func() {
                conf.Init("../../../../conf/app.ini")
                Init()
                ga := &GadgetAttribute{}
                err := ga.getOneByID("1000bae96ca731748b795c221897885e")
                fmt.Println("err:=", err)
                //fmt.Printf("ga:=%#v\n", ga)
                By("Fetching list list:=")
                //log.Println("list:=", list)

                It("add", func() {
                    Expect(Add(1,1)).Should(Equal(2))
                })
                It("add", func() {
                    Expect(ga).Should(Not(Equal(nil)))
                })
                It("getAll", func() {
                    Expect(err).Should(BeNil())
                })

                var list DataList
                rlist, err := list.GetAttribute()
                //log.Println("list:=", list)

                It("getAttribute", func() {
                    Expect(err).Should(BeNil())
                })
                Write(rlist)

            })

        })


我只想执行一个查询数据的语句，（单元测试涉及到 conf,DB 时需要注意的问题）
------------------------------------------------------------------------------

在测试 DB是，配置文件总是读取错误，这一点与 Erlang 有很多区别，例如在 Erlang 里面只要在根目录配置后，启动加载即可。Go 里面 不可以，会在执行文件所在的目录开始寻找。

例如:   
配置文件录 ::  

    tree conf/
    conf/
    ├── app.ini

    └── conf.go

routers和 conf在同一个目录下，执行文件在 ::

    routers/api/v2/data_analysis/

初始化文件要如下方式传递参数 ::

    conf.Init("../../../../conf/app.ini")//这个问题也花了非常多时间来研究 主要是对 Go 欠熟悉。
    conf.Init("../../../../conf/app.ini")//这个问题也花了非常多时间来研究 主要是对 Go 欠熟悉。

代码：

.. code-block:: go
   :emphasize-lines: 3,5

    func init() {
        var err error
        Cfg, err = ini.Load("conf/app.ini") //初始化配置文件
        if err != nil {
            log.Fatal("Loade config err")
        }
        log.Println(Cfg)
    }

        // 加载配置文件
    func Init(path string) {
        log.Println("path:=", path)
        var err error
        if path == "" {
            Cfg, err = ini.Load("conf/app.ini")
        } else {
            Cfg, err = ini.Load(path)
        }
        if err != nil {
            log.Fatal("Load config err:=", err, "path:=", path)
        }
        log.Println(Cfg)
    }

   






