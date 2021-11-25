package harbor

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/jettjia/go-micro-frame-cli/util"
)

// RunHarbor 安装 harbor
func RunHarbor() {
	// down
	_, err := gproc.ShellExec("sudo curl -L https://github.com/goharbor/harbor/releases/download/v2.2.1/harbor-offline-installer-v2.2.1.tgz -o /tmp")
	if err != nil {
		mlog.Fatal("down harbor : ", err)
		return
	}

	gproc.ShellExec("sudo mkdir /opt/harbor")

	gproc.ShellExec("tar -zxf /tmp/harbor-offline-installer-v2.2.1.tgz")

	gproc.ShellExec("mv /tmp/harbor/* /opt/harbor")

	gproc.ShellExec("cp /opt/harbor/harbor.yml.tmpl /opt/harbor/harbor.yml")

	// 修改ip和端口
	gproc.ShellExec(`sed -i "s/hostname: reg.mydomain.com/hostname: ` + util.GetOutboundIP() + `/g"  /opt/harbor/harbor.yml`)
	gproc.ShellExec(`sed -i "s/port: 80/port: 85/g"  /opt/harbor/harbor.yml`)

	// 安装
	gproc.ShellExec("shell /opt/harbor/prepare")
	gproc.ShellExec("shell /opt/harbor/install.sh  --with-chartmuseum")
	gproc.ShellExec("docker-compose up -d")


	mlog.Print("http://ip:85")
	mlog.Print("The harbor account password is admin/Harbor12345，Please keep it properly")
	mlog.Print("done!")
}
