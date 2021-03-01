package client

import (
	"context"
	gidSrv "gid/api"
	"testing"
	"time"
)

func TestInitGrpc(t *testing.T) {
	cli, err := InitGrpc([]string{"127.0.0.1:2379"}, 15)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		c, _ := cli.GetGidGrpcClient()
		_, err := c.GetId(context.TODO(), &gidSrv.ReqId{
			BizTag: "111",
		})
		if err != nil {
			continue
		}
	}
}
