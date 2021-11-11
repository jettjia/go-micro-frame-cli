package ops

import "github.com/jettjia/go-micro-frame-cli/cmd/ops/service/mysql"

func RunOps(args []string) {
	serviceName := args[0]

	if serviceName == "mysql" {
		mysql.RunMysql()
	}
}