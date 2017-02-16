// MIT License

// Copyright (c) 2016 rutcode-go

package ddd_isolator

const (
	logicsGo = `package logic

import (
	"fmt"

	"{{.PathRepoes}}"
)

type Logics interface {
	SayHello() (interface{}, error)
}

type defaultLogics struct{}

var logics *defaultLogics

func NewLogics() Logics {
	if logics == nil {
		logics = &defaultLogics{}
	}

	return logics
}

func (p *defaultLogics) SayHello() (interface{}, error) {
	return func(userRepo repo.UserRepository) error {
		user := userRepo.GetUser()

		fmt.Println("logic say hello:", user)

		return nil
	}, nil
}
`
)
