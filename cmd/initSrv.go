package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/jettjia/go-micro-frame-cli/cmd/initialize"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create and initialize an empty project",
	Long:  `create and initialize an empty project...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please enter project name")
			return
		}
		initialize.InitSrv(args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
