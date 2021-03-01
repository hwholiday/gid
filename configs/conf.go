package configs

import (
	"flag"
	"gid/library/database/mysql"
	"gid/library/log"
	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf     = new(Config)
)

type Config struct {
	Development bool
	SnowFlakeId int64
	Etcd        []string
	Log         *log.Options
	Mysql       *mysql.Config
	Server      *Srv
}

type Srv struct {
	Addr string
}

func init() {
	flag.StringVar(&confPath, "conf", "./gid.toml", "default config path")
}

func Init() error {
	return local()
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	Conf.Mysql.Debug = Conf.Development
	return
}
