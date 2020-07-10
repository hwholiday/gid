package service

import "gid/library/tool"

func (s *Service) NewAllocSnowFlakeId() (a *tool.Worker, err error) {
	return tool.NewWorker(s.c.SnowFlakeId)
}

func (s *Service) SnowFlakeGetId() int64 {
	return s.snowFlake.GetId()
}
