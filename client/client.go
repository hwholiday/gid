package client

import (
	"context"
	gidSrv "gid/api"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"time"
)

type client struct {
	etcd     *clientv3.Client
	node     string
	key      string
	change   bool
	ttl      int64
	revision int64
	conn     *grpc.ClientConn
}

func InitGrpc(etcdAddr []string, ttl int64) (*client, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	cli := &client{
		etcd:   c,
		change: false,
		key:    "/gid/master",
		ttl:    ttl,
	}
	cli.cornTTL()
	return cli, nil
}

func (c *client) watch() {
	watcher := clientv3.NewWatcher(c.etcd)
	watchChan := watcher.Watch(context.Background(), c.key, clientv3.WithRev(c.revision+1))
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.DELETE:
				go c.getMasterNode()
			}
		}
	}
}

func (c *client) cornTTL() {
	if err := c.getMasterNode(); err != nil {
		panic(err)
	}
	go c.watch()
	ticker := time.NewTicker(time.Duration(c.ttl) * time.Second)
	go func() {
		for {
			select {
			case _ = <-ticker.C:
				_ = c.getMasterNode()
			}
		}
	}()
}

func (c *client) getMasterNode() error {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	res, err := c.etcd.Get(ctx, c.key)
	if err != nil {
		return err
	}
	for _, v := range res.Kvs {
		val := v
		if string(val.Key) == c.key {
			newNode := string(val.Value)
			if c.node != newNode {
				c.change = true
			}
			c.node = string(val.Value)
		}
	}
	if c.revision != res.Header.Revision {
		c.revision = res.Header.Revision
	}
	return nil
}

func (c *client) GetGidGrpcClient() (gidSrv.GidClient, error) {
	var err error
	if c.change {
		c.conn, err = grpc.Dial(c.node, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		c.change = false
	}
	return gidSrv.NewGidClient(c.conn), nil
}
