package client

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	gidSrv "github.com/hwholiday/gid/v2/api"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	etcd     *clientv3.Client
	node     string
	key      string
	change   bool
	ttl      int64
	revision int64
	conn     *grpc.ClientConn
}

func InitGrpc(etcdAddr []string, ttl int64) (*Client, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	cli := &Client{
		etcd:   c,
		change: false,
		key:    "/gid/master",
		ttl:    ttl,
	}
	cli.cornTTL()
	return cli, nil
}

func (c *Client) watch() {
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

func (c *Client) cornTTL() {
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

func (c *Client) getMasterNode() error {
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

func (c *Client) GetGidGrpcClient() (gidSrv.GidClient, error) {
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
