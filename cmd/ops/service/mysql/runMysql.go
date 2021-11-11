package mysql

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/util"
)

// RunMysql 安装mysql 5.7
func RunMysql() {
	mlog.Print("init mysql5.7 start...")

	// docker pull image
	_, _ = gproc.ShellExec("sudo docker pull mysql:5.7")

	// docker run mysql
	_, _ = gproc.ShellExec(`
sudo docker run -p 3306:3306 --name mysql \
-v /mydata/mysql/log:/var/log/mysql \
-v /mydata/mysql/data:/var/lib/mysql \
-v /mydata/mysql/conf:/etc/mysql \
-e MYSQL_ROOT_PASSWORD=root \
-d mysql:5.7
`)

	// write config
	configStr := `
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8
[mysqld]
init_connect='SET collation_connection = utf8_unicode_ci'
init_connect='SET NAMES utf8'
character-set-server=utf8
collation-server=utf8_unicode_ci
skip-character-set-client-handshake
skip-name-resolve
`
	util.WriteStringToFileMethod("/mydata/mysql/conf/my.conf", configStr)

	// start docker
	_, _ = gproc.ShellExec("sudo docker restart mysql")

	mlog.Print("done!")
}
