package service

import (
	"gid/entity"
)

func (s *Service) GetId(tag string) (id int64, err error) {
	s.alloc.Mu.Lock()
	defer s.alloc.Mu.Unlock()
	val, ok := s.alloc.BizTagMap[tag]
	if !ok {
		if err = s.CreateTag(&entity.Segments{
			BizTag: tag,
			MaxId:  1,
			Step:   1000,
		}); err != nil {
			return 0, err
		}
		val, _ = s.alloc.BizTagMap[tag]
	}
	return val.GetId(s)
}

func (s *Service) CreateTag(e *entity.Segments) error {
	data, err := s.r.SegmentsCreate(e)
	if err != nil {
		return err
	}
	b := &BizAlloc{
		BazTag:  e.BizTag,
		GetDb:   false,
		IdArray: make([]*IdArray, 0),
	}
	b.IdArray = append(b.IdArray, &IdArray{Start: data.MaxId, End: data.MaxId + data.Step})
	s.alloc.BizTagMap[e.BizTag] = b
	return nil
}
