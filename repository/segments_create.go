package repository

import (
	"gid/entity"
	"gid/library/log"
	"gid/library/tool"
	"go.uber.org/zap"
)

func (r *Repository) SegmentsCreate(s *entity.Segments) (data *entity.Segments, err error) {
	var has bool
	data = new(entity.Segments)
	if has, err = r.db.Where("biz_tag = ?", s.BizTag).Get(data); err != nil {
		log.GetLogger().Error("[SegmentsCreate] Exist", zap.Any("req", s), zap.Error(err))
		return
	}
	if has {
		return
	}
	s.CreateTime = tool.GetTimeUnix()
	s.UpdateTime = tool.GetTimeUnix()
	if _, err = r.db.Insert(s); err != nil {
		log.GetLogger().Error("[SegmentsCreate] Create", zap.Any("req", s), zap.Error(err))
		return
	}
	data = s
	return
}
