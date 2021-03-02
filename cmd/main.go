package main

import (
	"flag"
	"github.com/hwholiday/gid/v2/configs"
	"github.com/hwholiday/gid/v2/library/log"
	"github.com/hwholiday/gid/v2/library/tool"
	"github.com/hwholiday/gid/v2/server/grpc"
	"github.com/hwholiday/gid/v2/service"
)

func main() {
	flag.Parse()
	if err := configs.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(configs.Conf.Log)
	s := service.NewService(configs.Conf)
	grpc.Init(configs.Conf, s)
	if err := tool.InitMasterNode(configs.Conf.Etcd, configs.Conf.Server.Addr, 30); err != nil {
		panic(err)
	}
	tool.QuitSignal(func() {
		s.Close()
		tool.MasterNode.CloseApplyMasterNode()
		log.GetLogger().Info("gid exit success")
	})
}
