package builtin

import (
	"path"
	"syscall"

	"github.com/codegangsta/cli"
	"github.com/go-rut/files"
	"github.com/go-rut/gorut/tpl"
)

const (
	cmdCreateForce    = "f"
	cmdCreateTemplate = "t"
)

func (p *GoRut) cmdCreate(c *cli.Context) {

	p.readConfig()

	appPath := path.Join(p.Config.Envs.GoPath, "src", p.Config.Project)

	var err error

	isForce := c.Bool(cmdCreateForce)
	if isForce {
		if err = files.DeleteDir(appPath); err != nil {
			p.printError(err)
			return
		}
	}

	if !files.IsDirExists(appPath) {
		p.printInfo("app path not exists, it will be created")
		if err = files.MkCommonDirAll(appPath); err != nil {
			p.printError(err)
			return
		}
	}

	cTpl := c.String(cmdCreateTemplate)
	if cTpl != "" {
		template := tpl.NewTemplate(cTpl, p.Config.Envs.GoPath, p.Config.Project)
		if template == nil {
			err = ErrTemplateNameNotExists
			p.printError(err)
			return
		}
		if err = template.Create(); err != nil {
			p.printError(err)
			return
		}

		syscall.Exec("gofmt", []string{"-w"}, nil)
	}
	return
}
