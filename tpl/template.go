// MIT License

// Copyright (c) 2016 rutcode-go

package tpl

import (
	"github.com/go-rut/gorut/tpl/ddd_isolator"
	"github.com/go-rut/gorut/tpl/ddd_sample"
)

const (
	TemplateDDD         = "ddd"
	TemplateDDDSample   = "ddd_sample"
	TemplateDDDIsolator = "ddd_isolator"
)

type GoRutTemplate interface {
	Create() error
}

func NewTemplate(tplName, goPath, appPath string) GoRutTemplate {
	switch tplName {
	case TemplateDDD:
		return ddd_sample.NewDDDFoldersTemplate(goPath, appPath)
	case TemplateDDDSample:
		return ddd_sample.NewDDDSampleTemplate(goPath, appPath)
	case TemplateDDDIsolator:
		return ddd_isolator.NewIosolatorTemplate(goPath, appPath)
	default:
		return nil
	}
}
