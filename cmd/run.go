package cmd

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/jettjia/go-micro-frame-cli/cmd/ops"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Install common services, like mysql, redis...",
	Long:  `Install common services, like mysql, redis...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			helpRun()
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
    go-micro-frame-cli run

EXAMPLES
	go-micro-frame-cli run mysql		[Initialize service]
`))
}
