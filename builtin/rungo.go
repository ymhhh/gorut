package builtin

import (
	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
)

func (p *GoRut) cmdRunGo(c *cli.Context) {

	p.readConfig()

	var err error

	if len(c.Args()) == 0 {
		err = ErrCommandNotExists
		p.printError(err)
		return
	}

	err = p.runGo(c.Args())

	return
}

func (p *GoRut) runGo(cArgs []string) (err error) {

	var a []interface{}
	a = append(a, cArgs)

	if err = sh.Command("go", a...).Run(); err != nil {
		p.printError(err)
		return
	}
	return
}
