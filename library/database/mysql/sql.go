package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/hwholiday/gid/v2/library/log"
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
func NewMysql(c *Config) (db *xorm.Engine) {
	var err error
	db, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@(%s)/%s?%s", c.User, c.Password, c.Addr, c.DbName, c.Parameters))
	if err != nil {
		log.GetLogger().Error("[NewMysql] Open", zap.Any("conf", c), zap.Error(err))
		panic(err)
	}
	db.SetMaxIdleConns(c.IdleConn)
	db.SetMaxOpenConns(c.MaxConn)
	db.SetConnMaxLifetime(time.Duration(c.IdleTimeout) * time.Millisecond)
	db.ShowSQL(c.Debug)
	db.ShowExecTime(c.Debug)
	if err = db.DB().Ping(); err != nil {
		log.GetLogger().Error("[NewMysql] Ping", zap.Any("conf", c), zap.Error(err))
		panic(err)
	}
	return
}
