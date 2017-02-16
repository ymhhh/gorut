// MIT License

// Copyright (c) 2016 rutcode-go

package ddd_sample

import (
	"path"

	"github.com/go-rut/files"
)

const (
	_DDDPathConf       = "conf"
	_DDDPathRepository = "repository"
	_DDDPathRepoLogic  = "repository/repo"
	_DDDPathDomain     = "domain"
	_DDDPathServices   = "services"
	_DDDPathDocs       = "docs"
	_DDDPathCommon     = "common"
)

type DDDTemplate struct {
	GoPath  string
	AppPath string
}

var dddTemplate *DDDTemplate

func NewDDDFoldersTemplate(goPath, appPath string) *DDDTemplate {
	if dddTemplate == nil {
		dddTemplate = &DDDTemplate{
			GoPath:  goPath,
			AppPath: appPath,
		}
	}
	return dddTemplate
}

func (p *DDDTemplate) Create() error {
	return p.mkDirs()
}

func (p *DDDTemplate) mkDirs() (err error) {
	if err = p.mkdirAll(_DDDPathConf); err != nil {
		return
	}
	if err = p.mkdirAll(_DDDPathRepoLogic); err != nil {
		return
	}
	if err = p.mkdirAll(_DDDPathDomain); err != nil {
		return
	}
	if err = p.mkdirAll(_DDDPathServices); err != nil {
		return
	}
	if err = p.mkdirAll(_DDDPathDocs); err != nil {
		return
	}
	if err = p.mkdirAll(_DDDPathCommon); err != nil {
		return
	}
	return
}

func (p *DDDTemplate) mkdirAll(name string) (err error) {
	return files.MkCommonDirAll(pathJoin(p.GoPath, _pathSrc, p.AppPath, name))
}

func pathJoin(pathes ...string) string {
	return path.Join(pathes...)
}
