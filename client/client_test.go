package client

import (
	"context"
	"fmt"
	gidSrv "github.com/hwholiday/gid/v2/api"
	"testing"
)

func TestInitGrpc(t *testing.T) {
	cli, err := InitGrpc([]string{"127.0.0.1:2379"}, 15)
	c, _ := cli.GetGidGrpcClient()
	res, err := c.GetId(context.TODO(), &gidSrv.ReqId{
		BizTag: "111",
	})
	fmt.Println(res, err)
}
