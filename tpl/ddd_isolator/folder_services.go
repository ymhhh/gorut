package ddd_isolator

const (
	serviceGo = `package services

import (
 	"github.com/gogap/isolator"

	"{{.PathLogics}}"
)

var service *Service

type Service struct {
	Isolator *isolator.Isolator
	Logics   logic.Logics
}

// NewProductHandler return default product handler
func NewService(isor *isolator.Isolator) *Service {
	if service == nil {
		service = &Service{
			Logics: logic.NewLogics(),
		}
	}
	service.Isolator = isor
	return service
}
`

	serviceFuncsGo = `package services

import (
	"golang.org/x/net/context"
)

type ReqNull struct{}

func (p *Service) SayHello(_ context.Context, _ *ReqNull) (interface{}, error) {

	return nil, p.Isolator.Invoke(p.Logics.SayHello).End(
		func(e error) {
			if e != nil {
				return
			}
			return
		})
}
`
)
