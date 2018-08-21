package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zerocruft/pharos/cmd/pharos/config"
	"go.uber.org/zap"
)

var (
	flagStart      bool
	flagBuild      bool
	flagConfigFile string
)

func main() {
	flag.BoolVar(&flagStart, "start", false, "")
	flag.BoolVar(&flagBuild, "build", false, "")
	flag.StringVar(&flagConfigFile, "config", "config.toml", "")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger.Info("testing")
	cfg := config.New(flagConfigFile, logger)
	if cfg == nil {
		logger.Error("config failed to build. exiting")
		os.Exit(1)
	}

	if flagBuild {
		build(*cfg, logger)
		return
	}
	if flagStart {
		go start()
	}
}
