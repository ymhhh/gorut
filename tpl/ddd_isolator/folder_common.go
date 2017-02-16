// MIT License

// Copyright (c) 2016 rutcode-go

package ddd_isolator

const (
	commonGo = `package common

import (
	"errors"

	"github.com/gogap/isolator"
)

// GenerateReposIsolatorWithOptions get isolator with repos and options
func GenerateReposIsolatorWithOptions(
	objects []isolator.Object, nullObjects []isolator.Object, opts ...isolator.SessionOption) (
	isor *isolator.Isolator, err error) {

	if isor, err = GenerateReposIsolator(objects...); err != nil {
		return
	}

	err = InjectIsolatorSessionOptions(isor, nullObjects, opts...)
	return
}

// GenerateReposIsolator get isolator with repos
func GenerateReposIsolator(objects ...isolator.Object) (isor *isolator.Isolator, err error) {
	if len(objects) == 0 {
		return nil, errors.New("initial repo isolator object is nil")
	}
	objBuilder := isolator.NewClassicObjectBuilder()
	if err = objBuilder.RegisterObjects(objects...); err != nil {
		return
	}

	isor = isolator.NewIsolator(isolator.IsolatorObjectBuilder(objBuilder))
	return
}

// InjectIsolatorSessionOptions inject isolator options
func InjectIsolatorSessionOptions(isor *isolator.Isolator, objects []isolator.Object, opts ...isolator.SessionOption) error {
	if len(objects) == 0 {
		return errors.New("inject isolator objects is nil")
	}

	if len(opts) == 0 {
		return errors.New("inject isolator options is nil")
	}

	if isor == nil {
		return errors.New("initial isolator is nil")
	}

	isor.ObjectsSessionOptions(
		objects,
		opts...,
	)
	return nil
}`

	initGo = `package common

func Init() error {
	return nil
}
`
)
