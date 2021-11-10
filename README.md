go-micro-frame cli tool

# install
```
git clone git@github.com:jettjia/go-micro-frame-cli.git
cd go-micro-frame-cli 

go build -ldflags "-w -s" -o go-micro-frame-cli main.go
or
go build -ldflags "-w -s" -o go-micro-frame-cli.exe main.go

go-micro-frame-cli install
```



# Command

todo

```
$ go-micro-frame-cli

Usage:
  go-micro-frame-cli [command]

Available Commands:
  build       build go project
  completion  generate the autocompletion script for the specified shell
  docker      create a docker image for current project
  drone       create .drone for ci/cd
  env         Print go-micro-frame version and environment info
  gen         automatically generate go files for ORM model,service, repository, handler, pb
  gofmt       gofmt your project
  help        Help about any command
  init        create and initialize an empty project
  install     install gf binary to system (might need root/admin permission)
  version     Show current binary version info

Flags:
      --config string   config file (default is $HOME/.go-micro-frame-cli.yaml)
  -h, --help            help for go-micro-frame-cli
  -t, --toggle          Help message for toggle

Use "go-micro-frame-cli [command] --help" for more information about a command.
```

