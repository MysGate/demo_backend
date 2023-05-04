package main

import (
	"flag"
	"time"

	"github.com/MysGate/demo_backend/conf"
	"github.com/MysGate/demo_backend/service"
	"github.com/MysGate/demo_backend/util"
)

func fixTimeZone() {
	time.Local = time.FixedZone("CST", 3600*8)
}

func initConfig(yaml string) *conf.MysGateConfig {
	conf.ParseYaml(yaml)
	c := conf.GetConfig()
	return c
}

func initLogger(cfg *conf.MysGateConfig) {
	util.InitLog(cfg.Logger.Path, cfg.Logger.Level)
}

func main() {
	var yaml string
	flag.StringVar(&yaml, "c", "", "config yaml file")
	flag.Parse()

	fixTimeZone()

	c := initConfig(yaml)
	defer c.CloseClient()

	initLogger(c)

	s := service.NewHttpServer(c)
	s.RunHttpService()

	util.Logger().Info("system shutdown")
}
