package ops

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/cicd/gogs"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/cicd/harbor"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/initGo"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/micro/consul"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/micro/jaeger"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/micro/konga"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/micro/nacos"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/service/docker"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/service/mysql"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/service/rabbitmq"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/service/redis"
)

func RunOps(args []string) {
	serviceName := args[0]

	// 基础服务
	{
		if serviceName == "mysql" {
			mysql.RunMysql()
		}

		if serviceName == "redis" {
			redis.RunRedis()
		}

		if serviceName == "rabbit" || serviceName == "rabbitmq" {
			rabbitmq.RunRabbit()
		}
	}

	// micro service
	{
		if serviceName == "nacos" {
			nacos.RunNacos()
		}
		if serviceName == "jaeger" {
			jaeger.RunJaeger()
		}
		if serviceName == "kong" {
			konga.RunKonga()
		}
		if serviceName == "consul" {
			consul.RunConsul()
		}
		if serviceName == "docker" {
			docker.RunDocker()
		}
		if serviceName == "docker-compose" {
			docker.RunDockerCompose()
		}
	}

	// cicd
	{
		if serviceName == "gogs" {
			gogs.RunGogs()
		}
		if serviceName == "harbor" {
			harbor.RunHarbor()
		}
	}

	// 项目环境
	{
		if serviceName == "go" {
			initGo.RunGo()
		}
	}

	mlog.Print("The Command not found")
}
