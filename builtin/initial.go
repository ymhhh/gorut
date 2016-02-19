package builtin

import (
	"path"

	"github.com/codegangsta/cli"
	"github.com/go-rut/files"
)

func (p *GoRut) cmdInitial(c *cli.Context) {

	p.readConfig()

	goSrc := path.Dir(p.Config.Envs.GoPath + "/src/")
	if !files.IsDirExists(goSrc) {
		if err := files.MkCommonDirAll(goSrc); err != nil {
			panic(err)
		}
	}

	goPkg := path.Dir(p.Config.Envs.GoPath + "/pkg/")
	if !files.IsDirExists(goPkg) {
		if err := files.MkCommonDir(goPkg); err != nil {
			panic(err)
		}
	}

	goBin := path.Dir(p.Config.Envs.GoPath + "/bin/")
	if !files.IsDirExists(goBin) {
		if err := files.MkCommonDir(goBin); err != nil {
			panic(err)
		}
	}
	return
}
