---
categories:
- mq
tags:
- rabbitmq
---

<!--more-->


## 消息队列协议
传递、存储、分发  
### 面试题：为什么消息中间件不直接使用http协议。
1.http协议请求报文头和响应报文头是比较复杂的，包含了消息传递不需要的部分。  
2.http大部分是短连接，一个请求到响应很可能会中断，中断以后就不会持久化，消息中间件可能是一个长期的获取消息的过程，出现问题和故障要对数据或消息持久化。  
### 常见消息中间件协议
AMQP、MQTT、Kafka、OpenMessage
#### AMQP
Erlang语言开发，默认端口5672  
产品支持：rabbmitmq acitvemq  

* 分布式事务支持
* 消息的持久化支持
* 高性能和高可靠的消息处理优势
#### MQTT协议
MQTT协议是IBM开放的一个即时通讯协议，物联网系统架构中重要组成部分  
产品支持：rabbmitmq acitvemq  

特点  
1. 轻量  
2. 结构简单  
3. 传输快、不支持事务  
4. 没有持久化设计。
应用场景：   
1. 适用于计算能力有限   
2. 低带宽 
3. 网络不稳地的场景  

#### OpenMessage
由阿里、雅虎、滴滴等共同参与创立的分布式消息中间件（布道者保持中立态度。。。）
1. 结构简单  
2. 解析速度快  
3. 支持事务和持久化  

产品支持：rocketmq  

#### kafka协议
kafka协议是基于TCP/IP的二进制协议，消息内部是通过长度分割，由一些基本的数据类型组成  
特点是：  
1. 结构简单  
2. 解析速度快  
3. 无事务支持  
4. 有持久化设计  

## 消息持久化
将数据存入磁盘中

## 消息的分发策略
消息队列MQ是一种推送的过程。  

### 消息分发策略的机制和对比
||ActiveMQ|RabbitMQ|Kafka|RocketMQ|
||:---:|:---:|:---:|:---:|
|发布订阅|支持|支持|支持|支持|
|轮询分发|支持|支持|支持|/|
|公平分发|/|支持|支持|/|
|重发|支持|支持|/|支持|
|消息拉取|/|支持|支持|支持|

> 轮询分发，消费分发数量是均匀的。
> 公平分发会根据服务器的性能分发消息，会造成消息的倾斜。rabbitmq 使用公平分发需要使用手动确认的方式。








