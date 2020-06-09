package repository

import (
	"gid/entity"
	"gid/library/log"
	"gid/library/tool"
	"go.uber.org/zap"
)

func (r *Repository) SegmentsIdNext(tag string) (id *entity.Segments, err error) {
	var (
		tx = r.db.Begin()
	)
	id = &entity.Segments{}
	if err = tx.Exec("update segments set max_id=max_id+step,update_time = ? where biz_tag = ?", tool.GetTimeUnix(), tag).Error; err != nil {
		log.GetLogger().Error("[Repository] SegmentsIdNext Update", zap.String("tag", tag), zap.Error(err))
		tx.Rollback()
		return
	}
	if err = tx.Where("biz_tag = ?", tag).First(&id).Error; err != nil {
		log.GetLogger().Error("[Repository] SegmentsIdNext First", zap.String("tag", tag), zap.Error(err))
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}
