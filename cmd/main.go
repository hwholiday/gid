package main

import (
	"flag"
	"gid/configs"
	"gid/library/log"
	"gid/library/tool"
	"gid/server/grpc"
	"gid/server/http"
	"gid/service"
)

func main() {
	flag.Parse()
	if err := configs.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(configs.Conf.Log)
	s := service.NewService(configs.Conf)
	http.Init(configs.Conf, s)
	grpc.Init(configs.Conf, s)
	tool.QuitSignal(func() {
		s.Close()
		log.GetLogger().Info("gid exit success")
	})
}
