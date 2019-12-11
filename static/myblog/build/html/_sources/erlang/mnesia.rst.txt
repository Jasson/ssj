Mnesia遇到的问题
######################

从 Ejabberd开源项目中了解到了 Mnesia的用法，之后好多长一段时间都用 Mnesia 设计分布式缓存。


Mnesia 数据丢失（并未真正丢失，只是索引丢失）
------------------------------------------------

用Mnesia构建的 ws_session,数据读写都有用了dirty操作。

系统刚上线没有出现问题，过了一段时间发现用户登录出现异常（总结发现都是用户切换网络时候出现）。进一步分析发现数据不同步了，例如根据  mnesia:dirty_index_read(ws_session, 2, id) 查数据总是查出空，偶尔在节点上查到数据（后来发现是根据 Key），起初误以为集群节点数据不同步。多次试验后发现 index查不到数据，key 能查到数据，得出结论是索引丢失了，百思不得其解，进一步了解发现是mnesia:dirty_delete_object删除时候只删除了索引，没有删除数据。

原来delete_object时候原来record 已被修改过了，才导致此次删除数据时只删除了索引，现在看上去比较合理，毕竟只丢失了索引。

但是更让人意外的是即使数据重写也不会再生成索引了，必须先根据 Key 把数据删除。

解决方案，删除数据的时候避免采用 delete_object,改为直接根据 Key 删除。查询数据时，先根据索引查询数据，若发现有索引丢失，立马根据 Key 把数据删除，重新生成一份数据，这样才能保证生成索引。


手册中有明确的解释::

    delete_object(Tab, Record, LockKind) -> transaction abort | ok 

    If a table is of type bag, it can sometimes be needed to delete only some of the records with a certain key. This can be done with the function delete_object/3. A complete record must be supplied to this function.

    The semantics of this function is context-sensitive. For details, see mnesia:activity/4. In transaction-context, it acquires a lock of type LockKind on the record. Currently, the lock types write and sticky_write are supported.




