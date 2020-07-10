package http

import (
	"encoding/json"
	"gid/entity"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetId(g *gin.Context) {
	var (
		tag string
		id  int64
		err error
	)
	tag = g.Param("tag")
	if tag == "" {
		g.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "parameter error",
		})
		return
	}
	if id, err = srv.GetId(tag); err != nil {
		g.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": id,
	})
}

func GetRandId(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": srv.SnowFlakeGetId(),
	})
}

func CreateTag(g *gin.Context) {
	var data entity.Segments
	info, err := ioutil.ReadAll(g.Request.Body)
	if err != nil {
		g.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if err = json.Unmarshal(info, &data); err != nil {
		g.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	if data.Step == 0 || data.BizTag == "" {
		g.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "parameter error",
		})
		return
	}
	if err = srv.CreateTag(&data); err != nil {
		g.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
