package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zerocruft/pharos/cmd/pharos/config"
)

var (
	flagStart bool
	flagBuild bool
	flagConf  string
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(wd)
	flag.BoolVar(&flagStart, "start", false, "")
	flag.BoolVar(&flagBuild, "build", false, "")
	flag.StringVar(&flagConf, "conf", "config.toml", "")
	flag.Parse()

	cfg := config.New(flagConf)
	if cfg == nil {
		fmt.Println("No Config. Exiting")
		os.Exit(1)
	}

	if flagBuild {
		build(*cfg)
		// build(cfg)
		return
	}
	if flagStart {
		go start()
	}

	fmt.Println("hw")
}
