package repository

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRepository_SegmentsGetAll(t *testing.T) {
	convey.Convey("TestRepository_SegmentsGetAll", t, func(c convey.C) {
		res, err := r.SegmentsGetAll()
		c.So(err, convey.ShouldBeNil)
		t.Log(res)
	})
}
