// MIT License

// Copyright (c) 2016 rutcode-go

package tpl

const (
	TemplateDDD       = "ddd"
	TemplateDDDSample = "ddd_sample"
)

type GoRutTemplate interface {
	Create() error
}

func NewTemplate(tplName, goPath, appPath string) GoRutTemplate {
	if tplName == TemplateDDD {
		t := newDDDTemplate(goPath, appPath)
		return t
	} else if tplName == TemplateDDDSample {
		tpl := newDDDSampleTemplate(goPath, appPath)
		return tpl
	}
	return nil
}
