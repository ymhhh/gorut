package ddd_isolator

const (
	userRepoGo = `package repo

import (
	"github.com/gogap/isolator"

	"{{.PathDomain}}"
)

type UserRepository interface {
	isolator.Object

	GetUser() domain.User
}
`

	userRepoRedisGo = `package repo

import (
	"fmt"

	"github.com/gogap/isolator"

	"{{.PathDomain}}"
)

func NewRedisRepo() UserRepository {
	return &RedisRepo{}
}

type RedisRepo struct {
	session *isolator.Session
}

func (p *RedisRepo) Derive(session *isolator.Session) (obj isolator.Object, err error) {
	return &RedisRepo{
		session: session,
	}, nil
}

func (p *RedisRepo) GetUser() domain.User {
    fmt.Println("Hello, repo Redis! - Get user name from Redis")
    return domain.User{Id: "2", Name: "Redis"}
}
`

	userRepoSqlGo = `package repo

import (
	"fmt"

	"github.com/gogap/isolator"

	"{{.PathDomain}}"
)

import _ "github.com/go-sql-driver/mysql"

func NewSqlRepo() UserRepository {
	return &SqlRepo{}
}

type SqlRepo struct {
	session *isolator.Session
}

func (p *SqlRepo) Derive(session *isolator.Session) (obj isolator.Object, err error) {
	return &SqlRepo{
		session: session,
	}, nil
}

func (p *SqlRepo) GetUser() domain.User {
    fmt.Println("Hello, repo Sql! - Get user name from Sql")
    return domain.User{Id: "1", Name: "Sql"}
}
`
)
