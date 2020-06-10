
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





