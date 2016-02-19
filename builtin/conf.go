package builtin

import (
	"os"
	"strings"

	"github.com/go-rut/config_reader"
)

const (
	confPath = "gorut.conf"
)

type GoRutConfig struct {
	Project string     `json:"project"`
	Envs    Envs       `json:"envs"`
	Deps    []RepoDeps `json:"deps"`
}

// RepoDeps describe the repository
// name: downloaning path name
// repo: the repository
// version: branch version, version is option
// vcs: go | git
type RepoDeps struct {
	Name    string      `json:"name"`
	Repo    string      `json:"repo"`
	Vcs     string      `json:"vcs"`
	Options DepsOptions `json:"options,omitempty"`
}

type DepsOptions struct {
	Remote string   `json:"remote,omitempty"`
	Branch string   `json:"branch,omitempty"`
	Flags  []string `json:"flags,omitempty"`
}

// Env for golang
// if not set gopath, if $GOROOT is not "", it will be $GOPATH, then gopath equals "."
type Envs struct {
	GoPath     string `json:"gopath,omitempty"`
	CGOEnabled bool   `json:"cgo_enabled,omitempty"`
}

func (p *GoRut) readConfig() {
	config := new(GoRutConfig)
	err := config_reader.NewConfigReader().JsonFileReader(confPath, config)
	if err != nil {
		panic(err)
	}

	if config.Envs.GoPath == "" {
		gopath := os.Getenv("GOPATH")
		if gopath != "" {
			goos := os.Getenv("GOOS")
			if goos == "windows" {
				gopath = strings.Split(gopath, ";")[0]
			} else {
				gopath = strings.Split(gopath, ":")[0]
			}
		}
		config.Envs.GoPath = gopath
	}
	if err = os.Setenv("GOPATH", config.Envs.GoPath); err != nil {
		panic(err)
	}

	CGOEnabled := "0"
	if config.Envs.CGOEnabled == true {
		CGOEnabled = "1"
	}
	if err = os.Setenv("CGO_ENABLED", CGOEnabled); err != nil {
		panic(err)
	}

	p.Config = config
	return
}
