package cmd

import (
	"fmt"
	"github.com/jettjia/go-micro-frame-cli/constant"
	"github.com/spf13/cobra"
	"log"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current binary version info",
	Long:  `All software has versions. This is go-micro-frame's.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func PrintVersion() {
	log.Println(fmt.Sprintf(constant.PROJECTNAME+":%q\n", constant.VERSION))
}
