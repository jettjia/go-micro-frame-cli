package cmd

import (
	"fmt"
	"runtime"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops"
	"github.com/spf13/cobra"
)

var (
	serviceList = "Install common services, like mysql, go"
)
var runCmd = &cobra.Command{
	Use:   "run",
	Short: serviceList,
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
	go-micro-frame-cli run mysql		[Initialize service]
	go-micro-frame-cli run go		[Initialize go env, 1.16.7]
`))
}
