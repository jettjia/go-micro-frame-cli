package cmd

import (
	"github.com/spf13/cobra"
	"log"

	"github.com/jettjia/go-micro-frame-cli/cmd/initialize"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create and initialize an empty project",
	Long:  `create and initialize an empty project...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("Please enter project name")
			return
		}
		initialize.InitSrv(args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
