package cmd

import (
	"runtime"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/cobra"

	"github.com/jettjia/go-micro-frame-cli/cmd/ops"
	"github.com/jettjia/go-micro-frame-cli/constant"
)

var (
	serviceList = `Install common service, like:
	go-micro-frame-cli run mysql
	go-micro-frame-cli run redis
	go-micro-frame-cli run rabbit
	go-micro-frame-cli run es

	go-micro-frame-cli run consul
	go-micro-frame-cli run jaeger
	go-micro-frame-cli run nacos
	go-micro-frame-cli run kong

	go-micro-frame-cli run gogs
	go-micro-frame-cli run harbor
	go-micro-frame-cli run drone

	go-micro-frame-cli run go
	go-micro-frame-cli run docker
	go-micro-frame-cli run docker-compose
`
)
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Install common service, like go-micro-frame-cli run mysql",
	Long:  serviceList,
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			mlog.Print("The run command must be in linux")
			return
		}

		if len(args) == 0 {
			helpRun()
			return
		}

		ops.RunOps(args)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func helpRun() {
	mlog.Print(gstr.TrimLeft(`
USAGE
    go-micro-frame-cli run xx

EXAMPLES
	go-micro-frame-cli run mysql		[Initialize mysql,` + constant.MysqlVersion + `]
	go-micro-frame-cli run redis		[Initialize redis,` + constant.RedisVersion + `]
	go-micro-frame-cli run rabbit		[Initialize rabbit, ` + constant.RabbitmqVersion + `]
	go-micro-frame-cli run es		[Initialize elasticsearch, ` + constant.EsVersion + `]

	go-micro-frame-cli run consul		[Initialize consul,` + constant.ConsulVersion + `]
	go-micro-frame-cli run nacos		[Initialize nacos,` + constant.NacosVersion + `]
	go-micro-frame-cli run jaeger		[Initialize jaeger,` + constant.JaegerVersion + `]
	go-micro-frame-cli run kong		[Initialize kong,` + constant.KongVersion + `]

	go-micro-frame-cli run gogs		[Initialize gogs, ` + constant.GogsVersion + `]
	go-micro-frame-cli run harbor		[Initialize harbor, ` + constant.HarborVersion + `]
	go-micro-frame-cli run drone		[Initialize drone, ` + constant.DroneVersion + `]

	go-micro-frame-cli run go		[Initialize go env, ` + constant.GOVERSION + `]
	go-micro-frame-cli run docker		[Initialize docker, ` + constant.DockerVersion + `]
	go-micro-frame-cli run docker-compose	[Initialize docker-compose, ` + constant.DockerComposeVersion + `]
`))
}
