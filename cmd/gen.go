package cmd

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/jettjia/go-micro-frame-cli/cmd/gen"
	"github.com/spf13/cobra"
)

var (
	addr, user, pwd, port, db, table, serverName, protoName string
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "automatically generate go files for ORM model, service, repository, handler, pb",
	Long:  `automatically generate go files for ORM model, service, repository, handler, pb`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(addr) == 0 || len(user) == 0 || len(port) == 0 || len(db) == 0 || len(table) == 0 || len(serverName) == 0 {
			help()
			return
		}
		gen.Run(addr, user, pwd, port, db, table, serverName, protoName)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1", "Enter MySQL addr")
	genCmd.Flags().StringVarP(&user, "user", "u", "root", "Enter MySQL user")
	genCmd.Flags().StringVarP(&pwd, "pwd", "", "root", "Enter MySQL password")
	genCmd.Flags().StringVarP(&port, "port", "p", "3306", "Enter MySQL port")
	genCmd.Flags().StringVarP(&db, "db", "d", "", "Enter MySQL database")
	genCmd.Flags().StringVarP(&table, "table", "t", "", "Enter MySQL table")
	genCmd.Flags().StringVarP(&serverName, "serverName", "s", "", "Enter project server name")
	genCmd.Flags().StringVarP(&protoName, "protoName", "n", "", "Enter project protoName")
}

func help() {
	mlog.Print(gstr.TrimLeft(`
USAGE
    go-micro-frame-cli gen [OPTION]

ARGUMENT
    OPTION
	-a	Enter MySQL addr 
	-u	Enter MySQL user 
	-pwd	Enter MySQL password 
	-p	Enter MySQL port 
	-d	Enter MySQL database 
	-t	Enter MySQL table 
	-s	Enter project server name
	-n	Enter project proto name

EXAMPLES
    go-micro-frame-cli gen -a 10.4.7.71 -u root -pwd root -p 3307 -d zhe_pms -t product -s goods-srv -n goods 
`))
}
