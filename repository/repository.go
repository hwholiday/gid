package repository

import (
	"github.com/go-xorm/xorm"
	"github.com/hwholiday/gid/v2/configs"
	"github.com/hwholiday/gid/v2/library/database/mysql"
)

type Repository struct {
	c  *configs.Config
	db *xorm.Engine
}

func NewRepository(c *configs.Config) (r *Repository) {
	r = &Repository{
		c:  c,
		db: mysql.NewMysql(c.Mysql),
	}
	return r
}

func (r *Repository) Close() {
	_ = r.db.Close()
}
