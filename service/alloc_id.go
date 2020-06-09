package service

import "errors"

func (s *Service) GetId(tag string) (id int64, err error) {
	s.alloc.Mu.Lock()
	defer s.alloc.Mu.Unlock()
	val, ok := s.alloc.BizTagMap[tag]
	if !ok {
		return 0, errors.New("not find tag")
	}
	return val.GetId()
}
