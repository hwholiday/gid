package repository

import (
	"gid/entity"
	"gid/library/tool"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRepository_SegmentsCreate(t *testing.T) {
	convey.Convey("TestRepository_SegmentsCreate", t, func(c convey.C) {
		err := r.SegmentsCreate(&entity.Segments{
			BizTag:     "test2",
			MaxId:      1,
			Step:       10000,
			Remark:     "test2",
			CreateTime: tool.GetTimeUnix(),
			UpdateTime: tool.GetTimeUnix(),
		})
		c.So(err, convey.ShouldBeNil)
	})
}
