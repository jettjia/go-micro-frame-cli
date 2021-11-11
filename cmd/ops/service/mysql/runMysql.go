package mysql

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
	"github.com/jettjia/go-micro-frame-cli/util"
)

// RunMysql 安装mysql 5.7
func RunMysql() {
	mlog.Print("init mysql:" + constant.MysqlVersion + " start...")

	// docker pull image
	_, err := gproc.ShellExec("sudo docker pull mysql:" + constant.MysqlVersion)
	if err != nil {
		mlog.Fatal("pull mysql image err", err)
	}

	// docker run mysql
	_, err = gproc.ShellExec(`
sudo docker run -p 3306:3306 --name `+constant.MysqlName+` \
-v /mydata/mysql/log:/var/log/mysql \
-v /mydata/mysql/data:/var/lib/mysql \
-v /mydata/mysql/conf:/etc/mysql \
-e MYSQL_ROOT_PASSWORD=root \
-d mysql:`+constant.MysqlVersion)

	if err != nil {
		mlog.Fatal("run mysql err", err)
	}

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
	_, err = gproc.ShellExec("sudo docker restart "+ constant.MysqlName)
	if err != nil {
		mlog.Fatal("docker restart mysql err", err)
	}

	mlog.Print("done!")
}
