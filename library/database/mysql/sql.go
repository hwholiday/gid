package mysql

import (
	"fmt"
	"gid/library/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

type Config struct {
	Addr       string
	User       string
	Password   string
	DbName     string
	Parameters string

	MaxConn      int
	IdleConn     int
	Debug        bool
	IdleTimeout  int
	QueryTimeout int //查询时间
	ExecTimeout  int //执行时间
}

//user:password@(addr)/dbname?charset=utf8&parseTime=True&loc=Local
func NewMysql(c *Config) (db *gorm.DB) {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?%s", c.User, c.Password, c.Addr, c.DbName, c.Parameters))
	if err != nil {
		log.GetLogger().Error("[NewMysql] Open", zap.Any("conf", c), zap.Error(err))
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.IdleConn)
	db.DB().SetMaxOpenConns(c.MaxConn)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) * time.Millisecond)
	db.LogMode(c.Debug)
	if err = db.DB().Ping(); err != nil {
		log.GetLogger().Error("[NewMysql] Ping", zap.Any("conf", c), zap.Error(err))
		panic(err)
	}
	return
}
