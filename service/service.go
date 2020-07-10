package service

import (
	"gid/configs"
	"gid/library/log"
	"gid/library/tool"
	"gid/repository"
	"go.uber.org/zap"
)

type Service struct {
	c         *configs.Config
	r         *repository.Repository
	alloc     *Alloc
	snowFlake *tool.Worker
}

func NewService(c *configs.Config) (s *Service) {
	var err error
	s = &Service{
		c: c,
		r: repository.NewRepository(c),
	}
	if s.alloc, err = s.NewAllocId(); err != nil {
		log.GetLogger().Error("[NewService] NewAllocId ", zap.Error(err))
		panic(err)
	}
	if s.snowFlake, err = s.NewAllocSnowFlakeId(); err != nil {
		log.GetLogger().Error("[NewService] NewAllocSnowFlakeId ", zap.Error(err))
		panic(err)
	}
	return s
}

func (s *Service) Close() {
	s.r.Close()
}
