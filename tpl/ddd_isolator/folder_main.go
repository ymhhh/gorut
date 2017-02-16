// MIT License

// Copyright (c) 2016 rutcode-go

package ddd_isolator

const (
	_DDDMainGo = `package main

import (
	"fmt"

	"github.com/gogap/isolator"
	"golang.org/x/net/context"
	// "github.com/gogap/isolator/extension/xorm"

	"{{.PathCommon}}"
	"{{.PathRepoes}}"
	"{{.PathServices}}"
)

func main() {
	if err := common.Init(); err != nil {
		panic(err)
		return
	}

	sqlIsor, err := generateIsolator(repo.NewSqlRepo(), (*repo.SqlRepo)(nil))
	if err != nil {
		panic(err)
		return
	}

	service := services.NewService(sqlIsor)
	service.SayHello(context.Background(), (*services.ReqNull)(nil))

	redisIsor, err := generateIsolator(repo.NewRedisRepo(), (*repo.RedisRepo)(nil))
	if err != nil {
		panic(err)
		return
	}

	service = services.NewService(redisIsor)
	service.SayHello(context.Background(), (*services.ReqNull)(nil))
}

func generateIsolator(r isolator.Object, os ...isolator.Object) (
	isor *isolator.Isolator, err error) {

	onErrorRollbackFn := func(isor *isolator.Session, err error) {
		fmt.Println("## Rollback ...", isor.CreateTime)
	}

	onSuccessFn := func(isor *isolator.Session) {
		fmt.Println("## Success ...", isor.CreateTime)
	}

	isor, err = common.GenerateReposIsolator(r)
	if err != nil {
		return
	}

	if err = common.InjectIsolatorSessionOptions(
		isor,
		[]isolator.Object(os),
		isolator.SessionOnSuccess(onSuccessFn),
		isolator.SessionOnError(onErrorRollbackFn),

		// Use xorm session
		// common.XormEngines.NewXORMSession(common.DBNameUser, false),
		// isolator.SessionOnSuccess(xorm.OnXORMSessionSuccess),
		// isolator.SessionOnError(xorm.OnXORMSessionError),
	); err != nil {
		return
	}
	return
}
`
)
