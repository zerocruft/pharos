package main

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"

	"github.com/zerocruft/pharos"

	"github.com/zerocruft/pharos/cmd/pharos/config"
	"go.uber.org/zap"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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

func build(cfg config.Config, logger *zap.Logger) {
	logger.Info("(re)building sources")

	logger.Info("creating workspace directory if not exists", zap.String("location", cfg.Workspace))
	err := os.Mkdir(cfg.Workspace, 0755)
	if err != nil {
		logger.Debug("failed to create directory, assuming its because it already exists", zap.Error(err))
	}

	logger.Info("cleaning build directory for build process")
	buildDir := cfg.Workspace + cfg.PathSep + "build"
	if err := os.RemoveAll(buildDir); err != nil {
		logger.Error("failure to purge build directory", zap.Error(err), zap.String("location", buildDir))
		return //TODO Rollback or clean ?
	}
	if err := os.Mkdir(buildDir, 0755); err != nil {
		logger.Error("failure to create build directory", zap.Error(err), zap.String("location", buildDir))
	}

	for _, source := range cfg.Sources {
		if err := cloneGitSources(source, buildDir, logger); err != nil {
			logger.Error("failed to clone git sources", zap.Error(err))
		}
		fb, err := ioutil.ReadFile(buildDir + cfg.PathSep + source.Name + cfg.PathSep + "manifest.toml")
		if err != nil {
			logger.Error("source is missing manifest.toml", zap.Error(err), zap.String("source", source.Name))
			continue
		}
		pfb := preprocess(fb)
		var manifest pharos.Manifest
		_, err = toml.Decode(string(pfb), &manifest)
		if err != nil {
			logger.Error("unable to parse manifest.toml", zap.Error(err))
			continue
		}
	}

}

func cloneGitSources(source config.Source, buildDir string, logger *zap.Logger) error {
	if source.RemoteHead == "" {
		source.RemoteHead = "master"
	}
	r, err := git.PlainClone(buildDir+string(os.PathSeparator)+source.Name, false, &git.CloneOptions{
		URL:               source.GitURL,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		ReferenceName:     plumbing.ReferenceName("refs/heads/" + source.RemoteHead),
	})
	if err != nil {
		logger.Error("failed to clone source repository", zap.Error(err), zap.Any("source", source))
		return err
	}
	logger.Info("cloned source", zap.Any("source", source), zap.Any("repo", r))
	return nil
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
