package http

import (
	"gid/entity"
	"github.com/gin-gonic/gin"
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

func CreateTag(g *gin.Context) {
	var (
		err error
	)
	var data entity.Segments
	if err = g.BindJSON(&data); err != nil {
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
