package repository

import (
	"errors"
	"gid/entity"
	"gid/library/log"
	"go.uber.org/zap"
)

func (r *Repository) SegmentsCreate(s *entity.Segments) (err error) {
	var has bool
	if has, err = r.db.Where("biz_tag = ?", s.BizTag).Exist(&entity.Segments{}); err != nil {
		log.GetLogger().Error("[SegmentsCreate] Exist", zap.Any("req", s), zap.Error(err))
		return
	}
	if has {
		err = errors.New("tag already exists")
		return
	}
	if _, err = r.db.Insert(s); err != nil {
		log.GetLogger().Error("[SegmentsCreate] Create", zap.Any("req", s), zap.Error(err))
		return
	}
	return
}
