# 简介
gid 是使用golang开发的生成分布式Id系统，基于数据库号段算法实现
### 性能
- id 基本上从内存生成，如果(step)步长设置的足够大,qps理论上只受物理限制
### 可用性
- id 分配依赖mysql ,当mysql不可用的,如果内存上还有的可以继续分配
### 特性
1. 全局唯一的int64型id  
2. 分配ID只访问内存  
3. 可无限横向扩展  
4. 依赖mysql恢复服务迅速
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
    remark      varchar(200) null,
    create_time bigint       null,
    update_time bigint       null,
    constraint segments_pk
        primary key (biz_tag)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

INSERT INTO segments(`biz_tag`, `max_id`, `step`, `remark`, `create_time`, `update_time`)
VALUES ('test', 0, 100000, 'test', 1591706686, 1591706686);
```
- 编译运行项目
```base
    git clone https://github.com/hwholiday/gid.git
    cd gid/cmd
    go build -o gidsrv
   ./gidsrv -conf ./gid.toml
```
 
#### 健康检查
- curl http://127.0.0.1:8080/ping

#### 获取ID
 - test6 获取该 tag 类型的 id
 - curl http://127.0.0.1:8080/id/test6
 
#### 创建 tag
- biz_tag tag 名称
- max_id  从这里开始派发ID
- step 步长
- remark 备注
- {"biz_tag":"test6","max_id":0,"step":10,"remark":"test6 tag"}
- curl -H "Content-Type:application/json" -X POST --data '{"biz_tag":"test6","max_id":0,"step":10,"remark":"test6 tag"}' http://127.0.0.1:8080/tag

### 文献
[美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)


