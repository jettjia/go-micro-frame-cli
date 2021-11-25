package mysql

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
)

func StartMysql() {
	mlog.Print("start mysql5.7...")
	_, _ = gproc.ShellExec("sudo docker restart mysql-jett")

	mlog.Print("done!")
}
