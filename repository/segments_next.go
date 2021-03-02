package repository

import (
	"github.com/hwholiday/gid/v2/entity"
	"github.com/hwholiday/gid/v2/library/log"
	"github.com/hwholiday/gid/v2/library/tool"
	"go.uber.org/zap"
)

func (r *Repository) SegmentsIdNext(tag string) (id *entity.Segments, err error) {
	var (
		tx = r.db.Prepare()
	)
	id = &entity.Segments{}
	if _, err = tx.Exec("update segments set max_id=max_id+step,update_time = ? where biz_tag = ?", tool.GetTimeUnix(), tag); err != nil {
		log.GetLogger().Error("[Repository] SegmentsIdNext Update", zap.String("tag", tag), zap.Error(err))
		_ = tx.Rollback()
		return
	}
	if _, err = tx.Where("biz_tag = ?", tag).Get(id); err != nil {
		log.GetLogger().Error("[Repository] SegmentsIdNext Get", zap.String("tag", tag), zap.Error(err))
		_ = tx.Rollback()
		return
	}
	err = tx.Commit()
	return
}
