package grpc

import (
	gidSrv "github.com/hwholiday/gid/v2/api"
	"github.com/hwholiday/gid/v2/configs"
	"github.com/hwholiday/gid/v2/library/log"
	"github.com/hwholiday/gid/v2/service"
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
