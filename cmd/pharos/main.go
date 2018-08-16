package main

import (
	"flag"
	"fmt"

	"github.com/zerocruft/pharos/cmd/pharos/config"
)

var (
	flagStart bool
	flagBuild bool
	flagConf  string
)

func main() {
	flag.BoolVar(&flagStart, "start", false, "")
	flag.BoolVar(&flagBuild, "build", false, "")
	flag.StringVar(&flagConf, "conf", "config.toml", "")
	flag.Parse()

	cfg := config.Config{
		Source: "/home/jab/code/dev/pharos101",
	}
	if flagBuild {
		build(cfg)
		return
	}
	if flagStart {
		go start()
	}

	fmt.Println("hw")
}
