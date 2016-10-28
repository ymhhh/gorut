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
    repoSQL := repo.NewSqlRepo()
    repoRedis := repo.NewRedisRepo()

    sampleHandler := handlers.NewUserHandler(repoSQL)
    sampleHandler.SayHello()

    sampleHandler = handlers.NewUserHandler(repoRedis)
    sampleHandler.SayHello()
}
`

	_DDDUserRepository = `
package repository

type UserRepository interface {
    GetUserName() (name string)
}
`

	_DDDSqlRepo = `package repo

import (
    "fmt"

    "{{.RepositoryPath}}"
)

type SqlRepo struct {
}

func NewSqlRepo() (repository.UserRepository) {
    return new(SqlRepo)
}

func (p *SqlRepo) GetUserName() (string) {
    fmt.Println("Hello, repo sql! - Get user name from sql")
    return "SQL"
}
`
	_DDDRedisRepo = `package repo

import (
    "fmt"

    "{{.RepositoryPath}}"
)

type RedisRepo struct {
}

func NewRedisRepo() (repository.UserRepository) {
    return new(RedisRepo)
}

func (p *RedisRepo) GetUserName() (name string){
    fmt.Println("Hello, repo redis! - Get user name from redis")
    return "Redis"
}
`

	_DDDuserHandler = `package handlers

import (
	"fmt"

    "{{.RepositoryPath}}"
)

var userHandler *UserHandler

type UserHandler struct {
    UserRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
    if userHandler == nil {
        userHandler = new(UserHandler)
    }
    userHandler.UserRepository = userRepository
    return userHandler
}

func (p *UserHandler) SayHello() {
    username := p.UserRepository.GetUserName()
    fmt.Println("hello", username)
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
	dddFilesMap[_DDDPathHandlers+"handler_user.go"] = _DDDuserHandler
	dddFilesMap[_DDDPathRepoLogic+"repo_sql.go"] = _DDDSqlRepo
	dddFilesMap[_DDDPathRepoLogic+"repo_redis.go"] = _DDDRedisRepo
	dddFilesMap[_DDDPathRepository+"repository_user.go"] = _DDDUserRepository
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
