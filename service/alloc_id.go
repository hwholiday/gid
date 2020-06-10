package service

import (
	"errors"
	"gid/entity"
)

func (s *Service) GetId(tag string) (id int64, err error) {
	s.alloc.Mu.Lock()
	defer s.alloc.Mu.Unlock()
	val, ok := s.alloc.BizTagMap[tag]
	if !ok {
		return 0, errors.New("not find tag")
	}
	return val.GetId(s)
}

func (s *Service) CreateTag(e *entity.Segments) (err error) {
	if err = s.r.SegmentsCreate(e); err != nil {
		return
	}
	b := &BizAlloc{
		BazTag:  e.BizTag,
		GetDb:   false,
		IdArray: make([]*IdArray, 0),
	}
	b.IdArray = append(b.IdArray, &IdArray{
		Cur:   1,
		Start: 0,
		End:   e.Step,
	})
	s.alloc.BizTagMap[e.BizTag] = b
	return
}
