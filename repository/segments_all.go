package repository

import (
	"gid/entity"
	"gid/library/log"
	"go.uber.org/zap"
)

func (r *Repository) SegmentsGetAll() (res []entity.Segments, err error) {
	if err = r.db.Find(&res); err != nil {
		log.GetLogger().Error("[Repository] SegmentsGetAll Find", zap.Error(err))
	}
	return
}
