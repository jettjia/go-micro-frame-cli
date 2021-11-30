package docker

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/constant"
	"github.com/jettjia/go-micro-frame-cli/util"
	"strings"
)

// RunDocker 安装 docker
func RunDocker() {
	mlog.Print("init docker:" + constant.DockerVersion + " start...")

	// install docker
	_, err := gproc.ShellExec("sudo yum install docker-ce-" + constant.DockerVersion + " -y")
	if err != nil {
		mlog.Fatal("yum install docker-ce err: ", err)
		return
	}

	_, err = gproc.ShellExec("sudo mkdir -p /data/docker")
	if err != nil {
		mlog.Fatal("sudo mkdir -p /data/docker err: ", err)
		return
	}

	// 启动
	_, err = gproc.ShellExec("systemctl daemon-reload && systemctl enable --now docker")
	if err != nil {
		mlog.Fatal("docker start err: ", err)
		return
	}

	// 写入 daemon.json
	dockerDaemonStr := `{
	"registry-mirrors": ["http://f1361db2.m.daocloud.io"],
	"insecure-registries":["10.4.7.100:85"],
	"graph": "/data/docker"
}`
	strings.Replace(dockerDaemonStr, "10.4.7.100", util.GetOutboundIP(), -1)

	err = gfile.PutContents("/etc/docker/daemon.json", dockerDaemonStr)

	// 重启
	_, err = gproc.ShellExec("systemctl restart docker")
	if err != nil {
		mlog.Fatal("docker restart err: ", err)
		return
	}

	mlog.Print("done!")
}
