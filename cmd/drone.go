package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jettjia/go-micro-frame-cli/cmd/cicd"
)

// droneCmd represents the drone command
var droneCmd = &cobra.Command{
	Use:   "drone",
	Short: "create .drone for ci/cd",
	Long: `create .drone for ci/cd`,
	Run: func(cmd *cobra.Command, args []string) {
		cicd.Run(args)
	},
}

func init() {
	rootCmd.AddCommand(droneCmd)
}
