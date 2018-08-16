package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/zerocruft/pharos"
	"github.com/zerocruft/pharos/cmd/pharos/config"
)

func build(cfg config.Config) {
	preManifestBS, err := ioutil.ReadFile(cfg.Source + string(os.PathSeparator) + "manifest.toml")
	if err != nil {
		fmt.Println("manifest.toml is required in source directory")
		return
	}
	manifestBS := preprocess(preManifestBS)
	var manifest pharos.Manifest
	_, err = toml.Decode(string(manifestBS), &manifest)
	fmt.Println(manifest)

}

func preprocess(bs []byte) []byte {
	var rb []byte
	var sort int
	var mlsTracker bool

	for i, b := range bs {
		if i < 4 {
			rb = append(rb, b)
			continue
		}
		if b == '"' {
			if bs[i-1] == '"' && bs[i-2] == '"' {
				mlsTracker = !mlsTracker
			}
			rb = append(rb, b)
			continue
		}
		if mlsTracker {
			rb = append(rb, b)
			continue
		}

		// Logic for building sort ids
		if bs[i-1] == '\n' && bs[i-2] == ']' && bs[i-3] == ']' {
			rb = append(rb, []byte("sort = "+strconv.Itoa(sort)+"\n")...)
			rb = append(rb, b)
			sort++
			continue
		}
		rb = append(rb, b)
	}
	return rb
}
