package cmd

import (
	"log"

	"github.com/gogf/gf/v2/os/gproc"
	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print go-micro-frame version and environment info",
	Long: `Print go-micro-frame version and environment info. This is useful in go-micro-frame bug reports.

If you add the -v flag, you will get a full dependency list.
`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintVersion()

		goenv()
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}

func goenv() {
	result, err := gproc.ShellExec("go env")
	if err != nil{
		return
	}
	if result == "" {
		return
	}

	log.Println("go env: ")
	log.Println(result)
}