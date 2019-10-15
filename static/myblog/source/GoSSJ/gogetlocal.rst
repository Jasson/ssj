Go get 本地库
######################

习惯了 Erlang rebar 的 get deps，还以为 go get 也理所当然 go get 本地库即可，事实稍微有点麻烦，花了一下午时间没有解决效率有点低。网上好多方法，有好多年前的也有最近方法，也有代理解决的。


go get 获取默认是 ssl 验证的
------------------------------------------------

使用  -insecure  参数来获取 http的
如::

    go get -v   -insecure  gitlab.com/tools/auth




parsing go.mod: unexpected module path "auth"
------------------------------------------------------------------------------------------------

go: gitlab.com/tools/testid_auth@v0.0.1: parsing go.mod: unexpected module path "auth"
go: error loading module requirements

大部分时间都花在解决这问题上了，一只没明白什么意思，这主要是对 go.mod 理解不深刻原因造成的。

原先 auth工程中的 go.mod 内容为::

    auth
    go 1.12
    require github.com/gin-gonic/gin v1.4.0

修改为以下解决此问题:: 

    module gitlab.com/tools/auth
    go 1.12
    require github.com/gin-gonic/gin v1.4.0

问题解决好才恍然大悟 错误提醒还是非常明显的。只是自己没有理解，然后到 Google 一桶搜索，然后一步步都搞不定。


另外几个错误
------------------------------------------------------------------------------------------------

这几个错误是因为胡乱配导致导致的::

    Parsing meta tags from http://gitlab.com?go-get=1 (status code 200)
    go get gitlab.com/tools/auth: git ls-remote -q http://gitlab.com/tools/auth.git in /Users/jason/go/pkg/mod/cache/vcs/4bd4058a728ae179ff04381faa4fc658a16fa8a62c9c6d8aaa6288cf645c9cf0: exit status 128:
        fatal: 'git@gitlab.com/tools/auth.git' does not appear to be a git repository
        fatal: 无法读取远程仓库。
        
        请确认您有正确的访问权限并且仓库存在。

    git ls-remote -q http://gitlab.com/tools/auth.git in /Users/jason/go/pkg/mod/cache/vcs/4bd4058a728ae179ff04381faa4fc658a16fa8a62c9c6d8aaa6288cf645c9cf0: exit status 128:
        fatal: could not read Username for 'http://gitlab.com': terminal prompts disabled


    Parsing meta tags from http://gitlab.com/tools/auth?go-get=1 (status code 200)
    get "gitlab.com/tools/auth": found meta tag get.metaImport{Prefix:"gitlab.com/tools/auth", VCS:"git", RepoRoot:"http://gitlab.com/tools/auth.git"} at http://gitlab.com/tools/auth?go-get=1
    jason@gitlab.com's password: 


解决方法::
    
    正常配置 OK，我手工修改时 把git@gitcodecloud.test.com.cn：后面冒号删除了
    [merge]
         tool = vimdiff
    [url "git@gitcodecloud.test.com.cn:"]                               
             insteadof = http://gitcodecloud.test.com.cn 

可以按部就班的用以下命令设置::

    git config --global http.extraheader "PRIVATE-TOKEN: YOUR_PRIVATE_TOKEN"
    git config --global url."git@gitlab.com:groupName/projectName.git".insteadOf "https://gitlab.com/groupName/projectName.git"`
    或用
    git config --global url."git@gitlab.yoursite.com:".insteadof "https://gitlab.yoursite.com/"
    配网就可以拉 代码了go get -u -v gitlab.com/groupName/projectName

来自 `zhoushuaime简书说 <https://www.jianshu.com/p/ca4404512cf3>`_ 
的说法 ssh 转为 http::

     git config --global url."git@gitlab.com:groupName/projectName.git".insteadOf "https://gitlab.com/groupName/projectName.git"`



来自 `golangdoc <https://golang.org/doc/faq#git_https>`_

Git can also be configured to use SSH in place of HTTPS for URLs matching a given prefix. For example, to use SSH for all GitHub access, add these lines to your ~/.gitconfig:::

    [url "ssh://git@github.com/"]
        insteadOf = https://github.com/


本人没有配置 http.extraheader 啰嗦了那么多就是通过两个命令解决问题的::

    1.添加  -insecure 
        go get -v insecure gitlab.com/tools/autgitlabh
    2. git config --global url."git@gitlab.yoursite.com:".insteadof "https://gitlab.yoursite.com/"

若是设置了代理 export GOPROXY=https://goproxy.io 一定要关掉
在执行如::

    export GOPROXY=
    go get -u -v gitlab.com/groupName/projectName


