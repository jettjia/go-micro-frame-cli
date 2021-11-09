package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jettjia/go-micro-frame-cli/cmd/gen"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "automatically generate go files for ORM model,service, repository, handler, pb",
	Long: `automatically generate go files for ORM model,service, repository, handler, pb`,
	Run: func(cmd *cobra.Command, args []string) {
		// todo
		// -h -u -p -d -t -s
		// 解析上面的参数
		// cli gen -h 127.0.0.1 -u root -p root -p 3307 -d zhe_pms -t category -s goods_srv
		gen.Run()
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
