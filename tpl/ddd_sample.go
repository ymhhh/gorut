// MIT License

// Copyright (c) 2016 rutcode-go

package tpl

import (
	"bytes"
	"html/template"

	"github.com/go-rut/files"
)

const (
	_pathSrc = "/src/"

	_DDDMainGo = `
package main

import (
    "{{.RepoLogicPath}}"
    "{{.ServicePath}}"
)

func main() {
    repoSQL := repo.NewSqlRepo()
    repoRedis := repo.NewRedisRepo()

    dddHandler := services.NewDDDService(repoSQL)
    dddHandler.SayHello()

    dddHandler = services.NewDDDService(repoRedis)
    dddHandler.SayHello()
}
`

	_DDDUserRepository = `
package repository

import "{{.DomainPath}}"

type UserRepository interface {
    GetUser() (domain.User)
}
`

	_DDDSqlRepo = `package repo

import (
    "fmt"

    "{{.RepositoryPath}}"
    "{{.DomainPath}}"
)

type SqlRepo struct {}

func NewSqlRepo() (repository.UserRepository) {
    return new(SqlRepo)
}

func (p *SqlRepo) GetUser() (domain.User) {
    fmt.Println("Hello, repo sql! - Get user name from sql")
    return domain.User{Id: "1", Name: "SQL"}
}
`
	_DDDRedisRepo = `package repo

import (
    "fmt"

    "{{.RepositoryPath}}"
    "{{.DomainPath}}"
)

type RedisRepo struct{}

func NewRedisRepo() (repository.UserRepository) {
    return new(RedisRepo)
}

func (p *RedisRepo) GetUser() (domain.User){
    fmt.Println("Hello, repo redis! - Get user name from redis")
    return domain.User{Id: "2", Name: "Redis"}
}
`

	_DDDuserHandler = `package services

import (
	"fmt"

    "{{.RepositoryPath}}"
)

var dddService *DDDService

type DDDService struct {
    UserRepository repository.UserRepository
}

func NewDDDService(userRepository repository.UserRepository) *DDDService {
    if dddService == nil {
        dddService = new(DDDService)
    }
    dddService.UserRepository = userRepository
    return dddService
}

func (p *DDDService) SayHello() {
    user := p.UserRepository.GetUser()
    fmt.Println("hello:", user.Name)
}
`

	_DDDUserDomain = `
package domain

type User struct {
    Id string
    Name string
}`
)

type DDDSampleTemplate struct {
	GoPath  string
	AppPath string

	PathMap  map[string]string
	FilesMap map[string]string
}

var dddSampleTemplate *DDDSampleTemplate

func newDDDSampleTemplate(goPath, appPath string) *DDDSampleTemplate {
	if dddSampleTemplate == nil {
		dddSampleTemplate = &DDDSampleTemplate{
			GoPath:  goPath,
			AppPath: appPath,

			PathMap: map[string]string{
				"ServicePath":    pathJoin(appPath, _DDDPathServices),
				"RepoLogicPath":  pathJoin(appPath, _DDDPathRepoLogic),
				"RepositoryPath": pathJoin(appPath, _DDDPathRepository),
				"DomainPath":     pathJoin(appPath, _DDDPathDomain),
			},
			FilesMap: map[string]string{
				"main.go": _DDDMainGo,
				pathJoin(_DDDPathServices, "service_user.go"):      _DDDuserHandler,
				pathJoin(_DDDPathRepoLogic, "repo_sql.go"):         _DDDSqlRepo,
				pathJoin(_DDDPathRepoLogic, "repo_redis.go"):       _DDDRedisRepo,
				pathJoin(_DDDPathRepository, "repository_user.go"): _DDDUserRepository,
				pathJoin(_DDDPathDomain, "user.go"):                _DDDUserDomain,
			},
		}
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
	return p.mkFiles()
}

func (p *DDDSampleTemplate) mkFiles() (err error) {
	for k, v := range p.FilesMap {
		t := template.Must(template.New(k).Parse(v))
		var buf bytes.Buffer
		if err = t.Execute(&buf, p.PathMap); err != nil {
			return
		}
		filename := pathJoin(p.GoPath, _pathSrc, p.AppPath, k)
		_, err = files.WriteFile(filename, buf.String())
		if err != nil {
			return
		}
	}
	return
}
