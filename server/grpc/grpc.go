package grpc

import (
	gidSrv "gid/api"
	"gid/configs"
	"gid/library/log"
	"gid/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	srv *service.Service
}

func Init(c *configs.Config, s *service.Service) {
	lis, err := net.Listen("tcp", c.Server.Addr)
	if err != nil {
		panic(err)
	}
	g := grpc.NewServer()
	gidSrv.RegisterGidServer(g, &Server{srv: s})
	log.GetLogger().Info("gid grpc server start", zap.Any("addr", c.Server.Addr))
	go g.Serve(lis)
}
