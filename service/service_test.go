package service

import (
	"flag"
	"gid/configs"
	"gid/library/log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = flag.Set("conf", "./../cmd/gid.toml")
	flag.Parse()
	if err := configs.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(configs.Conf.Log)
	NewService(&configs.Conf)
	os.Exit(m.Run())
}
