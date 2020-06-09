package repository

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRepository_IdNext(t *testing.T) {
	convey.Convey("TestRepository_IdNext", t, func(c convey.C) {
		res, err := r.SegmentsIdNext("test")
		c.So(err, convey.ShouldBeNil)
		t.Log(res)
	})
}
