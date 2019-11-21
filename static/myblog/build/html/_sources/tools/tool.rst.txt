常用工具 随手记
################

screen 配置
------------------

::

    # The resultsing command give you a status line that
    # # # # looks like this:
    # # # # | Screens: 0* bash <blanks zapped> 5:30PM Fri, Jun/25/2004 |
    # # # # (The pipes indicate the edges of the xterm/console).
    # # #
    # # # # Green text, time, and date; windows in blue:
    # #
    hardstatus alwayslastline
    # # Red Hat's normal status line
    # # #hardstatus string "[screen %n%?: %t%?] %h"
    hardstatus alwayslastline "%{=b}%{G} Screen(s): %{b}%w %=%{kG}%C%A %D, %M/%d/%Y "
    defscrollback 10000
    vbell off
    #把ctrl +a 替换成ctrl+z
    escape ^Zz
    hardstatus alwayslastline
    defscrollback 10000



    screen 配置

    ctr+A shift: 输入 fit

    消除 线 适应屏幕

    ctr+A然后 shift+: 输入 fit


reStructureText
---------------

超好用的一个插件能快速输入

`reStructureText`_ 

.. _reStructureText: http://knowledge.zhaoweiguo.com/7assists/emacs/sublimes/plugins/plugin_rst.html



`reStructureText <http://knowledge.zhaoweiguo.com/7assists/emacs/sublimes/plugins/plugin_rst.html>`_ 


.. code-block:: go

    type OrderParams struct {
        app.BaseForm

        OrderId      int64  `form:"order_id"`
        UserId       int    `form:"user_id"`
        Username     string `form:"username"`
        OrderStatus  *int   `form:"order_status"`
        CreatedStart string `form:"created_start"`
        CreatedEnd   string `form:"created_end"`
    }

