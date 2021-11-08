package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jettjia/go-micro-frame-cli/cmd/docker"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "create a docker image for current project",
	Long:  `create a docker image for current project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please enter file, like project-image")
			return
		}
		docker.Run(args)
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
