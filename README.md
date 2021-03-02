# 简介

gid 是使用golang开发的生成分布式Id系统，基于数据库号段算法实现

# gid V2版本实现了高可用，主从架构，简化了调用逻辑

###### GRPC 对外服务

### 性能

- id 从内存生成，如果(step)步长设置的足够大,qps可达到千万+

### 可用性

- id 分配依赖mysql ,当mysql不可用的,如果内存上还有的可以继续分配

### 特性

1. 全局唯一的int64型id
2. 分配ID只访问内存
3. 可无限横向扩展
4. 依赖mysql恢复服务迅速
4. 依赖etcd实现服务的高可用 ......

### 高可用

1. server 基于 ETCD Lease 实现了自动抢主，主节点挂了，从节点自动申请为主节点

### 简化调用逻辑

```base
  go get github.com/hwholiday/gid/v2@gidV2
  
  //go mod 内
  //github.com/hwholiday/gid/v2 v2.0.6
```
1. 封住了 client 实现了自动识别服务主节点
2. 只需要实现 client 并调用 GetId(GRPC方法)，无需其他接口，自动创建BizTag，并预加载

```base
package main

import (
	"context"
	"fmt"
	gidSrv "github.com/hwholiday/gid/v2/api"
)

func main() {
	cli, err := InitGrpc([]string{"127.0.0.1:2379"}, 15)
	c, _ := cli.GetGidGrpcClient()
	res, err := c.GetId(context.TODO(), &gidSrv.ReqId{
		BizTag: "111",
	})
	fmt.Println(res, err)
}
```
### 安装

- 初始化 mysql

```base
create database gid;
use gid;
create table segments
(
    biz_tag     varchar(128) not null,
    max_id      bigint       null,
    step        int          null,
    create_time bigint       null,
    update_time bigint       null,
    constraint segments_pk
        primary key (biz_tag)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

```

- 编译运行项目

```base
    go get github.com/hwholiday/gid/v2@gidV2  OR  git clone -b gidV2 https://github.com/hwholiday/gid.git 
    cd gid/cmd
    go build -o gid
   ./gid -conf ./gid.toml
```

#### 压测

```base
BenchmarkService_GetId-4         2046296               583 ns/op 
```

#### 重点SQL

```base
Begin
UPDATE table SET max_id=max_id+step WHERE biz_tag=xxx
SELECT tag, max_id, step FROM table WHERE biz_tag=xxx
Commit
```

### 文献

[美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)


