package service

import (
	"github.com/hwholiday/gid/v2/configs"
	"github.com/hwholiday/gid/v2/library/log"
	"github.com/hwholiday/gid/v2/repository"
	"go.uber.org/zap"
)

type Service struct {
	c     *configs.Config
	r     *repository.Repository
	alloc *Alloc
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
	return s
}

func (s *Service) Close() {
	s.r.Close()
}
