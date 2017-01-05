# gorut
it is a tool for managing the go environment and project

## Introduction

* [WIKI:en](https://en.wikipedia.org/wiki/Domain-driven_design)


## Installation

```bash
go get -u github.com/go-rut/gorut
cd $GOPATH/src/github.com/go-rut/gorut
go build
mv gorut /usr/local/bin/
```

------

## GoRut Usage

```
NAME:
   gorut - it is a tool for managing the go environment and project

USAGE:
   gorut [global options] command [command options] [arguments...]

VERSION:
   0.0.2

AUTHOR(S):
   Henry Huang <hhh@rutcode.com>

COMMANDS:
   init		Initial go environment and go source path
   create	Create a new project
   deps		Get dependence of the project
   go		Run command go - link go: ./gorut go [go_subcommand args]
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Config Example

Key|Information|Sample
:-:|:-:|:-:
project|Path: github.com/go-rut; ProjectName: gorut_test|github.com/go-rut/gorut/example/gorut_test
gopath|Go Source Path|/Users/henryhuang/code/golang
cgo_enabled| C Go Enabled Flag| true or false
deps| Project Dependences | Support Git or Go


```json
{
  "project": "github.com/go-rut/gorut/example/gorut_test",
  "envs": {
    "gopath": "/Users/henryhuang/code/golang",
    "cgo_enabled": false
  },
  "deps": [
    {
      "name": "gorut-utils",
      "repo": "https://github.com/go-rut/utils",
      "vcs": "git",
      "options": {
        "flags": [],
        "remote": "origin",
        "branch": "master"
      }
    },
    {
      "name": "gorut-fsm",
      "repo": "github.com/go-rut/fsm",
      "vcs": "go",
      "options": {
        "flags": [
          "-u"
        ]
      }
    }
  ]
}
```

## Step 1 - Initial go environment and go source path

```bash
gorut init
```

**It will create go source path**

* $gopath/src
* $gopath/pkg
* $gopath/bin

## Step 2 - Example config

```bash
cd example
```

[example](example/gorut.conf)


## Step 3 - Create a new project with gorut.config

Key|Information
:-:|:-:
-f|if it is true, it will delete the old project, then create an new project
-t|create the project with template, default is null. buildin: ddd, ddd_sample

### Sample 1 - Project Without Code

If create an new project like this:

```bash
gorut create
```

It will create an new project (ProjectPath: $gopath/src/github.com/go-rut/gorut_test) without any code.


### Sample 2 - DDD(Domain Driven Design) Project Without Code

If create an new ddd project like this:

```bash
gorut create -f -t=ddd
```

It will create an new project (ProjectPath: $gopath/src/github.com/go-rut/gorut_test) without code.

But it has ddd directories:

```
$ProjectPath/common/
$ProjectPath/conf/
$ProjectPath/docs/
$ProjectPath/domain/
$ProjectPath/services/
$ProjectPath/logics/
$ProjectPath/models/
$ProjectPath/repository/
```

### Sample 3 - DDD Sample Project

If create an new ddd project like this:

```bash
cd example
gorut create -f -t=ddd_sample
```

It will create project (ProjectPath: $gopath/src/github.com/go-rut/gorut_test) with ddd sample code.

![ProjectFiles](images/ddd_sample.png)

### Other Project Template

If you have other good practice, write your template into tpl path.

Your struct must be GoRutTemplate:

```go
type GoRutTemplate interface {
  Create() error
}
```

## Step 4 - Move sample config into project path

```bash
mv gorut.conf /Users/henryhuang/code/golang/src/github.com/go-rut/gorut/example/gorut_test/
```

## Step 5 - Run ddd sample

* command : gorut go
* Run gorut go - link go: go go_subcommand [args]

```bash
cd gorut_test/
gorut go run main.go
```

![ProjectFiles](images/ddd_sample_run.png)

## Dependences

If your project need dependences, you can write into config. It Supported Git And Go.

### Get dependences

```bash
gorut deps
```

It will get dependences like this:

![ProjectFiles](images/dependences.png)
