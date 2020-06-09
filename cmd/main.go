package main

import (
	"flag"
	"gid/configs"
	"gid/library/log"
)

func main() {
	flag.Parse()
	if err := configs.Init(); err != nil {
		panic(err)
	}
	log.NewLogger(configs.Conf.Log)
}
