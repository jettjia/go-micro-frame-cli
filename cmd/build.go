package cmd

import (
	"fmt"
	"github.com/jettjia/go-micro-frame-cli/cmd/build"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build go project",
	Long:  `build go project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please enter file, like main.go")
			return
		}
		//处理传入的参数
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Please enter output file name, like main")
			return
		}
		build.Run(args, name)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// 接受传入的参数：name， 短拼是：n， 默认值是：main， 提示信息是： 编译后文件名字
	buildCmd.Flags().StringP("name", "n", "main", "output file name")
}
