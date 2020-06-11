package grpc

import (
	"fmt"
	gid "gid/api/grpc"
	"gid/configs"
	"gid/library/log"
	"gid/library/net/ip"
	"gid/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	srv *service.Service
}

func Init(c *configs.Config, s *service.Service) {
	addr := fmt.Sprintf("%s:%d", ip.InternalIP(), c.Grpc.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	g := grpc.NewServer()
	gid.RegisterGidServer(g, &Server{srv: s})
	log.GetLogger().Info("gid grpc server start", zap.Any("addr", addr))
	go g.Serve(lis)
}
