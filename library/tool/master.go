package tool

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/hwholiday/gid/v2/library/log"
	"go.uber.org/zap"
	"time"
)

type master struct {
	cli          *clientv3.Client
	ip           string
	key          string
	ttl          int64
	isMasterNode bool
	revision     int64
	id           clientv3.LeaseID
	isClose      bool
}

var MasterNode *master

func InitMasterNode(etcdAddr []string, ip string, ttl int64) error {
	MasterNode = &master{
		key: "/gid/master",
		ttl: ttl,
		ip:  ip,
	}
	var err error
	MasterNode.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.GetLogger().Error("[ApplyMasterNode] New", zap.Any("ip", ip), zap.Any("etcdAddr", etcdAddr), zap.Error(err))
		return err
	}
	go MasterNode.cornTTl()
	return nil
}

func (c *master) cornTTl() {
	if c == nil {
		panic("InitMasterNode is nil")
	}
	if err := c.applyMasterNode(); err != nil {
		panic(err)
	}
	c.watch()
	ticker := time.NewTicker(time.Duration(c.ttl) * time.Second)
	go func() {
		for {
			select {
			case _ = <-ticker.C:
				_ = c.applyMasterNode()
			}
		}
	}()
}

func (c *master) watch() {
	go func() {
		watcher := clientv3.NewWatcher(c.cli)
		watchChan := watcher.Watch(context.Background(), c.key, clientv3.WithRev(c.revision+1))
		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				switch event.Type {
				case mvccpb.DELETE:
					if !c.isClose {
						go c.applyMasterNode()
					}
				}
			}

		}
	}()
}

func (c *master) applyMasterNode() error {
	if c == nil {
		panic("InitMasterNode is nil")
	}
	lease := clientv3.NewLease(c.cli)
	if !c.isMasterNode {
		txn := clientv3.NewKV(c.cli).Txn(context.TODO())
		grantRes, err := lease.Grant(context.TODO(), c.ttl+1)
		if err != nil {
			log.GetLogger().Error("[ApplyMasterNode] New", zap.Any("ip", c.ip), zap.Error(err))
			c.isMasterNode = false
			return err
		}
		c.id = grantRes.ID
		txn.If(clientv3.Compare(clientv3.CreateRevision(c.key), "=", 0)).
			Then(clientv3.OpPut(c.key, c.ip, clientv3.WithLease(grantRes.ID))).
			Else()
		txnResp, err := txn.Commit()
		if err != nil {
			log.GetLogger().Error("[ApplyMasterNode] New", zap.Any("ip", c.ip), zap.Error(err))
			c.isMasterNode = false
			return err
		}
		if txnResp.Succeeded {
			c.isMasterNode = true
		} else {
			c.isMasterNode = false
		}
		if c.revision != txnResp.Header.Revision {
			c.revision = txnResp.Header.Revision
		}
	}
	_, err := lease.KeepAliveOnce(context.TODO(), c.id)
	if err != nil {
		c.isMasterNode = false
		log.GetLogger().Error("[ApplyMasterNode] New", zap.Any("ip", c.ip), zap.Error(err))
		return err
	}
	return nil
}

func (c *master) CloseApplyMasterNode() {
	if c != nil {
		c.isClose = true
		if _, err := c.cli.Delete(context.Background(), c.key); err != nil {
			log.GetLogger().Error("[CloseApplyMasterNode] Delete", zap.Any("ip", c.ip), zap.Error(err))
		}
	}
}
