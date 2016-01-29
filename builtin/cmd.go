package builtin

import (
	"github.com/codegangsta/cli"
)

func (p *GoRut) commands() []cli.Command {
	return []cli.Command{
		{
			Name:   "init",
			Usage:  "Initial go environment and go source path",
			Action: p.cmdInitial,
		}, {
			Name:   "create",
			Usage:  "Create a new project",
			Action: p.cmdCreate,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  cmdCreateForce,
					Usage: "if force is true, it will delete the old project, then create a new one",
				}, cli.StringFlag{
					Name:  cmdCreateTemplate,
					Value: "",
					Usage: "create the project with template, default is null",
				},
			},
		}, {
			Name:   "deps",
			Usage:  "Get dependence of the project",
			Action: p.cmdDeps,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  cmdDepsDelete,
					Value: "",
					Usage: "use -d=deps delete dependence in config, or use -d=all to delete GOPATH/src, default: \"\". sample: gorut deps -d=deps",
				},
			},
		}, {
			Name:   "go",
			Usage:  "Run command go - link go: ./gorut go [go_subcommand args]",
			Action: p.cmdRunGo,

			SkipFlagParsing: true,
		},
	}
}
