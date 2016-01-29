package builtin

import (
	"fmt"
)

const (
	_PTypeCritical = iota
	_PTypeError
	_PTypeWarning
	_PTypeInfo

	SprintString = "\x1b[0;%dm%s %s %v\x1b[0m"

	_PStringCritical = "<C>"
	_PStringError    = "<E>"
	_PStringWarning  = "<W>"
	_PStringInfo     = "<I>"

	_ColorCritical = 35
	_ColorError    = 31
	_ColorWarning  = 33
	_ColorInfo     = 36
)

func (p *GoRut) println(pType int, msg ...interface{}) {
	var color int
	var pStr string
	switch pType {
	case _PTypeCritical:
		color = _ColorCritical
		pStr = _PStringCritical
	case _PTypeError:
		color = _ColorError
		pStr = _PStringError
	case _PTypeWarning:
		color = _ColorWarning
		pStr = _PStringWarning
	default:
		color = _ColorInfo
		pStr = _PStringInfo
	}

	pStr = p.fmtPrintStr(color, pStr, msg)
	fmt.Println(pStr)
}

func (p *GoRut) fmtPrintStr(pColor int, ptype string, msg ...interface{}) string {
	return fmt.Sprintf(SprintString, pColor, p.CliApp.Name, ptype, msg)
}

func (p *GoRut) printError(err error) {
	if err != nil {
		p.println(_PTypeError, err.Error())
	}
}

func (p *GoRut) printInfo(msg ...interface{}) {
	p.println(_PTypeInfo, msg)
}
