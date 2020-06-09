package repository

import (
	"gid/configs"
	"gid/library/database/mysql"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	c  *configs.Config
	db *gorm.DB
}

func NewRepository(c *configs.Config) (r *Repository) {
	r = &Repository{
		c:  c,
		db: mysql.NewMysql(c.Mysql),
	}
	return r
}

func (r *Repository)Close()  {
	_ = r.db.Close()
}


