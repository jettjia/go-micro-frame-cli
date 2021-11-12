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
	go-micro-frame-cli run consul
	go-micro-frame-cli run jaeger
	go-micro-frame-cli run nacos
	go-micro-frame-cli run konga
	go-micro-frame-cli run go
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
	go-micro-frame-cli run consul		[Initialize run,latest]
	go-micro-frame-cli run nacos		[Initialize nacos,latest]
	go-micro-frame-cli run jaeger		[Initialize jaeger,latest]
	go-micro-frame-cli run konga		[Initialize konga,latest]
	go-micro-frame-cli run go		[Initialize go env, 1.16.7]
`))
}
