// MIT License

// Copyright (c) 2016 rutcode-go

package main

import (
	"os"

	"github.com/go-rut/gorut/builtin"
)

var (
	appName     = "gorut"
	author      = "Henry Huang"
	email       = "hhh@rutcode.com"
	version     = "0.0.1"
	description = "it is a tool for managing the go environment and project"
)

func main() {

	gorut, err := builtin.NewGoRut(appName,
		author,
		email,
		version,
		description)
	if err != nil {
		return
	}
	gorut.CliApp.Run(os.Args)
}
