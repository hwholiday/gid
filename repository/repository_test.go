package repository

import (
	"flag"
	"gid/configs"
	"gid/library/log"
	"os"
	"testing"
)

var (
	r *Repository
)

func TestMain(m *testing.M) {
	_ = flag.Set("conf", "./../cmd/gid.toml")
	flag.Parse()
	if err := configs.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(configs.Conf.Log)
	r = NewRepository(configs.Conf)
	os.Exit(m.Run())
}
