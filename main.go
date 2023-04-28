package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MysGate/demo_backend/conf"
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

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
	util.Logger().Info("system shutdown")
}
