package repository

import (
	"gid/entity"
	"gid/library/log"
	"gid/library/tool"
	"go.uber.org/zap"
)

//只取6个小时有变化的Tag
func (r *Repository) SegmentsGetAll() (res []entity.Segments, err error) {
	if err = r.db.Where("update_time >= ?", tool.GetTimeUnix()-21600).Find(&res); err != nil {
		log.GetLogger().Error("[Repository] SegmentsGetAll Find", zap.Error(err))
	}
	return
}
