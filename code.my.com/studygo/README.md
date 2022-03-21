## 1. Mysql

#### 1.1 下载依赖库
`go get -u github.com/go-sql-driver/mysql`
`go get -u github.com/jmoiron/sqlx`

将第三方依赖库默认保存到 `$GOPATH/src/`

#### 1.2 实现代码

1. 先定义连接池对象

`X:\GoStudy\src\code.my.com\studygo\day08_db\01mysql\main.go`

#### 1.3 一些概念

###### 1.3.1 什么是事务？
`事务：`一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)，同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联合完成。A转账给B，这里面就需要执行两次update操作。

在MySQL中只有使用了Innodb数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行。


###### 1.3.2 事务的ACID
通常事务必须满足4个条件（ACID）：原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）


#### 1.4 注入

`xxx' or 1=1 #`
`xxx' union select * from user #`
`xxx' and (select count(*) from user) <10 #`

## 2. Redis

- Redis实战

#### 2.1 安装依赖库

`go get -u github.com/go-redis/redis`


## 3. 消息队列 NSQ

`go get -u github.com/nsqio/go-nsq`

## 4. 服务端Agent开发

> zookeeper、kafka部署文档：https://docs.qq.com/doc/DTmdldEJJVGtTRkFi

#### 4.1 Kafka

<div><img src='assets\kafka.png'></div>

1. Kafka集群的架构
    1. broker
    2. topic：
    3. partition：分区，把同一个topic分成不同的分区，分区的作用是做负载，提高kafka的吞吐量。（提高负载）
        1. leader：分区的主节点
        2. fpllower：分区的从节点
    4.  Consumer Group
2. 生产者往kafka发送数据的流程(6步)
<div><img src='assets\producer.png'></div>

3. Kafka选择分区的模式(3种)
    1. 指定往哪个分区写
    2. 指定key，kafka根据key做hash然后决定写哪个分区
    3. 轮询方式
4. 生产者往kafka发送数据的模式(3种)
    1. `0`: 把数据发送给leader就成功，效率最高，安全性最低
    2. `1`: 把数据发给leader，等待leader回ACK
    3. `all`: 把数据发给leader，确保follower从leader拉取数据回复ack给leader，leader再回复ACK，安全性最高
5. 分区存储文件的原理
6. 为什么kafka快
7. 消费者组消费数据的原理

#### 4.2 项目架构设计

下载地址：http://kafka.apache.org/downloads

###### 4.2.1 LogAgent的工作流程

1. 读日志 —— `tailf`第三方库
```cmd
set GO111MODULE=on
SET GOPROXY=https://goproxy.cn
go mod init example.com/m
go get github.com/hpcloud/tail
```

2. 往kafka写日志 —— `samara` 第三方库

- win 平台请用v1.19版本的samara

```cmd
go get github.com/Shopify/sarama
```

3. 

## Gin

`go get -u github.com/gin-gonic/gin`
