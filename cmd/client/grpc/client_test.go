package grpc

import (
	"context"
	"flag"
	"fmt"
	gid "gid/api/grpc"
	"gid/configs"
	"gid/library/net/ip"
	"github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"os"
	"testing"
)

var client gid.GidClient

func TestMain(m *testing.M) {
	_ = flag.Set("conf", "./../../gid.toml")
	flag.Parse()
	if err := configs.Init(); err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip.InternalIP(), configs.Conf.Grpc.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = gid.NewGidClient(conn)
	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	convey.Convey("TestPing", t, func(c convey.C) {
		res, err := client.Ping(context.Background(), &gid.ReqPing{})
		c.So(err, convey.ShouldBeNil)
		c.So(res, convey.ShouldNotBeNil)
		t.Log(res)
	})
}

func TestCreateTag(t *testing.T) {
	convey.Convey("TestCreateTag", t, func(c convey.C) {
		res, err := client.CreateTag(context.Background(), &gid.ReqTagCreate{
			Tag:    "test01",
			Step:   10000,
			Remark: "test01 tag",
		})
		c.So(err, convey.ShouldBeNil)
		c.So(res, convey.ShouldNotBeNil)
		t.Log(res)
	})
}

func TestGetId(t *testing.T) {
	convey.Convey("TestGetId", t, func(c convey.C) {
		res, err := client.GetId(context.Background(), &gid.ReqId{
			Tag: "test01",
		})
		c.So(err, convey.ShouldBeNil)
		c.So(res, convey.ShouldNotBeNil)
		t.Log(res)
	})
}

//go test -bench=. -run=none
func BenchmarkGetId(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := client.GetId(context.Background(), &gid.ReqId{
			Tag: "test01",
		})
		if err != nil {
			b.Error(err)
		}
		if res == nil {
			b.Error("res is nil")
		}
		if res.Status.Code != 200 {
			b.Error("req err")
		}
	}
}
