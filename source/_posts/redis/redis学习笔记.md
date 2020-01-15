#### redis 特性
* 速度快
* 持久化
* 多种数据结构
* 支持多种编程语言
* 功能丰富
* 简单
* 主从复制
* 高可用,分布式

#### 通用命令
```bash
keys * #列出所有的key ,不建议线上使用
dbsize #统计所有的key数量
exists #检查某个key 是否存在
del # 删除某个key
expire key seconds #key 在seconds秒后过期
ttl key #查询key剩余的过期时间
persist key # 去掉key的过期时间
type key #返回key的类型

```
#### 单线程
1. 一次只运行一条命令
2. 拒绝长（慢）的命令


#### 字符串 操作
1.get
2.set
3.del
4.incr
5.decr
6.incrby
7.decrby