package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zerocruft/pharos/cmd/pharos/config"
	git "gopkg.in/src-d/go-git.v4"
)

func saveForLater(cfg config.Config) {
	// preManifestBS, err := ioutil.ReadFile(cfg.Source + string(os.PathSeparator) + "manifest.toml")
	// if err != nil {
	// 	fmt.Println("manifest.toml is required in source directory")
	// 	return
	// }
	// manifestBS := preprocess(preManifestBS)
	// var manifest pharos.Manifest
	// _, err = toml.Decode(string(manifestBS), &manifest)
	// if err != nil {
	// 	fmt.Println(err)
	// 	//DO something here
	// 	return
	// }
	// fmt.Println(manifest)
}

func build(cfg config.Config) {
	// TODO make sure build directory is ready

	err := os.Mkdir(cfg.BuildDir, 0755)
	if err != nil {
		fmt.Println(err)
	}
	if err := os.RemoveAll(cfg.BuildDir + string(os.PathSeparator) + "temp"); err != nil {
		fmt.Println(err)
	}
	if err := os.Mkdir(cfg.BuildDir+string(os.PathSeparator)+"temp", 0755); err != nil {
		fmt.Println(err)
	}
	for _, source := range cfg.Sources {
		r, err := git.PlainClone(cfg.BuildDir+string(os.PathSeparator)+"temp"+string(os.PathSeparator)+source.Name, false, &git.CloneOptions{
			URL:               source.GitURL,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(r)
	}
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
