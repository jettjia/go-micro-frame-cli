package cmd

import (
	"runtime"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops"
	"github.com/spf13/cobra"
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
	go-micro-frame-cli run mysql		[Initialize mysql,5.7]
	go-micro-frame-cli run redis		[Initialize redis,6.2]
	go-micro-frame-cli run rabbit		[Initialize rabbit,3.7.7-management]
	go-micro-frame-cli run es		[Initialize elasticsearch,7.7.1]

	go-micro-frame-cli run consul		[Initialize run,latest]
	go-micro-frame-cli run nacos		[Initialize nacos,latest]
	go-micro-frame-cli run jaeger		[Initialize jaeger,latest]
	go-micro-frame-cli run kong		[Initialize kong,latest]

	go-micro-frame-cli run gogs		[Initialize gogs]
	go-micro-frame-cli run harbor		[Initialize harbor]
	go-micro-frame-cli run drone		[Initialize drone]

	go-micro-frame-cli run go		[Initialize go env, 1.16.7]
	go-micro-frame-cli run docker		[Initialize docker-ce-19.03.*]
	go-micro-frame-cli run docker-compose	[Initialize docker-compose-1.21.2]
`))
}
