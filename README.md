go-micro-frame cli tool

# install
```
git clone git@github.com:jettjia/go-micro-frame-cli.git
cd go-micro-frame-cli 

go build

go-micro-frame-cli install
```



# Command

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
  run         Install common service, like go-micro-frame-cli run mysql
  start       A brief description of your command
  version     Show current binary version info

Flags:
      --config string   config file (default is $HOME/.go-micro-frame-cli.yaml)
  -h, --help            help for go-micro-frame-cli
  -t, --toggle          Help message for toggle

Use "go-micro-frame-cli [command] --help" for more information about a command.
```

命令说明

```
  build       编译项目
  docker      生成 Dockerfile 文件，方便快速构建项目
  drone       生成 .drone.yaml 文件，方便项目快速cicd构建
  env         输出go env 的运行环境，go-micro-frame 的框架版本
  gen         根据表名，自动化生成代码 ORM model,service, repository, handler, pb等
  gofmt       gofmt，规范项目代码格式
  help        Help about any command
  init        创建一个空的基于 go-micro-frame 框架的项目
  install     安装 go-micro-frame-cli 命令到 win/linux 的bin下，全局使用命令
  run         [ops]快速安装一个开发环境，方便开发、测试、运维； 目前只支持 linux环境，作者使用 centos7环境
  			 已经支持的有：mysql5.7/redis6.2/rabbitmq3.3.7/go1.16.7
  start       [ops]快速启动，run 已经运行中的环境，这里不会覆盖。会保留数据
  version     查看 go-micro-frame 框架版本
```

