package http

import (
	"fmt"
	"gid/configs"
	"gid/library/log"
	"gid/library/net/ip"
	"gid/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	srv *service.Service
)

func Init(c *configs.Config, s *service.Service) {
	srv = s
	if !c.Development {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	setRouters(g)
	addr := fmt.Sprintf("%s:%d", ip.InternalIP(), c.Http.Port)
	h := &http.Server{
		Handler:        g,
		Addr:           addr,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.GetLogger().Info("gid http server start", zap.Any("addr", addr))
	go h.ListenAndServe()
}

func setRouters(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	r.GET("/ping", func(context *gin.Context) {
		_, _ = context.Writer.WriteString("pong")
	})
	r.POST("/tag", CreateTag)
	r.GET("/id/:tag", GetId)

}
