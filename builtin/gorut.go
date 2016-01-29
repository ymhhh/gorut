package builtin

import (
	"github.com/codegangsta/cli"
)

var (
	newGoRut *GoRut
)

type GoRut struct {
	CliApp *cli.App
	Config *GoRutConfig
}

func NewGoRut(appName, author, email, appVersion, description string) (*GoRut, error) {
	if newGoRut == nil {
		newGoRut = new(GoRut)
	}

	cliApp := cli.NewApp()
	cliApp.Name = appName
	cliApp.Author = author
	cliApp.Email = email
	cliApp.Usage = description
	cliApp.Version = appVersion
	cliApp.Commands = newGoRut.commands()
	newGoRut.CliApp = cliApp

	return newGoRut, nil
}
