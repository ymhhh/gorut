// MIT License

// Copyright (c) 2016 rutcode-go

package builtin

import (
	"path"
	"regexp"

	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
	"github.com/go-rut/files"
)

const (
	cmdDepsDelete = "d"

	deleteTypeAll = "all"

	VcsGo  = "go"
	VcsGit = "git"
)

func (p *GoRut) cmdDeps(c *cli.Context) {

	p.readConfig()

	dFlag := c.String(cmdDepsDelete)

	var err error

	if dFlag != "" {
		srcPath := path.Join(p.Config.Project, "src")
		switch dFlag {
		case deleteTypeAll:
			if err = files.DeleteDir(srcPath); err != nil {
				p.printError(err)
				return
			}
		default:
			for _, v := range p.Config.Deps {
				srcPath = path.Join(srcPath, getRepoPath(v.Vcs, v.Repo))
				if err = files.DeleteDir(srcPath); err != nil {
					p.printError(err)
				}
				p.printInfo("delete path: " + srcPath)
			}
			return
		}
		return
	}

	for _, v := range p.Config.Deps {

		p.printInfo("start get deps:", v.Name, v.Repo)

		srcPath := path.Join(p.Config.Envs.GoPath, "src", getRepoPath(v.Vcs, v.Repo))

		if v.Vcs == VcsGo {

			var options []interface{}
			if len(v.Options.Flags) != 0 {
				options = append(options, "get", v.Repo)
			} else {
				options = append(options, "get", v.Options.Flags, v.Repo)
			}

			if err = sh.Command("go", options...).Run(); err != nil {
				p.printError(err)
				return
			}

		} else if v.Vcs == VcsGit {

			if !files.IsDirExists(srcPath) {
				var options []interface{}

				options = append(options, "clone", v.Repo, srcPath)
				if v.Options.Branch != "" {
					options = append(options, "-b", v.Options.Branch)
				}

				if err = sh.Command("git", options...).Run(); err != nil {
					p.printError(err)
					return
				}

			} else {
				if err = sh.Command("cd", srcPath).Run(); err != nil {
					p.printError(err)
					return
				}

				if v.Options.Branch != "" {
					if err = sh.Command("git", "checkout", v.Options.Branch).Run(); err != nil {
						p.printError(err)
						return
					}
				}

				if v.Options.Remote != "" && v.Options.Branch != "" {
					if err = sh.Command("git", "pull", v.Options.Remote, v.Options.Branch).Run(); err != nil {
						p.printError(err)
						return
					}
				}
			}
		}
	}

	return
}

func getRepoPath(t, path string) string {
	if t == VcsGo {
		return path
	}

	re := regexp.MustCompile("^((https|http)?://)")

	bytes := re.ReplaceAll([]byte(path), []byte(""))

	return string(bytes)
}
