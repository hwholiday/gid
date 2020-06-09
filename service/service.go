package service

import (
	"fmt"
	"gid/configs"
	"gid/library/log"
	"gid/repository"
	"go.uber.org/zap"
)

var s *Service

type Service struct {
	c     *configs.Config
	r     *repository.Repository
	alloc *Alloc
}

func NewService(c *configs.Config) {
	var err error
	s = &Service{
		c: c,
		r: repository.NewRepository(c),
	}
	if s.alloc, err = s.NewAllocId(); err != nil {
		log.GetLogger().Error("[NewService] NewAllocId ", zap.Error(err))
		panic(err)
	}
	fmt.Println(s)
}

func GetService() *Service {
	return s
}

func (s *Service) Close() {
	s.r.Close()
}
