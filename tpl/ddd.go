// MIT License

// Copyright (c) 2016 rutcode-go

package tpl

import (
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
	_DDDPathLogic      = "logics"
)

type DDDTemplate struct {
	GoPath  string
	AppPath string
}

var (
	dddTemplate *DDDTemplate
)

func newDDDTemplate(goPath, appPath string) *DDDTemplate {
	if dddTemplate == nil {
		dddTemplate = new(DDDTemplate)
		dddTemplate.GoPath = goPath
		dddTemplate.AppPath = appPath
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
	if err = p.mkdirAll(_DDDPathLogic); err != nil {
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
