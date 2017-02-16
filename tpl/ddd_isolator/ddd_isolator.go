// MIT License

// Copyright (c) 2016 rutcode-go

package ddd_isolator

import (
	"bytes"
	"fmt"
	"html/template"
	"path"

	"github.com/go-rut/files"
)

const (
	_pathSrc = "/src/"
)

const (
	// _DDDPathConf       = "conf"
	// _DDDPathRepoLogic  = "repository/repo"
	_PathConf     = "conf"
	_PathServices = "services"
	_PathDocs     = "docs"
	_PathCommon   = "common"
	_PathLogics   = "logics"
	_PathDomain   = "domain"
	_PathRepoes   = "repo"
)

type iosolatorTemplate struct {
	GoPath  string
	AppPath string

	PathMap  map[string]string
	FilesMap map[string]string
	Imports  map[string]string
}

var defaultTemplate *iosolatorTemplate

func NewIosolatorTemplate(goPath, appPath string) *iosolatorTemplate {
	if defaultTemplate == nil {
		defaultTemplate = &iosolatorTemplate{
			GoPath:  goPath,
			AppPath: appPath,

			PathMap: map[string]string{
				"PathCommon":   _PathCommon,
				"PathDocs":     _PathDocs,
				"PathLogics":   _PathLogics,
				"PathDomain":   _PathDomain,
				"PathRepoes":   _PathRepoes,
				"PathServices": _PathServices,
				"PathConf":     _PathConf,
			},
			Imports: map[string]string{
				"PathCommon":   pathJoin(appPath, _PathCommon),
				"PathLogics":   pathJoin(appPath, _PathLogics),
				"PathDomain":   pathJoin(appPath, _PathDomain),
				"PathRepoes":   pathJoin(appPath, _PathRepoes),
				"PathServices": pathJoin(appPath, _PathServices),
			},
			FilesMap: map[string]string{
				"main.go": _DDDMainGo,

				pathJoin(_PathCommon, "common.go"):          commonGo,
				pathJoin(_PathCommon, "init.go"):            initGo,
				pathJoin(_PathLogics, "logics.go"):          logicsGo,
				pathJoin(_PathRepoes, "user_repository.go"): userRepoGo,
				pathJoin(_PathRepoes, "redis_repo.go"):      userRepoRedisGo,
				pathJoin(_PathRepoes, "sql_repo.go"):        userRepoSqlGo,
				pathJoin(_PathDomain, "user.go"):            userDomainGo,
				pathJoin(_PathServices, "service.go"):       serviceGo,
				pathJoin(_PathServices, "service_funcs.go"): serviceFuncsGo,
			},
		}
	}
	return defaultTemplate
}

func (p *iosolatorTemplate) Create() (err error) {
	if err = p.mkDirs(); err != nil {
		return
	}
	return p.mkFiles()
}

func (p *iosolatorTemplate) mkDirs() (err error) {
	for k, v := range p.PathMap {
		fmt.Println("create folder:", k, v)
		if err = p.mkdirAll(v); err != nil {
			return
		}
	}
	return
}

func (p *iosolatorTemplate) mkdirAll(name string) (err error) {
	return files.MkCommonDirAll(pathJoin(p.GoPath, _pathSrc, p.AppPath, name))
}

func (p *iosolatorTemplate) mkFiles() (err error) {
	for k, v := range p.FilesMap {
		t := template.Must(template.New(k).Parse(v))
		var buf bytes.Buffer
		if err = t.Execute(&buf, p.Imports); err != nil {
			return
		}
		filename := pathJoin(p.GoPath, _pathSrc, p.AppPath, k)
		if _, err = files.WriteFile(filename, buf.String()); err != nil {
			return
		}
	}
	return
}

func pathJoin(pathes ...string) string {
	return path.Join(pathes...)
}
