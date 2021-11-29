package mysql

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
)

// RunMysql 安装mysql 5.7
func RunMysql() {
	mlog.Print("init mysql:" + constant.MysqlVersion + " start...")

	// pull image
	has, _ := gproc.ShellExec("docker images -q mysql:"+constant.MysqlVersion)
	if has == "" {
		_, err := gproc.ShellExec("sudo docker pull mysql:" + constant.MysqlVersion)
		if err != nil {
			mlog.Fatal("pull mysql image err: ", err)
			return
		}
	}


	// docker run mysql
	_, err := gproc.ShellExec(`
sudo docker run -p 3306:3306 --name `+constant.MysqlName+` \
-v /mydata/mysql/log:/var/log/mysql \
-v /mydata/mysql/data:/var/lib/mysql \
-v /mydata/mysql/conf:/etc/mysql \
-e MYSQL_ROOT_PASSWORD=root \
-d mysql:`+constant.MysqlVersion)

	if err != nil {
		mlog.Fatal("run mysql err: ", err)
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
	err = gfile.PutContents("/mydata/mysql/conf/my.conf", configStr)
	if err != nil {
		mlog.Fatal("write /mydata/mysql/conf/my.conf err: ", err)
	}

	// start docker
	_, err = gproc.ShellExec("sudo docker restart "+ constant.MysqlName)
	if err != nil {
		mlog.Fatal("docker restart mysql err: ", err)
	}

	mlog.Print("The Mysql account password is root/root，Please keep it properly")
	mlog.Print("done!")
}
