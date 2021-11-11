package ops

import (
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/initGo"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/service/mysql"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops/service/redis"
)

func RunOps(args []string) {
	serviceName := args[0]

	if serviceName == "mysql" {
		mysql.RunMysql()
	}

	if serviceName == "redis" {
		redis.RunRedis()
	}

	if serviceName == "go" {
		initGo.RunGo()
	}
}
