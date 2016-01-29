package tpl

import (
	"path"

	"github.com/go-rut/utils/files"
)

const (
	_DDDPathConf           = "/conf/"
	_DDDPathRepository     = "/repository/"
	_DDDPathRepoLogic      = "/repository/repo/"
	_DDDPathHandlers       = "/handlers/"
	_DDDPathDocs           = "/docs/"
	_DDDPathInfrastructure = "/infrastructure/"
	_DDDPathModels         = "/models/"
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
	appConfigPath := path.Join(p.GoPath, _pathSrc, p.AppPath, _DDDPathConf)
	if err = files.MkCommonDirAll(appConfigPath); err != nil {
		return
	}

	appRepoLogicPath := path.Join(p.GoPath, _pathSrc, p.AppPath, _DDDPathRepoLogic)
	if err = files.MkCommonDirAll(appRepoLogicPath); err != nil {
		return
	}

	appHandlersPath := path.Join(p.GoPath, _pathSrc, p.AppPath, _DDDPathHandlers)
	if err = files.MkCommonDirAll(appHandlersPath); err != nil {
		return
	}

	appDocsPath := path.Join(p.GoPath, _pathSrc, p.AppPath, _DDDPathDocs)
	if err = files.MkCommonDirAll(appDocsPath); err != nil {
		return
	}

	appInfrastructurePath := path.Join(p.GoPath, _pathSrc, p.AppPath, _DDDPathInfrastructure)
	if err = files.MkCommonDirAll(appInfrastructurePath); err != nil {
		return
	}

	appModelsPath := path.Join(p.GoPath, _pathSrc, p.AppPath, _DDDPathModels)
	if err = files.MkCommonDirAll(appModelsPath); err != nil {
		return
	}
	return
}
