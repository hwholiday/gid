package main

import (
	"flag"
	"gid/configs"
	"gid/library/log"
	"gid/library/tool"
	"gid/server/grpc"
	"gid/service"
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
