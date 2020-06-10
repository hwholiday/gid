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
	Log         *log.Options
	Mysql       *mysql.Config
	Http        *httpConf
}

type httpConf struct {
	Port int
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

func Init() error {
	return local()
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	Conf.Mysql.Debug = Conf.Development
	return
}
