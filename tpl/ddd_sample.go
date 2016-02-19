package tpl

import (
	"bytes"
	"html/template"
	"path"

	"github.com/go-rut/files"
)

const (
	_pathSrc = "/src/"

	_DDDMainGo = `
package main

import (
    "{{.RepoLogicPath}}"
    "{{.HandlersPath}}"
)

func main() {
    repoGoRut := repo.NewGoRutRepo()
    repoHelloWorld := repo.NewHelloWorldRepo()

    sampleHandler := handlers.NewSampleHandler(repoGoRut)
    sampleHandler.SayHelloOK()

    sampleHandler = handlers.NewSampleHandler(repoHelloWorld)
    sampleHandler.SayHelloOK()
}
`

	_DDDRepositorySample = `
package repository

type SampleRepository interface {
    SayHello()
}
`

	_DDDRepoGoRut = `package repo

import (
    "fmt"

    "{{.RepositoryPath}}"
)

type GoRutRepo struct {
    repository.SampleRepository
}

func NewGoRutRepo() (r *GoRutRepo) {
    r = new(GoRutRepo)
    return
}

func (p *GoRutRepo) SayHello() {
    fmt.Println("Hello, repo gorut! - It's gorut world.")
}
`
	_DDDRepoHelloWorld = `package repo

import (
    "fmt"

    "{{.RepositoryPath}}"
)

type HelloWorldRepo struct {
    repository.SampleRepository
}

func NewHelloWorldRepo() (repo *HelloWorldRepo) {
    repo = new(HelloWorldRepo)
    return
}

func (p *HelloWorldRepo) SayHello() {
    fmt.Println("Hello, repo world! - It's hello world.")
}
`

	_DDDSampleHandler = `package handlers

import (
    "{{.RepositoryPath}}"
)

var sampleHandler *SampleHandler

type SampleHandler struct {
    SampleRepository repository.SampleRepository
}

func NewSampleHandler(sampleRepository repository.SampleRepository) *SampleHandler {
    if sampleHandler == nil {
        sampleHandler = new(SampleHandler)
    }
    sampleHandler.SampleRepository = sampleRepository
    return sampleHandler
}

func (p *SampleHandler) SayHelloOK() {
    p.SampleRepository.SayHello()
}
`
)

type DDDSampleTemplate struct {
	GoPath  string
	AppPath string

	HandlersPath   string
	RepoLogicPath  string
	RepositoryPath string
}

var (
	dddFilesMap       = make(map[string]string)
	dddSampleTemplate *DDDSampleTemplate
)

func init() {
	dddFilesMap["main.go"] = _DDDMainGo
	dddFilesMap[_DDDPathHandlers+"handler_sample.go"] = _DDDSampleHandler
	dddFilesMap[_DDDPathRepoLogic+"repo_gorut.go"] = _DDDRepoGoRut
	dddFilesMap[_DDDPathRepoLogic+"repo_hello_world.go"] = _DDDRepoHelloWorld
	dddFilesMap[_DDDPathRepository+"repository_sample.go"] = _DDDRepositorySample
}

func newDDDSampleTemplate(goPath, appPath string) *DDDSampleTemplate {
	if dddSampleTemplate == nil {
		dddSampleTemplate = new(DDDSampleTemplate)
		dddSampleTemplate.GoPath = goPath
		dddSampleTemplate.AppPath = appPath
		dddSampleTemplate.HandlersPath = path.Join(appPath, _DDDPathHandlers)
		dddSampleTemplate.RepoLogicPath = path.Join(appPath, _DDDPathRepoLogic)
		dddSampleTemplate.RepositoryPath = path.Join(appPath, _DDDPathRepository)
	}
	return dddSampleTemplate
}

func (p *DDDSampleTemplate) Create() (err error) {
	ddd := new(DDDTemplate)
	ddd.AppPath = p.AppPath
	ddd.GoPath = p.GoPath
	if err = ddd.mkDirs(); err != nil {
		return
	}
	if err = p.mkFiles(); err != nil {
		return
	}
	return
}

func (p *DDDSampleTemplate) mkFiles() (err error) {
	for k, v := range dddFilesMap {
		t := template.Must(template.New(k).Parse(v))
		var buf bytes.Buffer
		if err = t.Execute(&buf, p); err != nil {
			return
		}

		filePath := path.Join(p.GoPath, _pathSrc, p.AppPath, k)
		if _, err = files.WriteFile(filePath, buf.String()); err != nil {
			return
		}
	}
	return
}
