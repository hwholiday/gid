package service

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_GetId(t *testing.T) {
	convey.Convey("TestService_GetId", t, func(c convey.C) {
		res, err := GetService().GetId("test")
		convey.So(err, convey.ShouldBeNil)
		t.Logf("res %v", res)
	})
}

func BenchmarkService_GetId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetService().GetId("test")
		if err != nil {
			b.Error(err)
		}
	}
}
